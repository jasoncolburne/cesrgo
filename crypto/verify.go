package crypto

import (
	"fmt"

	"github.com/jasoncolburne/cesrgo/crypto/ed25519"
	"github.com/jasoncolburne/cesrgo/crypto/secp256k1"
	"github.com/jasoncolburne/cesrgo/crypto/secp256r1"
	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/types"
)

func VerifySignature(
	code types.Code,
	vk types.Raw,
	sig, ser []byte,
) error {
	switch code {
	case codex.Ed25519N, codex.Ed25519:
		return ed25519.Verify(sig, []byte(vk), ser)
	case codex.ECDSA_256k1N, codex.ECDSA_256k1:
		return secp256k1.Verify(sig, []byte(vk), ser)
	case codex.ECDSA_256r1N, codex.ECDSA_256r1:
		return secp256r1.Verify(sig, []byte(vk), ser)
	default:
		return fmt.Errorf("unsupported code: %s", code)
	}
}
