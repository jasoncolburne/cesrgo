package cesr

import (
	"fmt"
	"strings"
	"time"

	"github.com/jasoncolburne/cesrgo/common"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

type Dater struct {
	matter
}

func NewDater(dts *types.DateTime, opts ...options.MatterOption) (*Dater, error) {
	d := &Dater{}

	config := &options.MatterOptions{}
	for _, opt := range opts {
		opt(config)
	}

	if config.Code != nil && *config.Code != codex.DateTime {
		return nil, fmt.Errorf("code %s is not valid for a dater", *config.Code)
	}

	var qb64 *types.Qb64
	if config.Raw == nil && config.Qb2 == nil && config.Qb64 == nil && config.Qb64b == nil {
		var dtsStr string
		if dts == nil {
			dtsStr = string(common.NowISO8601())
		} else {
			dtsStr = string(*dts)
		}

		dtsStr = strings.ReplaceAll(dtsStr, ":", "c")
		dtsStr = strings.ReplaceAll(dtsStr, ".", "d")
		dtsStr = strings.ReplaceAll(dtsStr, "+", "p")

		qb64Str := types.Qb64(fmt.Sprintf("%s%s", codex.DateTime, dtsStr))
		qb64 = &qb64Str
	}

	if qb64 != nil {
		opts = append(opts, options.WithQb64(*qb64))

		err := NewMatter(d, opts...)
		if err != nil {
			return nil, err
		}
	} else {
		err := NewMatter(d, opts...)
		if err != nil {
			return nil, err
		}
	}

	if d.GetCode() != codex.DateTime {
		return nil, fmt.Errorf("code %s is not valid for a dater", d.GetCode())
	}

	return d, nil
}

func (d *Dater) DTS() (types.DateTime, error) {
	qb64, err := d.Qb64()
	if err != nil {
		return "", err
	}

	szg, ok := codex.Sizes[d.GetCode()]
	if !ok {
		return "", fmt.Errorf("code %s is not valid for a dater", d.GetCode())
	}

	if szg.Fs == nil {
		return "", fmt.Errorf("programmer error: fs nil for dater")
	}

	dts := string(qb64)[szg.Hs:]

	dts = strings.ReplaceAll(dts, "c", ":")
	dts = strings.ReplaceAll(dts, "d", ".")
	dts = strings.ReplaceAll(dts, "p", "+")

	return types.DateTime(dts), nil
}

func (d *Dater) DTSb() ([]byte, error) {
	dts, err := d.DTS()
	if err != nil {
		return nil, err
	}

	return []byte(dts), nil
}

func (d *Dater) DateTime() (time.Time, error) {
	dts, err := d.DTS()
	if err != nil {
		return time.Time{}, err
	}

	return time.Parse(time.RFC3339, string(dts))
}
