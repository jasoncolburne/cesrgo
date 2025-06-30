package cesrgo

import (
	"fmt"

	codex "github.com/jasoncolburne/cesrgo/indexer"
	"github.com/jasoncolburne/cesrgo/indexer/options"
	"github.com/jasoncolburne/cesrgo/types"
)

type Siger struct {
	indexer
	verfer *Verfer
}

var validSigerCodes = []types.Code{
	codex.Ed25519,
	codex.Ed25519_Crt,
	// codex.ECDSA_256k1,
	// codex.ECDSA_256k1_Crt,
	codex.ECDSA_256r1,
	codex.ECDSA_256r1_Crt,
	// codex.Ed448,
	// codex.Ed448_Crt,
	codex.Ed25519_Big,
	codex.Ed25519_Big_Crt,
	// codex.ECDSA_256k1_Big,
	// codex.ECDSA_256k1_Big_Crt,
	codex.ECDSA_256r1_Big,
	codex.ECDSA_256r1_Big_Crt,
	// codex.Ed448_Big,
	// codex.Ed448_Big_Crt,
}

func NewSiger(verfer *Verfer, opts ...options.IndexerOption) (*Siger, error) {
	s := &Siger{}

	if err := NewIndexer(s, opts...); err != nil {
		return nil, err
	}

	if !validateCode(s.GetCode(), validSigerCodes) {
		return nil, fmt.Errorf("unexpected code: %s", s.GetCode())
	}

	if verfer != nil {
		s.verfer = verfer
	}

	return s, nil
}
