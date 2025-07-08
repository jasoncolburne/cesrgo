package cesr

import (
	"crypto/rand"
	"fmt"
	"slices"

	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

type Noncer struct {
	matter
}

func NewNoncer(nonce []byte, opts ...options.MatterOption) (*Noncer, error) {
	n := &Noncer{}

	config := &options.MatterOptions{}
	for _, opt := range opts {
		opt(config)
	}

	var code types.Code

	if config.Code == nil {
		code = codex.Salt_128
	} else {
		code = *config.Code
	}

	if config.Raw == nil && config.Qb2 == nil && config.Qb64 == nil && config.Qb64b == nil {
		if nonce != nil {
			if len(nonce) == 0 {
				code = codex.Empty
				opts = []options.MatterOption{options.WithCode(code), options.WithRaw(nonce)}
			} else {
				opts = []options.MatterOption{options.WithQb64b(nonce)}
			}
		} else {
			switch code {
			case codex.Salt_128:
				raw := [16]byte{}
				_, err := rand.Read(raw[:])
				if err != nil {
					return nil, err
				}
				opts = []options.MatterOption{options.WithCode(code), options.WithRaw(raw[:])}
			case codex.Salt_256:
				raw := [32]byte{}
				_, err := rand.Read(raw[:])
				if err != nil {
					return nil, err
				}
				opts = []options.MatterOption{options.WithCode(code), options.WithRaw(raw[:])}
			default:
				return nil, fmt.Errorf("invalid nonce code: %s", code)
			}
		}
	}

	err := NewMatter(n, opts...)
	if err != nil {
		return nil, err
	}

	if !slices.Contains(codex.NonceCodex, n.GetCode()) {
		return nil, fmt.Errorf("invalid nonce code: %s", n.GetCode())
	}

	return n, nil
}

func (n *Noncer) Nonce() (string, error) {
	if n.GetCode() == codex.Empty {
		return "", nil
	}

	qb64, err := n.Qb64()
	if err != nil {
		return "", err
	}

	return string(qb64), nil
}

func (n *Noncer) Nonceb() ([]byte, error) {
	if n.GetCode() == codex.Empty {
		return []byte{}, nil
	}

	qb64b, err := n.Qb64()
	if err != nil {
		return nil, err
	}

	return []byte(qb64b), nil
}
