package secp256k1

import (
	"crypto/sha256"
	"fmt"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
)

func Verify(sig, vk, ser []byte) error {
	pub, err := secp256k1.ParsePubKey(vk)
	if err != nil {
		return err
	}

	r := &secp256k1.ModNScalar{}
	s := &secp256k1.ModNScalar{}

	r.SetByteSlice(sig[:32])
	s.SetByteSlice(sig[32:])

	signature := ecdsa.NewSignature(r, s)
	hash := sha256.Sum256(ser)

	if !signature.Verify(hash[:], pub) {
		return fmt.Errorf("invalid signature")
	}

	return nil
}
