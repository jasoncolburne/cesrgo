package crypto

import (
	"fmt"

	"github.com/jasoncolburne/cesrgo/crypto/ed25519"
	"github.com/jasoncolburne/cesrgo/crypto/secp256r1"
	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/types"
)

func VerifySignature(
	code types.Code,
	raw types.Raw,
	sig, ser []byte,
) error {
	switch code {
	case codex.Ed25519N, codex.Ed25519:
		return ed25519.Verify(sig, []byte(raw), ser)
	case codex.ECDSA_256r1N, codex.ECDSA_256r1:
		return secp256r1.Verify(sig, []byte(raw), ser)
	default:
		return fmt.Errorf("unsupported code: %s", code)
	}
}
