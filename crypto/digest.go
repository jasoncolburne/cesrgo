package crypto

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/types"

	"github.com/zeebo/blake3"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/sha3"
)

func Digest(code types.Code, ser []byte) ([]byte, error) {
	switch code {
	case codex.Blake3_256:
		digest := blake3.Sum256(ser)
		return digest[:], nil
	case codex.Blake3_512:
		digest := blake3.Sum512(ser)
		return digest[:], nil
	case codex.Blake2b_256:
		digest := blake2b.Sum256(ser)
		return digest[:], nil
	case codex.Blake2b_512:
		digest := blake2b.Sum512(ser)
		return digest[:], nil
	case codex.Blake2s_256:
		digest := blake2s.Sum256(ser)
		return digest[:], nil
	case codex.SHA3_256:
		digest := sha3.Sum256(ser)
		return digest[:], nil
	case codex.SHA3_512:
		digest := sha3.Sum512(ser)
		return digest[:], nil
	case codex.SHA2_256:
		digest := sha256.Sum256(ser)
		return digest[:], nil
	case codex.SHA2_512:
		digest := sha512.Sum512(ser)
		return digest[:], nil
	default:
		return nil, fmt.Errorf("unimplemented digest code: %s", code)
	}
}
