package cesr

import (
	"encoding/base64"
	"fmt"
	"slices"
	"strings"

	"github.com/jasoncolburne/cesrgo/common"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

type Bexter struct {
	matter
}

func NewBexter(bext *string, opts ...options.MatterOption) (*Bexter, error) {
	b := &Bexter{}

	config := &options.MatterOptions{}

	for _, opt := range opts {
		opt(config)
	}

	if config.Code == nil && config.Qb2 == nil && config.Qb64 == nil && config.Qb64b == nil {
		opts = append(opts, options.WithCode(codex.StrB64_L0))
	}

	if config.Raw == nil && config.Qb2 == nil && config.Qb64 == nil && config.Qb64b == nil {
		if bext == nil {
			return nil, fmt.Errorf("missing bext, raw, qb2, qb64, or qb64b")
		}

		re, err := common.ReB64()
		if err != nil {
			return nil, err
		}

		if !re.MatchString(*bext) {
			return nil, fmt.Errorf("invalid bext: %s", *bext)
		}

		raw, err := rawify(*bext)
		if err != nil {
			return nil, err
		}

		opts = append(opts, options.WithRaw(raw))
	}

	err := NewMatter(b, opts...)
	if err != nil {
		return nil, err
	}

	if !slices.Contains(codex.BextCodex, b.GetCode()) {
		return nil, fmt.Errorf("invalid bext: %s", b.GetCode())
	}

	return b, nil
}

func rawify(bext string) (types.Raw, error) {
	ts := len(bext) % 4
	ws := (4 - ts) % 4
	ls := (3 - ts) % 3
	b64 := types.Code(fmt.Sprintf("%s%s", strings.Repeat("A", ws), bext))
	raw, err := base64.URLEncoding.DecodeString(string(b64))
	if err != nil {
		return nil, err
	}

	return raw[ls:], nil
}

func derawify(raw types.Raw, code types.Code) (string, error) {
	szg, ok := codex.Sizes[code]
	if !ok {
		return "", fmt.Errorf("invalid code: %s", code)
	}

	padded := make([]byte, len(raw)+int(szg.Ls))
	copy(padded[int(szg.Ls):], raw)

	bext := base64.URLEncoding.EncodeToString(padded)

	ws := 0
	if szg.Ls == 0 && bext != "" {
		if bext[0] == 'A' {
			ws = 1
		}
	} else {
		ws = (int(szg.Ls) + 1) % 4
	}

	return bext[ws:], nil
}

func (b *Bexter) Bext() (string, error) {
	return derawify(b.GetRaw(), b.GetCode())
}
