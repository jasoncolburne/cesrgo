package secp256r1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/jasoncolburne/cesrgo/types"
)

func Sign(sk types.Raw, ser []byte) (types.Raw, error) {
	curve := elliptic.P256()

	privateKey := new(big.Int).SetBytes(sk)
	if privateKey.Cmp(big.NewInt(0)) == 0 || privateKey.Cmp(curve.Params().N) >= 0 {
		return nil, fmt.Errorf("invalid private key")
	}

	priv := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
		},
		D: privateKey,
	}

	hash := sha256.Sum256(ser)

	r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
	if err != nil {
		return nil, err
	}

	sig := make([]byte, 64)
	copy(sig[:32], r.Bytes())
	copy(sig[32:], s.Bytes())

	return types.Raw(sig), nil
}
