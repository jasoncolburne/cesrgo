package cesr

import (
	"fmt"
	"slices"

	"github.com/jasoncolburne/cesrgo"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

type Ilker struct {
	Tagger
}

func NewIlker(ilk *types.Ilk, opts ...options.MatterOption) (*Ilker, error) {
	config := &options.MatterOptions{}

	for _, opt := range opts {
		opt(config)
	}

	var tag *string
	if config.Qb2 == nil && config.Qb64 == nil && config.Qb64b == nil {
		if ilk == nil {
			return nil, fmt.Errorf("ilk is required")
		}

		if !slices.Contains(cesrgo.ILKS, *ilk) {
			return nil, fmt.Errorf("ilk must be one of %v", cesrgo.ILKS)
		}

		ilkStr := string(*ilk)
		tag = &ilkStr
	}

	tagger, err := NewTagger(tag, opts...)
	if err != nil {
		return nil, err
	}

	i := &Ilker{Tagger: *tagger}

	if !slices.Contains([]types.Code{codex.Tag3}, i.GetCode()) {
		return nil, fmt.Errorf("ilker must be coded as tag3")
	}

	return i, nil
}
