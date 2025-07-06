package cesr

import (
	"fmt"
	"slices"

	"github.com/jasoncolburne/cesrgo/common"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

type Tagger struct {
	matter
}

func NewTagger(tag *string, opts ...options.MatterOption) (*Tagger, error) {
	t := &Tagger{}

	config := &options.MatterOptions{}
	for _, opt := range opts {
		opt(config)
	}

	if tag != nil && config.Soft != nil {
		return nil, fmt.Errorf("tag and soft are mutually exclusive")
	}

	if config.Soft != nil {
		tag = config.Soft
	}

	if tag != nil {
		re, err := common.Reb64()
		if err != nil {
			return nil, err
		}

		if !re.MatchString(*tag) {
			return nil, fmt.Errorf("tag %s is not a valid b64 string", *tag)
		}

		code, err := codify(*tag)
		if err != nil {
			return nil, err
		}

		soft := *tag

		if config.Code != nil && *config.Code != code {
			return nil, fmt.Errorf("provided code %s does not match derived code %s", *config.Code, code)
		}

		opts = append(opts, options.WithCode(code), options.WithSoft(soft))
		if err := NewMatter(t, opts...); err != nil {
			return nil, err
		}
	} else if config.Qb2 == nil && config.Qb64 == nil && config.Qb64b == nil {
		return nil, fmt.Errorf("tag is required when qb2, qb64, or qb64b is not provided")
	} else {
		if err := NewMatter(t, opts...); err != nil {
			return nil, err
		}
	}

	if !slices.Contains(codex.TagCodex, t.GetCode()) {
		return nil, fmt.Errorf("code %s is not valid for a tag", t.GetCode())
	}

	return t, nil
}

func codify(tag string) (types.Code, error) {
	l := len(tag)

	if l < 1 || l > len(codex.TagCodex) {
		return "", fmt.Errorf("tag %s empty of oversized", tag)
	}

	return codex.TagCodex[l-1], nil
}
