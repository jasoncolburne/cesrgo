package cesrgo

import (
	"fmt"

	"github.com/jasoncolburne/cesrgo/common/util"
	"github.com/jasoncolburne/cesrgo/crypto"
	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
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
	if !util.ValidateCode(v.GetCode(), codex.PreNonDigCodex) {
		return false, fmt.Errorf("unexpected code: %s", v.GetCode())
	}

	if err := crypto.VerifySignature(v.GetCode(), v.GetRaw(), sig, ser); err != nil {
		return false, err
	}

	return true, nil
}
