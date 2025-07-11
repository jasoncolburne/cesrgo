package cesr

import (
	"fmt"
	"slices"

	"github.com/jasoncolburne/cesrgo"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

type Traitor struct {
	Tagger
}

func NewTraitor(trait *types.Trait, opts ...options.MatterOption) (*Traitor, error) {
	config := &options.MatterOptions{}

	for _, opt := range opts {
		opt(config)
	}

	var tag *string
	if config.Qb2 == nil && config.Qb64 == nil && config.Qb64b == nil {
		if trait == nil {
			return nil, fmt.Errorf("ilk is required")
		}

		if !slices.Contains(cesrgo.TRAITS, *trait) {
			return nil, fmt.Errorf("ilk must be one of %v", cesrgo.ILKS)
		}

		ilkStr := string(*trait)
		tag = &ilkStr
	}

	tagger, err := NewTagger(tag, opts...)
	if err != nil {
		return nil, err
	}

	i := &Traitor{Tagger: *tagger}

	if !slices.Contains([]types.Code{codex.Tag2, codex.Tag3}, i.GetCode()) {
		return nil, fmt.Errorf("traitor must be coded as tag2 or tag3")
	}

	return i, nil
}

func (i *Traitor) Trait() (types.Trait, error) {
	szg, ok := codex.Sizes[i.GetCode()]
	if !ok {
		return "", fmt.Errorf("unknown code %s", i.GetCode())
	}

	xs := int(szg.Xs)
	return types.Trait(i.GetSoft()[xs:]), nil
}
