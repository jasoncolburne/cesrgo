//nolint:dupl
package cesr

import (
	"fmt"

	"github.com/jasoncolburne/cesrgo/common"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
)

type Cigar struct {
	matter
	verfer *Verfer
}

func (c *Cigar) GetVerfer() *Verfer {
	return c.verfer
}

func NewCigar(verfer *Verfer, opts ...options.MatterOption) (*Cigar, error) {
	c := &Cigar{}

	if err := NewMatter(c, opts...); err != nil {
		return nil, err
	}

	if !common.ValidateCode(c.GetCode(), codex.SigCodex) {
		return nil, fmt.Errorf("unexpected code: %s", c.GetCode())
	}

	if verfer != nil {
		c.verfer = verfer
	}

	return c, nil
}
