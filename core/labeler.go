package cesr

import (
	"fmt"
	"slices"
	"strings"

	"github.com/jasoncolburne/cesrgo/common"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

type Labeler struct {
	matter
}

func NewLabeler(label *string, opts ...options.MatterOption) (*Labeler, error) {
	l := &Labeler{}

	config := &options.MatterOptions{}

	for _, opt := range opts {
		opt(config)
	}

	var (
		soft *string
		code types.Code
		raw  *types.Raw
	)
	if label != nil {
		re, err := common.ReAtt()
		if err != nil {
			return nil, err
		}

		if !re.MatchString(*label) {
			return nil, fmt.Errorf("invalid label: %s", *label)
		}

		if len(*label) < 1 || len(*label) > len(codex.TagCodex) {
			ws := (4 - (len(*label) % 4)) % 4
			if (*label)[0] == 'A' && (ws == 0 || ws == 1) {
				labelStr := "-" + *label
				label = &labelStr
			}
			code = codex.StrB64_L0
			rawBytes, err := rawify(*label)
			if err != nil {
				return nil, err
			}
			raw = &rawBytes
		} else {
			code, err = codify(*label)
			if err != nil {
				return nil, err
			}

			soft = label
		}

		opts = append(opts, options.WithCode(code))
		if raw != nil {
			opts = append(opts, options.WithRaw(*raw))
		}
		if soft != nil {
			opts = append(opts, options.WithSoft(*soft))
		}
	}

	if err := NewMatter(l, opts...); err != nil {
		return nil, err
	}

	if !slices.Contains(codex.LabelCodex, l.GetCode()) {
		return nil, fmt.Errorf("invalid code for labeler: %s", l.GetCode())
	}

	return l, nil
}

func (l *Labeler) Label() (string, error) {
	var (
		label string
		err   error
	)
	if slices.Contains(codex.TagCodex, l.GetCode()) {
		label = l.GetSoft()
	} else if slices.Contains(codex.BextCodex, l.GetCode()) {
		label, err = derawify(l.GetRaw(), l.GetCode())
		if err != nil {
			return "", err
		}

		label = strings.TrimPrefix(label, "-")
	} else {
		label = string(l.GetRaw())
	}

	re, err := common.ReAtt()
	if err != nil {
		return "", err
	}

	if !re.MatchString(label) {
		return "", fmt.Errorf("invalid label: %s", label)
	}

	return label, nil
}
