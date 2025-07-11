package cesr

import (
	"fmt"
	"math/big"

	"github.com/jasoncolburne/cesrgo/common"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
)

type Seqner struct {
	matter
}

func NewSeqner(sn *big.Int, snh *string, opts ...options.MatterOption) (*Seqner, error) {
	s := &Seqner{}

	config := options.MatterOptions{}

	for _, opt := range opts {
		opt(&config)
	}

	if config.Raw == nil && config.Qb2 == nil && config.Qb64 == nil && config.Qb64b == nil {
		if sn == nil && snh == nil {
			sn = big.NewInt(0)
		} else if snh != nil {
			bigNum := &big.Int{}
			bigNum.SetString(*snh, 16)
			sn = bigNum
		}

		if !common.LessThanMaxON(*sn) {
			return nil, fmt.Errorf("sn is too large to be represented by 16 octets")
		}

		if sn.Cmp(big.NewInt(0)) < 0 {
			return nil, fmt.Errorf("sn must be greater than or equal to 0")
		}

		rs, err := rawSize(codex.Huge)
		if err != nil {
			return nil, err
		}

		raw := make([]byte, rs)
		sn.FillBytes(raw)
		opts = append(opts, options.WithRaw(raw))

		if config.Code == nil {
			opts = append(opts, options.WithCode(codex.Huge))
		}
	}
	if err := NewMatter(s, opts...); err != nil {
		return nil, err
	}

	if s.GetCode() != codex.Huge {
		return nil, fmt.Errorf("seqner must be coded as huge")
	}

	return s, nil
}

func (s *Seqner) Sn() big.Int {
	bigNum := big.Int{}
	bigNum.SetBytes(s.GetRaw())

	return bigNum
}

func (s *Seqner) Snh() (string, error) {
	sn := s.Sn()
	rs, err := rawSize(s.GetCode())
	if err != nil {
		return "", err
	}

	bytes := make([]byte, rs)
	sn.FillBytes(bytes)
	return fmt.Sprintf("%X", bytes), nil
}
