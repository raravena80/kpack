package registry

import (
	"time"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

type ImageFactory struct {
	KeychainFactory KeychainFactory
}

func (f *ImageFactory) NewRemote(imageRef ImageRef) (RemoteImage, error) {
	remoteImage, err := NewGoContainerRegistryImage(imageRef.Image(), f.KeychainFactory.KeychainForImageRef(imageRef))
	return remoteImage, err
}

func (f *ImageFactory) Rebase(orig, oldBase, latestBase ImageRef) (RemoteImage, error) {
	origImage, err := NewGoContainerRegistryImage(orig.Image(), f.KeychainFactory.KeychainForImageRef(orig))
	if err != nil {
		return nil, err
	}

	oldBaseImage, err := NewGoContainerRegistryImage(oldBase.Image(), f.KeychainFactory.KeychainForImageRef(oldBase))
	if err != nil {
		return nil, err
	}

	latestBaseImage, err := NewGoContainerRegistryImage(latestBase.Image(), f.KeychainFactory.KeychainForImageRef(latestBase))
	if err != nil {
		return nil, err
	}

	rebase, err := mutate.Rebase(origImage.image, oldBaseImage.image, latestBaseImage.image)
	if err != nil {
		return nil, err
	}

	reference, err := name.ParseReference(origImage.repoName, name.WeakValidation)
	if err != nil {
		return nil, err
	}

	rebasedImage := &GoContainerRegistryImage{
		repoName: origImage.repoName,
		image:    rebase,
	}

	return rebasedImage, remote.Write(reference, rebase, remote.WithAuthFromKeychain(f.KeychainFactory.KeychainForImageRef(orig)))
}

type KeychainFactory interface {
	KeychainForImageRef(ImageRef) authn.Keychain
}

type ImageRef interface {
	ServiceAccount() string
	Namespace() string
	Image() string
	HasSecret() bool
	SecretName() string
}

type noAuthImageRef struct {
	identifier string
}

func (na *noAuthImageRef) SecretName() string {
	return ""
}

func NewNoAuthImageRef(identifier string) *noAuthImageRef {
	return &noAuthImageRef{identifier: identifier}
}

func (na *noAuthImageRef) Image() string {
	return na.identifier
}

func (noAuthImageRef) ServiceAccount() string {
	return ""
}

func (noAuthImageRef) HasSecret() bool {
	return false
}

func (noAuthImageRef) Namespace() string {
	return ""
}

type RemoteImage interface {
	CreatedAt() (time.Time, error)
	Identifier() (string, error)
	Label(labelName string) (string, error)
	Env(key string) (string, error)
}

//go:generate counterfeiter . RemoteImageFactory
type RemoteImageFactory interface {
	NewRemote(imageRef ImageRef) (RemoteImage, error)
}
