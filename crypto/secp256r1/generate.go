package secp256r1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/jasoncolburne/cesrgo/types"
)

func GenerateSeed() (types.Raw, error) {
	curve := elliptic.P256()

	sk, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	seed := make([]byte, 32)
	sk.D.FillBytes(seed)

	return types.Raw(seed), nil
}

func DerivePublicKey(seed types.Raw) (types.Raw, error) {
	curve := elliptic.P256()

	privateKey := new(big.Int).SetBytes(seed)
	if privateKey.Cmp(big.NewInt(0)) == 0 || privateKey.Cmp(curve.Params().N) >= 0 {
		return nil, fmt.Errorf("invalid private key")
	}

	priv := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
		},
		D: privateKey,
	}

	pub, ok := priv.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to convert public key to bytes")
	}

	compressed := elliptic.MarshalCompressed(curve, pub.X, pub.Y)

	return types.Raw(compressed), nil
}
