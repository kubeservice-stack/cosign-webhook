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
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	wk "sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var cosignkeylog = logf.Log.WithName("cosignkey-resource")

func (r *CosignKey) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

var _ wk.Defaulter = &CosignKey{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *CosignKey) Default() {
	cosignkeylog.Info("default", "name", r.Name, "request", r)
	cosignkeylog.Info("Spec", "spec", r.Spec)
}

var _ wk.Validator = &CosignKey{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *CosignKey) ValidateCreate() (warnings admission.Warnings, err error) {
	cosignkeylog.Info("validate create", "name", r.Name, "request", r)

	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *CosignKey) ValidateUpdate(old runtime.Object) (warnings admission.Warnings, err error) {
	cosignkeylog.Info("validate update", "name", r.Name, "request", r)

	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *CosignKey) ValidateDelete() (warnings admission.Warnings, err error) {
	cosignkeylog.Info("validate delete", "name", r.Name)
	return nil, nil
}
