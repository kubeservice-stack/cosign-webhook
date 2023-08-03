/*
Copyright 2022 The KubeService-Stack Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package webhook

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	wk "sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var cosignkeylog = logf.Log.WithName("cosignkey-resource")

func (r *CustomCosignKey) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(&CosignKey{}).
		WithDefaulter(r).
		WithValidator(r).
		Complete()
}

type CustomCosignKey struct {
	Client client.Reader
}

var _ wk.CustomDefaulter = &CustomCosignKey{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *CustomCosignKey) Default(ctx context.Context, obj runtime.Object) error {
	n, ok := obj.(*CosignKey)
	if !ok {
		return errors.NewBadRequest(fmt.Sprintf("expected a CosignKey but got a %T", obj))
	}
	cosignkeylog.Info("default", "name", n.Name)
	cosignkeylog.Info("Spec", "spec", n.Spec)
	return nil
}

var _ wk.CustomValidator = &CustomCosignKey{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *CustomCosignKey) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	n, ok := obj.(*CosignKey)
	if !ok {
		return nil, errors.NewBadRequest(fmt.Sprintf("expected a CosignKey but got a %T", obj))
	}
	cosignkeylog.Info("validate create", "name", n.Name, "request", n)
	var allErrs field.ErrorList
	counter, err := getCRDCounter(r.Client, n.Namespace)
	if err != nil {
		allErrs = append(allErrs, field.Invalid(field.NewPath("spec").Child("authorities"),
			n.Spec.Auth,
			err.Error()))
	}
	if len(allErrs) == 0 && counter == 0 {
		return nil, nil
	}
	cosignkeylog.Info("validate cosignkey crd counter", "err", err, "field.ErrorList", allErrs)
	return nil, errors.NewInvalid(n.GroupVersionKind().GroupKind(), n.Name, allErrs)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *CustomCosignKey) ValidateUpdate(ctx context.Context, oldObj runtime.Object, newObj runtime.Object) (admission.Warnings, error) {
	cosignkeylog.Info("validate update")
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *CustomCosignKey) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	cosignkeylog.Info("validate delete")
	return nil, nil
}

func getCRDCounter(c client.Reader, namespace string) (int, error) {
	clrl := &CosignKeyList{}
	err := c.List(context.Background(), clrl, client.InNamespace(namespace))
	if err != nil {
		cosignkeylog.Info("Get CosignKey Resource Error", "namespace", namespace, "resource name", WebhookName)
		if errors.IsNotFound(err) {
			return 0, nil
		}
		return 0, ErrMissingCosignCRD
	}

	return len(clrl.Items), nil
}
