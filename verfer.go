package cesrgo

import (
	"fmt"

	"github.com/jasoncolburne/cesrgo/common/util"
	"github.com/jasoncolburne/cesrgo/crypto"
	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
)

type Verfer struct {
	matter
}

var implementedVerferCodes []types.Code = []types.Code{
	codex.Ed25519N,
	codex.Ed25519,
	codex.ECDSA_256k1N,
	codex.ECDSA_256k1,
	codex.ECDSA_256r1N,
	codex.ECDSA_256r1,
	// codex.Ed448N,
	// codex.Ed448,
}

func NewVerfer(opts ...options.MatterOption) (*Verfer, error) {
	v := &Verfer{}

	if err := NewMatter(v, opts...); err != nil {
		return nil, err
	}

	return v, nil
}

func (v *Verfer) Verify(sig, ser []byte) error {
	if !util.ValidateCode(v.GetCode(), implementedVerferCodes) {
		return fmt.Errorf("unexpected code: %s", v.GetCode())
	}

	return crypto.VerifySignature(v.GetCode(), v.GetRaw(), sig, ser)
}
