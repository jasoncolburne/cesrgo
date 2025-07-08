package cesr

import (
	"fmt"
	"slices"

	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
)

type Prefixer struct {
	matter
}

func NewPrefixer(opts ...options.MatterOption) (*Prefixer, error) {
	p := &Prefixer{}

	if err := NewMatter(p, opts...); err != nil {
		return nil, err
	}

	if !slices.Contains(codex.PreCodex, p.GetCode()) {
		return nil, fmt.Errorf("invalid code for prefixer: %s", p.GetCode())
	}

	return p, nil
}
