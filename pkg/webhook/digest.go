package webhook

import (
	"context"

	opt "github.com/google/go-containerregistry/pkg/authn/kubernetes"
	"github.com/google/go-containerregistry/pkg/crane"
)

func Digest(image string, ps opt.Options) (string, error) {
	keychain, err := opt.NewInCluster(context.Background(), ps)
	if err != nil {
		return "", err
	}
	return crane.Digest(image,
		crane.WithAuthFromKeychain(keychain),
	)
}
