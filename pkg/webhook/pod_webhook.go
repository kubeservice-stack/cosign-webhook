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
	"net/http"

	opt "github.com/google/go-containerregistry/pkg/authn/kubernetes"
	"github.com/kubeservice-stack/cosign-webhook/pkg/util"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var podlog = logf.Log.WithName("pod-webhook-resource")

// PodAnnotator validates Pods
type PodAnnotator struct {
	Client  client.Client
	decoder *admission.Decoder
}

func NewPodAnnotatorMutate(c client.Client, d *admission.Decoder) admission.Handler {
	return &PodAnnotator{Client: c, decoder: d}
}

// PodAnnotator adds an annotation to every incoming pods.
func (a *PodAnnotator) Handle(ctx context.Context, req admission.Request) admission.Response {
	podlog.Info("PodAnnotator", "req", req)
	pod := &corev1.Pod{}

	err := a.decoder.Decode(req, pod)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	ns := req.Namespace
	if ns == "" {
		ns = "default"
	}
	var imagepullsecretstrings []string
	for _, s := range pod.Spec.ImagePullSecrets {
		imagepullsecretstrings = append(imagepullsecretstrings, s.Name)
	}

	opts := opt.Options{
		Namespace:          pod.GetNamespace(),
		ServiceAccountName: pod.Spec.ServiceAccountName,
		ImagePullSecrets:   imagepullsecretstrings,
	}

	keys, err1 := a.GeCosignKeys(ns)

	for _, ic := range pod.Spec.InitContainers {
		_, err := util.Digest(ic.Image, opts)
		if err != nil {
			return admission.Denied(err.Error())
		}

		if err1 == nil && keys != nil {
			ok, err := a.ValidationCosignVerify(keys, ic.Image)
			if ok && err == nil {
				continue
			} else {
				return admission.Denied(err.Error())
			}
		}
	}

	for _, ic := range pod.Spec.Containers {
		_, err := util.Digest(ic.Image, opts)
		if err != nil {
			return admission.Denied(err.Error())
		}

		if err1 == nil && keys != nil {
			ok, err := a.ValidationCosignVerify(keys, ic.Image)
			if ok && err == nil {
				continue
			} else {
				return admission.Denied(err.Error())
			}
		}
	}

	return admission.Allowed("Check image cosign success")
}

func (a *PodAnnotator) GeCosignKeys(namespace string) ([]CosignKey, error) {
	clrl := &CosignKeyList{}
	err := a.Client.List(context.Background(), clrl, client.InNamespace(namespace))
	if err != nil {
		podlog.Info("Get CosignKey Resource Error", "namespace", namespace, "resource name", WebhookName)
		if errors.IsNotFound(err) {
			return nil, ErrMissingCosignCRD
		}
		return nil, ErrMissingCosignCRD
	}
	return clrl.Items, nil
}

func (a *PodAnnotator) ValidationCosignVerify(items []CosignKey, image string) (bool, error) {
	if len(items) > 1 {
		podlog.Info("Namespace has more than one CosignKey Resource", "count", len(items))
		return false, ErrInvalidCosignCRDMoreThanOne
	} else if len(items) <= 0 {
		podlog.Info("Namespace not found CosignKey Resource")
		return false, ErrMissingCosignCRDKeys
	} else {
		podlog.Info("PodAnnotator get CosignKey", "Cosign", items[0].Spec)

		for _, key := range items[0].Spec.Auth.Key {
			ok, err := util.VerifyPublicKey(image, key)
			if ok && err == nil {
				podlog.Info("VerifyPublicKey success", "public key", key, "image", image)
				return true, nil
			}
			continue
		}

	}
	podlog.Info("VerifyPublicKey failed", "image", image)
	return false, ErrInvalidCosignVerify
}
