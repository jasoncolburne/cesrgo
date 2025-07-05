package cesr

import (
	"fmt"

	"github.com/jasoncolburne/cesrgo/common"
	"github.com/jasoncolburne/cesrgo/core/crypto"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
)

type Verfer struct {
	matter
}

func NewVerfer(opts ...options.MatterOption) (*Verfer, error) {
	v := &Verfer{}

	if err := NewMatter(v, opts...); err != nil {
		return nil, err
	}

	return v, nil
}

func (v *Verfer) Verify(sig, ser []byte) (bool, error) {
	if !common.ValidateCode(v.GetCode(), codex.PreNonDigCodex) {
		return false, fmt.Errorf("unexpected code: %s", v.GetCode())
	}

	if err := crypto.VerifySignature(v.GetCode(), v.GetRaw(), sig, ser); err != nil {
		return false, err
	}

	return true, nil
}
