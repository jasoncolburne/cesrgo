package secp256r1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func Verify(sig, vk, ser []byte) error {
	if len(sig) != 64 {
		return fmt.Errorf("invalid signature length")
	}

	if len(vk) != 33 {
		return fmt.Errorf("invalid public key length")
	}

	curve := elliptic.P256()

	x, y := elliptic.UnmarshalCompressed(curve, vk)
	if x == nil || y == nil {
		return fmt.Errorf("invalid public key")
	}
	pub := &ecdsa.PublicKey{Curve: curve, X: x, Y: y}

	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:])

	hash := sha256.Sum256(ser)

	if !ecdsa.Verify(pub, hash[:], r, s) {
		return fmt.Errorf("invalid signature")
	}

	return nil
}
