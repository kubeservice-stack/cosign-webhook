/*
Copyright 2023 The KubeService-Stack Authors.

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

	opt "github.com/google/go-containerregistry/pkg/authn/kubernetes"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var cosignlog = logf.Log.WithName("cosign-webhook-resource")

// PodValidator validates Pods
type PodValidator struct{}

// validate admits a pod if a specific annotation exists.
func (v *PodValidator) validate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	pod, ok := obj.(*corev1.Pod)
	if !ok {
		return nil, fmt.Errorf("expected a Pod but got a %T", obj)
	}

	// It looks for image fields with tags in these sequence nodes:
	//   - `spec.containers`
	//   - `spec.initContainers`
	//   - `spec.template.spec.containers`
	//   - `spec.template.spec.initContainers`
	//   - `spec.jobTemplate.spec.template.spec.containers`
	//   - `spec.jobTemplate.spec.template.spec.initContainers`
	var imagepullsecretstrings []string
	for _, s := range pod.Spec.ImagePullSecrets {
		imagepullsecretstrings = append(imagepullsecretstrings, s.Name)
	}

	opts := opt.Options{
		Namespace:          pod.GetNamespace(),
		ServiceAccountName: pod.Spec.ServiceAccountName,
		ImagePullSecrets:   imagepullsecretstrings,
	}

	for _, ic := range pod.Spec.InitContainers {
		_, err := Digest(ic.Image, opts)
		if err != nil {
			return nil, fmt.Errorf("Parse image Digest fail. err: ", err)
		}

	}

	for _, ic := range pod.Spec.Containers {
		_, err := Digest(ic.Image, opts)
		if err != nil {
			return nil, fmt.Errorf("Parse image Digest fail. err: ", err)
		}
	}

	return nil, nil
}

func (v *PodValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	return v.validate(ctx, obj)
}

func (v *PodValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	return v.validate(ctx, newObj)
}

func (v *PodValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	return v.validate(ctx, obj)
}
