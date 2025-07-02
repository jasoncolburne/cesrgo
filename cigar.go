package cesrgo

import (
	"fmt"

	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
	"github.com/jasoncolburne/cesrgo/util"
)

type Cigar struct {
	matter
	verfer *Verfer
}

func (c *Cigar) GetVerfer() *Verfer {
	return c.verfer
}

var implementedCigarCodes = []types.Code{
	codex.Ed25519_Sig,
	codex.ECDSA_256k1_Sig,
	codex.ECDSA_256r1_Sig,
}

func NewCigar(verfer *Verfer, opts ...options.MatterOption) (*Cigar, error) {
	c := &Cigar{}

	if err := NewMatter(c, opts...); err != nil {
		return nil, err
	}

	if !util.ValidateCode(c.GetCode(), implementedCigarCodes) {
		return nil, fmt.Errorf("unexpected code: %s", c.GetCode())
	}

	if verfer != nil {
		c.verfer = verfer
	}

	return c, nil
}
