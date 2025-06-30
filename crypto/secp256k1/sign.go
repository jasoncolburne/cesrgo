package secp256k1

import (
	"crypto/sha256"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
	"github.com/jasoncolburne/cesrgo/types"
)

func Sign(sk types.Raw, ser []byte) (types.Raw, error) {
	priv := secp256k1.PrivKeyFromBytes(sk)
	hash := sha256.Sum256(ser)

	signature := ecdsa.Sign(priv, hash[:])
	bytes := signature.Serialize()

	return types.Raw(bytes), nil
}
