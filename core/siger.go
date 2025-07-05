//nolint:dupl
package cesr

import (
	"fmt"

	"github.com/jasoncolburne/cesrgo/common"
	codex "github.com/jasoncolburne/cesrgo/core/indexer"
	"github.com/jasoncolburne/cesrgo/core/indexer/options"
)

type Siger struct {
	indexer
	verfer *Verfer
}

func (s *Siger) GetVerfer() *Verfer {
	return s.verfer
}

func NewSiger(verfer *Verfer, opts ...options.IndexerOption) (*Siger, error) {
	s := &Siger{}

	if err := NewIndexer(s, opts...); err != nil {
		return nil, err
	}

	if !common.ValidateCode(s.GetCode(), codex.IndexedSigCodex) {
		return nil, fmt.Errorf("unexpected code: %s", s.GetCode())
	}

	if verfer != nil {
		s.verfer = verfer
	}

	return s, nil
}
