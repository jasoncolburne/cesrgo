package cesrgo

import (
	"crypto/subtle"
	"fmt"

	"github.com/jasoncolburne/cesrgo/crypto"
	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
)

type Diger struct {
	matter
}

var validDigerCodes = []types.Code{
	codex.Blake3_256,
	codex.Blake3_512,
	codex.Blake2b_256,
	codex.Blake2b_512,
	codex.Blake2s_256,
	codex.SHA3_256,
	codex.SHA3_512,
	codex.SHA2_256,
	codex.SHA2_512,
}

func NewDiger(ser []byte, opts ...options.MatterOption) (*Diger, error) {
	d := &Diger{}

	if ser != nil {
		config := &options.MatterOptions{}
		for _, opt := range opts {
			opt(config)
		}

		if config.Code == nil {
			return nil, fmt.Errorf("code is required")
		}

		if !validateCode(*config.Code, validDigerCodes) {
			return nil, fmt.Errorf("unexpected code: %s", *config.Code)
		}

		digest, err := crypto.Digest(*config.Code, ser)
		if err != nil {
			return nil, err
		}

		if err := NewMatter(d, options.WithCode(*config.Code), options.WithRaw(digest)); err != nil {
			return nil, err
		}
	} else {
		if err := NewMatter(d, opts...); err != nil {
			return nil, err
		}

		if !validateCode(d.GetCode(), validDigerCodes) {
			return nil, fmt.Errorf("unexpected code: %s", d.GetCode())
		}
	}

	return d, nil
}

func (d *Diger) Verify(ser []byte) (bool, error) {
	digest, err := crypto.Digest(d.GetCode(), ser)
	if err != nil {
		return false, err
	}

	return subtle.ConstantTimeCompare(d.GetRaw(), digest) == 1, nil
}
