package cesr

import (
	"crypto/rand"
	"fmt"

	"github.com/jasoncolburne/cesrgo/common"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
	"golang.org/x/crypto/argon2"
)

type Salter struct {
	matter
	tier types.Tier
}

func NewSalter(tier *types.Tier, opts ...options.MatterOption) (*Salter, error) {
	s := &Salter{}

	config := &options.MatterOptions{}
	for _, opt := range opts {
		opt(config)
	}

	code := codex.Salt_128
	if config.Code != nil {
		code = *config.Code
	}

	if tier == nil {
		tier = &common.TIER_LOW
	}

	if config.Raw != nil {
		if err := NewMatter(s, options.WithCode(code), options.WithRaw(*config.Raw)); err != nil {
			return nil, err
		}
	} else {
		var raw types.Raw
		switch code {
		case codex.Salt_128:
			bytes := [16]byte{}
			_, err := rand.Read(bytes[:])
			if err != nil {
				return nil, err
			}

			raw = bytes[:]
		case codex.Salt_256:
			bytes := [32]byte{}
			_, err := rand.Read(bytes[:])
			if err != nil {
				return nil, err
			}

			raw = bytes[:]
		default:
			return nil, fmt.Errorf("unimplemented code: %s", code)
		}

		if err := NewMatter(s, options.WithCode(code), options.WithRaw(raw)); err != nil {
			return nil, err
		}
	}

	s.tier = *tier

	return s, nil
}

func (s *Salter) Stretch(size *types.Size, path *string, tier *types.Tier, temp *bool) ([]byte, error) {
	if size == nil {
		sz := types.Size(32)
		size = &sz
	}

	if tier == nil {
		tier = &s.tier
	}

	if temp == nil {
		f := false
		temp = &f
	}

	opsLimit := uint32(1)
	memLimit := uint32(8192 / 1024)

	if !*temp {
		switch *tier {
		case common.TIER_LOW:
			opsLimit = 2
			memLimit = 67108864 / 1024
		case common.TIER_MED:
			opsLimit = 3
			memLimit = 268435456 / 1024
		case common.TIER_HIGH:
			opsLimit = 4
			memLimit = 1073741824 / 1024
		default:
			return nil, fmt.Errorf("unsupported security tier: %s", *tier)
		}
	}

	passwd := []byte{}
	if path != nil {
		passwd = []byte(*path)
	}

	seed := argon2.IDKey(passwd, s.GetRaw(), opsLimit, memLimit, 1, uint32(*size))

	return seed, nil
}
