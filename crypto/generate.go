package crypto

import (
	"fmt"

	"github.com/jasoncolburne/cesrgo/crypto/ed25519"
	"github.com/jasoncolburne/cesrgo/crypto/secp256k1"
	"github.com/jasoncolburne/cesrgo/crypto/secp256r1"
	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/types"
)

func GenerateSeed(code types.Code) (types.Raw, error) {
	var (
		seed types.Raw
		err  error
	)

	switch code {
	case codex.Ed25519_Seed:
		seed, err = ed25519.GenerateSeed()
	case codex.ECDSA_256r1_Seed:
		seed, err = secp256r1.GenerateSeed()
	case codex.ECDSA_256k1_Seed:
		seed, err = secp256k1.GenerateSeed()
	default:
		return nil, fmt.Errorf("unimplemented seed code: %s", code)
	}

	return seed, err
}

func DeriveCodeAndPublicKey(code types.Code, raw types.Raw, transferable bool) (types.Code, types.Raw, error) {
	var verferCode types.Code
	var verferRaw types.Raw

	switch code {
	case codex.Ed25519_Seed:
		if transferable {
			verferCode = codex.Ed25519
		} else {
			verferCode = codex.Ed25519N
		}

		var err error
		if verferRaw, err = ed25519.DerivePublicKey(raw); err != nil {
			return "", nil, err
		}
	case codex.ECDSA_256r1_Seed:
		if transferable {
			verferCode = codex.ECDSA_256r1
		} else {
			verferCode = codex.ECDSA_256r1N
		}

		var err error
		if verferRaw, err = secp256r1.DerivePublicKey(raw); err != nil {
			return "", nil, err
		}
	case codex.ECDSA_256k1_Seed:
		if transferable {
			verferCode = codex.ECDSA_256k1
		} else {
			verferCode = codex.ECDSA_256k1N
		}

		var err error
		if verferRaw, err = secp256k1.DerivePublicKey(raw); err != nil {
			return "", nil, err
		}
	default:
		return "", nil, fmt.Errorf("unimplemented seed code")
	}

	return verferCode, verferRaw, nil
}
