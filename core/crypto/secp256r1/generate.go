package secp256r1

import (
	"crypto/ecdh"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/jasoncolburne/cesrgo/core/types"
)

func GenerateSeed() (types.Raw, error) {
	curve := ecdh.P256()

	key, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	seed := [32]byte{}
	byteLen := len(key.Bytes())
	if byteLen <= 32 {
		copy(seed[32-byteLen:], key.Bytes())
	} else {
		copy(seed[:], key.Bytes()[byteLen-32:])
	}

	return types.Raw(seed[:]), nil
}

func DerivePublicKey(seed types.Raw) (types.Raw, error) {
	curve := ecdh.P256()

	privateKey, err := curve.NewPrivateKey([]byte(seed))
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.PublicKey()
	pubBytes := publicKey.Bytes() // 65 bytes: 0x04 || X (32) || Y (32)

	if len(pubBytes) != 65 || pubBytes[0] != 0x04 {
		return nil, fmt.Errorf("unexpected public key format")
	}

	x := new(big.Int).SetBytes(pubBytes[1:33])
	y := new(big.Int).SetBytes(pubBytes[33:65])

	// Compressed format: 0x02 if y is even, 0x03 if y is odd
	prefix := byte(0x02)
	if y.Bit(0) == 1 {
		prefix = 0x03
	}

	compressed := make([]byte, 33)
	compressed[0] = prefix
	bytes := [32]byte{}
	x.FillBytes(bytes[:])
	copy(compressed[1:], bytes[:])

	return types.Raw(compressed), nil
}
