package crypto

import (
	"fmt"

	"github.com/jasoncolburne/cesrgo/core/crypto/ed25519"
	"github.com/jasoncolburne/cesrgo/core/crypto/secp256k1"
	"github.com/jasoncolburne/cesrgo/core/crypto/secp256r1"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/types"
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
