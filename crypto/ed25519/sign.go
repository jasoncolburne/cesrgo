package ed25519

import (
	"crypto/ed25519"

	"github.com/jasoncolburne/cesrgo/types"
)

func Sign(sk types.Raw, ser []byte) (types.Raw, error) {
	var priv = ed25519.PrivateKey(sk)
	sig := ed25519.Sign(priv, ser)

	return types.Raw(sig), nil
}
