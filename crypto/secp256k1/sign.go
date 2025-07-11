package secp256k1

import (
	"crypto/sha256"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
	"github.com/jasoncolburne/cesrgo/core/types"
)

func Sign(sk types.Raw, ser []byte) (types.Raw, error) {
	priv := secp256k1.PrivKeyFromBytes(sk)
	hash := sha256.Sum256(ser)

	signature := ecdsa.Sign(priv, hash[:])
	r := signature.R()
	s := signature.S()

	rBytes := r.Bytes()
	sBytes := s.Bytes()

	bytes := make([]byte, 64)
	copy(bytes[:32], rBytes[:])
	copy(bytes[32:], sBytes[:])

	return types.Raw(bytes), nil
}
