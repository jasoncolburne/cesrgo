package cesrgo

import (
	"fmt"
	"strings"

	"github.com/jasoncolburne/cesrgo/crypto"
	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
)

type Saider struct {
	matter
}

var validSaiderCodes = []types.Code{
	codex.Blake3_256,
	codex.Blake3_512,
	codex.Blake2b_256,
	codex.Blake2b_512,
	codex.Blake2s_256,
	codex.SHA3_256,
	codex.SHA3_512,
	codex.SHA2_256,
	codex.SHA2_512,
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
		om := sad.Map()
		value, ok := om.Get(*label)
		if !ok {
			return nil, fmt.Errorf("label not found: %s", *label)
		}

		valueQb64, ok := value.(types.Qb64)
		if !ok {
			return nil, fmt.Errorf("value is not a string")
		}

		empty := valueQb64 == types.Qb64("")
		if code == nil {
			if !empty {
				err := NewMatter(s, options.WithQb64(valueQb64))
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

		if !validateCode(*code, validSaiderCodes) {
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

	if !validateCode(s.code, validSaiderCodes) {
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
		kindJson := Kind_JSON
		kind = &kindJson
	}

	if !validateCode(*code, validSaiderCodes) {
		return nil, types.Map{}, fmt.Errorf("unexpected code: %s", *code)
	}

	sadCopy := *sad
	sadOm := sadCopy.Map()

	_, ok := sadOm.Get(*label)
	if !ok {
		return nil, types.Map{}, fmt.Errorf("label not found: %s", *label)
	}

	szg, err := codex.GetSizage(*code)
	if err != nil {
		return nil, types.Map{}, fmt.Errorf("failed to get sizage: %w", err)
	}

	if szg.Fs == nil {
		return nil, types.Map{}, fmt.Errorf("programmer error: sizage fs is nil")
	}

	dummy := strings.Repeat("#", int(*szg.Fs))
	sadOm.Set(*label, dummy)

	for _, key := range ignore {
		sadOm.Delete(key)
	}

	cpa, err := marshal(sadCopy, kind)
	if err != nil {
		return nil, types.Map{}, fmt.Errorf("failed to marshal: %w", err)
	}

	digest, err := crypto.Digest(*code, []byte(cpa))
	if err != nil {
		return nil, types.Map{}, fmt.Errorf("failed to digest: %w", err)
	}

	return types.Raw(digest), sadCopy, nil
}
