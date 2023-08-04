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
	"path"

	"github.com/docker/distribution/reference"
	opt "github.com/google/go-containerregistry/pkg/authn/kubernetes"
	"github.com/google/go-containerregistry/pkg/crane"
)

type ContainerImage struct {
	Repository string
	Name       string
	Tag        string
	Digest     string
}

func ImagesToStuct(image string) (ContainerImage, error) {
	n, err := reference.ParseNamed(image)
	if err != nil {
		return ContainerImage{}, err
	}

	var repository, tag, digest string
	_, nameOnly := path.Split(reference.Path(n))
	if nameOnly != "" {
		lenOfCompleteName := len(n.Name())
		repository = n.Name()[:lenOfCompleteName-len(nameOnly)-1]
	}

	tagged, ok := n.(reference.Tagged)
	if ok {
		tag = tagged.Tag()
	}

	digested, ok := n.(reference.Digested)
	if ok {
		digest = digested.Digest().String()
	}

	return ContainerImage{Repository: repository, Name: nameOnly, Tag: tag, Digest: digest}, nil
}

func Digest(image string, ps opt.Options) (string, error) {
	keychain, err := opt.NewInCluster(context.Background(), ps)
	if err != nil {
		return "", err
	}
	return crane.Digest(image,
		crane.WithAuthFromKeychain(keychain),
	)
}
