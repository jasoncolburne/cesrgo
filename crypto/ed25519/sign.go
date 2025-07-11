package ed25519

import (
	"crypto/ed25519"

	"github.com/jasoncolburne/cesrgo/core/types"
)

func Sign(sk types.Raw, ser []byte) (types.Raw, error) {
	var priv = ed25519.NewKeyFromSeed(sk)
	sig := ed25519.Sign(priv, ser)

	return types.Raw(sig), nil
}
