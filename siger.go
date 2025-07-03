//nolint:dupl
package cesrgo

import (
	"fmt"

	"github.com/jasoncolburne/cesrgo/common/util"
	codex "github.com/jasoncolburne/cesrgo/indexer"
	"github.com/jasoncolburne/cesrgo/indexer/options"
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

	if !util.ValidateCode(s.GetCode(), codex.IndexedSigCodex) {
		return nil, fmt.Errorf("unexpected code: %s", s.GetCode())
	}

	if verfer != nil {
		s.verfer = verfer
	}

	return s, nil
}
