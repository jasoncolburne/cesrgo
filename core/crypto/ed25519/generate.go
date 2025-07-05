package ed25519

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"

	"github.com/jasoncolburne/cesrgo/core/types"
)

func GenerateSeed() (types.Raw, error) {
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	return priv.Seed(), nil
}

func DerivePublicKey(seed types.Raw) (types.Raw, error) {
	priv := ed25519.NewKeyFromSeed(seed)

	pub := priv.Public()

	pubBytes, ok := pub.(ed25519.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to convert public key to bytes")
	}

	return types.Raw(pubBytes), nil
}
