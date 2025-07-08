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

type Pather struct {
	matter
}

func NewPather(path *string, parts []string, relative, pathive bool, opts ...options.MatterOption) (*Pather, error) {
	if parts == nil {
		parts = []string{}
	}

	p := &Pather{}

	config := options.MatterOptions{}

	for _, opt := range opts {
		opt(&config)
	}

	var code types.Code
	if config.Code == nil && config.Qb2 == nil && config.Qb64 == nil && config.Qb64b == nil {
		code = codex.StrB64_L0
	} else if config.Qb2 == nil && config.Qb64 == nil && config.Qb64b == nil {
		code = *config.Code
	}

	var raw types.Raw
	if config.Qb2 == nil && config.Qb64 == nil && config.Qb64b == nil {
		if config.Raw == nil {
			if len(parts) == 0 && path == nil {
				return nil, fmt.Errorf("missing path or parts")
			}

			if len(parts) == 0 {
				if strings.Contains(*path, "/") {
					parts = strings.Split(*path, "/")
				} else {
					parts = strings.Split(*path, "-")
				}
			}

			bextable := true
			re, err := common.RePath()
			if err != nil {
				return nil, err
			}

			for _, part := range parts {
				if !re.MatchString(part) {
					bextable = false
					break
				}
			}

			if !relative {
				if len(parts) > 0 && parts[0] != "" {
					parts = append([]string{""}, parts...)
				} else if len(parts) == 0 {
					parts = []string{"", ""}
				}
			}

			if bextable {
				pathStr := strings.Join(parts, "-")
				path = &pathStr

				if strings.Contains(*path, "--") {
					return nil, fmt.Errorf("non-unitary path separators: %s", *path)
				}
			} else {
				pathStr := strings.Join(parts, "/")
				path = &pathStr

				if strings.Contains(*path, "//") {
					return nil, fmt.Errorf("non-unitary path separators: %s", *path)
				}
			}

			if bextable {
				code = codex.StrB64_L0
				ws := (4 - (len(*path) % 4)) % 4
				if (*path)[0] == 'A' && (ws == 0 || ws == 1) {
					pathStr := "--" + *path
					path = &pathStr
				}
				raw, err = rawify(*path)
				if err != nil {
					return nil, err
				}
			} else {
				code = codex.Bytes_L0
				raw = types.Raw(*path)
			}
		} else {
			raw = *config.Raw
		}

		opts = []options.MatterOption{options.WithCode(code), options.WithRaw(raw)}
	}

	err := NewMatter(p, opts...)
	if err != nil {
		return nil, err
	}

	if !slices.Contains(codex.BextCodex, p.GetCode()) && !slices.Contains(codex.TextCodex, p.GetCode()) {
		return nil, fmt.Errorf("invalid code: %s", p.GetCode())
	}

	return p, nil
}

func (p *Pather) Path() (string, error) {
	if slices.Contains(codex.BextCodex, p.GetCode()) {
		path, err := derawify(p.GetRaw(), p.GetCode())
		if err != nil {
			return "", err
		}

		path = strings.TrimPrefix(path, "--")
		parts := strings.Split(path, "-")
		path = strings.Join(parts, "/")

		return path, nil
	} else {
		return string(p.GetRaw()), nil
	}
}

func (p *Pather) Parts() ([]string, error) {
	var parts []string

	if slices.Contains(codex.BextCodex, p.GetCode()) {
		path, err := derawify(p.GetRaw(), p.GetCode())
		if err != nil {
			return nil, err
		}

		path = strings.TrimPrefix(path, "--")
		parts = strings.Split(path, "-")
	} else {
		path := string(p.GetRaw())
		parts = strings.Split(path, "/")
	}

	return parts, nil
}
