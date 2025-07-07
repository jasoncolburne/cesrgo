package cesr

import (
	"fmt"
	"slices"

	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

type Texter struct {
	matter
}

func NewTexter(text *string, opts ...options.MatterOption) (*Texter, error) {
	t := &Texter{}

	config := &options.MatterOptions{}

	for _, opt := range opts {
		opt(config)
	}

	if config.Raw == nil && config.Qb2 == nil && config.Qb64 == nil && config.Qb64b == nil {
		if text == nil {
			return nil, fmt.Errorf("text cannot be empty if raw and qualified options are omitted")
		}

		raw := types.Raw(*text)
		opts = append(opts, options.WithRaw(raw))
	}

	if err := NewMatter(t, opts...); err != nil {
		return nil, err
	}

	if !slices.Contains(codex.TextCodex, t.GetCode()) {
		return nil, fmt.Errorf("invalid code: %s", t.GetCode())
	}

	return t, nil
}

func (t *Texter) Text() string {
	return string(t.GetRaw())
}
