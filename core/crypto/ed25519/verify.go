package ed25519

import (
	"crypto/ed25519"
	"fmt"
)

const VK_SIZE = 32
const SIG_SIZE = 64

func Verify(sig, vk, ser []byte) error {
	if len(sig) != SIG_SIZE {
		return fmt.Errorf("invalid signature length")
	}

	if len(vk) != VK_SIZE {
		return fmt.Errorf("invalid public key length")
	}

	if !ed25519.Verify(vk, ser, sig) {
		return fmt.Errorf("invalid signature")
	}

	return nil
}
