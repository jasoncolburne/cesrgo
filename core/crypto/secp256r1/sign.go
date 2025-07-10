package secp256r1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/codahale/rfc6979"

	"github.com/jasoncolburne/cesrgo/core/types"
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

	r, s, err := rfc6979.SignECDSA(priv, hash[:], sha256.New)
	if err != nil {
		return nil, err
	}

	bytes := make([]byte, 64)
	r.FillBytes(bytes[:32])
	s.FillBytes(bytes[32:])

	return types.Raw(bytes), nil
}
