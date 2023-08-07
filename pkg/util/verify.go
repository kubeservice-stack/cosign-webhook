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

package util

import (
	"context"
	"crypto"
	"fmt"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/sigstore/cosign/v2/pkg/cosign"
	ociremote "github.com/sigstore/cosign/v2/pkg/oci/remote"
	sigs "github.com/sigstore/cosign/v2/pkg/signature"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// log is for logging in this package.
var utillog = logf.Log.WithName("cosign-util-resource")

func VerifyPublicKey(image string, pubkey string) (bool, error) {
	ref, err := name.ParseReference(image)
	if err != nil {
		utillog.Info("failed to parse image ref", "image", image, "err", err.Error())
		return false, fmt.Errorf("failed to parse image ref `%s`; %s", image, err.Error())
	}

	digested, err := ociremote.ResolveDigest(ref)
	if err != nil {
		utillog.Info("failed to parse ociremote digested", "image", image, "err", err.Error())
		return false, fmt.Errorf("failed to parse ociremote digested`%s`; %s", image, err.Error())
	}

	o := &cosign.CheckOpts{
		Annotations:   nil,
		ClaimVerifier: cosign.SimpleClaimVerifier,
		IgnoreSCT:     true,
		IgnoreTlog:    true,
		Offline:       true,
	}

	pubKeyVerifier, err := sigs.LoadPublicKeyRaw([]byte(pubkey), crypto.SHA256)
	if err != nil {
		utillog.Info("failed to load public key", "err", err.Error())
		return false, fmt.Errorf("failed to load public key; %s", err.Error())
	}

	o.SigVerifier = pubKeyVerifier

	checkedSigs, _, err := cosign.VerifyImageSignatures(context.Background(), digested, o)
	if err != nil {
		utillog.Info("error occured while verifying image", "image", image, "err", err.Error())
		return false, fmt.Errorf("error occured while verifying image `%s`; %s", image, err.Error())
	}
	if len(checkedSigs) == 0 {
		utillog.Info("no verified signatures in the image", "image", image, "err", err.Error())
		return false, fmt.Errorf("no verified signatures in the image `%s`; %s", image, err.Error())
	}

	return true, nil
}
