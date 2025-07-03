package cesrgo

import (
	"fmt"
	"strings"

	"github.com/jasoncolburne/cesrgo/common"
	"github.com/jasoncolburne/cesrgo/common/util"
	"github.com/jasoncolburne/cesrgo/crypto"
	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
)

type Saider struct {
	matter
}

func NewSaider(
	sad *types.Map,
	label *string,
	kind *types.Kind,
	opts ...options.MatterOption,
) (*Saider, error) {
	s := &Saider{}

	config := &options.MatterOptions{}
	for _, opt := range opts {
		opt(config)
	}

	code := config.Code
	raw := config.Raw

	if label == nil {
		labelStr := "d"
		label = &labelStr
	}

	if code != nil && raw != nil {
		err := NewMatter(s, options.WithCode(*code), options.WithRaw(*raw))
		if err != nil {
			return nil, err
		}
	} else if config.Qb2 != nil || config.Qb64 != nil || config.Qb64b != nil {
		err := NewMatter(s, opts...)
		if err != nil {
			return nil, err
		}
	} else if sad == nil {
		return nil, fmt.Errorf("code and raw or sad is required")
	} else {
		value, ok := sad.Get(*label)
		if !ok {
			return nil, fmt.Errorf("label not found: %s", *label)
		}

		valueQb64, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("value is not a string")
		}

		empty := valueQb64 == ""
		if code == nil {
			if !empty {
				err := NewMatter(s, options.WithQb64(types.Qb64(valueQb64)))
				if err != nil {
					return nil, err
				}

				codeStr := s.GetCode()
				code = &codeStr
			} else {
				codeStr := codex.Blake3_256
				code = &codeStr
			}
		}

		if !util.ValidateCode(*code, codex.DigCodex) {
			return nil, fmt.Errorf("unexpected code: %s", s.code)
		}

		rawValue, _, err := derive(sad, code, kind, label, []string{})
		if err != nil {
			return nil, err
		}

		if err := NewMatter(s, options.WithCode(*code), options.WithRaw(rawValue)); err != nil {
			return nil, err
		}
	}

	if !util.ValidateCode(s.code, codex.DigCodex) {
		return nil, fmt.Errorf("unexpected code: %s", s.code)
	}

	return s, nil
}

func derive(sad *types.Map, code *types.Code, kind *types.Kind, label *string, ignore []string) (types.Raw, types.Map, error) {
	if code == nil {
		codeBlake3 := codex.Blake3_256
		code = &codeBlake3
	}

	if label == nil {
		labelStr := "d"
		label = &labelStr
	}

	if kind == nil {
		kindJson := common.Kind_JSON
		kind = &kindJson
	}

	if !util.ValidateCode(*code, codex.DigCodex) {
		return nil, types.Map{}, fmt.Errorf("unexpected code: %s", *code)
	}

	sadCopy := sad.Clone()
	_, ok := sadCopy.Get(*label)
	if !ok {
		return nil, types.Map{}, fmt.Errorf("label not found: %s", *label)
	}

	szg, ok := codex.Sizes[*code]
	if !ok {
		return nil, types.Map{}, fmt.Errorf("unknown code: %s", *code)
	}

	if szg.Fs == nil {
		return nil, types.Map{}, fmt.Errorf("programmer error: sizage fs is nil")
	}

	dummy := strings.Repeat("#", int(*szg.Fs))
	_, ok = sadCopy.Set(*label, dummy)
	if !ok {
		return nil, types.Map{}, fmt.Errorf("failed to set dummy")
	}

	for _, key := range ignore {
		_, ok = sadCopy.Delete(key)
		if !ok {
			return nil, types.Map{}, fmt.Errorf("failed to delete key: %s", key)
		}
	}

	cpa, err := util.Marshal(sadCopy, kind)
	if err != nil {
		return nil, types.Map{}, fmt.Errorf("failed to marshal: %w", err)
	}

	digest, err := crypto.Digest(*code, []byte(cpa))
	if err != nil {
		return nil, types.Map{}, fmt.Errorf("failed to digest: %w", err)
	}

	return types.Raw(digest), sadCopy, nil
}
