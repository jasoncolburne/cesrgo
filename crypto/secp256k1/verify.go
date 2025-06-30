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

	signature, err := ecdsa.ParseDERSignature(sig)
	if err != nil {
		return err
	}

	if signature == nil {
		return fmt.Errorf("invalid signature")
	}

	hash := sha256.Sum256(ser)

	if !signature.Verify(hash[:], pub) {
		return fmt.Errorf("invalid signature")
	}

	return nil
}
