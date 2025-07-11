package crypto

import (
	"fmt"

	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/types"
	"github.com/jasoncolburne/cesrgo/crypto/ed25519"
	"github.com/jasoncolburne/cesrgo/crypto/secp256k1"
	"github.com/jasoncolburne/cesrgo/crypto/secp256r1"
)

func Sign(code types.Code, sk types.Raw, ser []byte) (types.Raw, error) {
	switch code {
	case codex.Ed25519_Seed:
		return ed25519.Sign(sk, ser)
	case codex.ECDSA_256k1_Seed:
		return secp256k1.Sign(sk, ser)
	case codex.ECDSA_256r1_Seed:
		return secp256r1.Sign(sk, ser)
	default:
		return nil, fmt.Errorf("unimplemented seed code: %s", code)
	}
}
