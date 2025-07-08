package cesr

import (
	"encoding/base64"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

type Decimer struct {
	matter
}

func NewDecimer(dns *string, decimal *float64, opts ...options.MatterOption) (*Decimer, error) {
	d := &Decimer{}

	config := &options.MatterOptions{}
	for _, opt := range opts {
		opt(config)
	}

	var code types.Code
	if config.Code == nil {
		code = codex.Decimal_L0
	} else {
		code = *config.Code
	}

	if config.Raw == nil && config.Qb2 == nil && config.Qb64 == nil && config.Qb64b == nil {
		if decimal != nil {
			dnsStr := fmt.Sprintf("%f", *decimal)
			dns = &dnsStr
		}

		if dns == nil {
			return nil, fmt.Errorf("one of dns, decimal, raw, qb2, qb64, or qb64b is required")
		}

		decimal, err := strconv.ParseFloat(*dns, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid dns: %w", err)
		}

		if math.Trunc(decimal) == decimal {
			dnsStr := strconv.FormatInt(int64(decimal), 10)
			dns = &dnsStr
		} else {
			dnsStr := strconv.FormatFloat(decimal, 'f', -1, 64)
			dns = &dnsStr
		}

		raw, err := drawify(*dns)
		if err != nil {
			return nil, err
		}

		opts = []options.MatterOption{options.WithCode(code), options.WithRaw(raw)}
	}

	if config.Code == nil && config.Raw != nil {
		opts = append(opts, options.WithCode(code))
	}

	err := NewMatter(d, opts...)
	if err != nil {
		return nil, err
	}

	if !slices.Contains(codex.DecimalCodex, d.GetCode()) {
		return nil, fmt.Errorf("invalid code: %s", d.GetCode())
	}

	return d, nil
}

func drawify(dns string) (types.Raw, error) {
	dns = strings.ReplaceAll(dns, ".", "p")
	ts := len(dns) % 4
	ws := (4 - ts) % 4
	ls := (3 - ts) % 3
	base := strings.Repeat("A", ws) + dns
	raw, err := base64.URLEncoding.DecodeString(base)
	if err != nil {
		return nil, err
	}

	return raw[ls:], nil
}

func (d *Decimer) Dns() (string, error) {
	szg, ok := codex.Sizes[d.GetCode()]
	if !ok {
		return "", fmt.Errorf("invalid code: %s", d.GetCode())
	}

	raw := append(slices.Repeat([]byte{0}, int(szg.Ls)), d.GetRaw()...)
	dns := base64.URLEncoding.EncodeToString(raw)

	ws := 0
	if szg.Ls == 0 && dns != "" {
		if dns[0] == 'A' {
			ws = 1
		}
	} else {
		ws = (int(szg.Ls) + 1) % 4
	}

	dns = dns[ws:]
	dns = strings.ReplaceAll(dns, "p", ".")

	return dns, nil
}

func (d *Decimer) Decimal() (float64, error) {
	dns, err := d.Dns()
	if err != nil {
		return 0, err
	}

	decimal, err := strconv.ParseFloat(dns, 64)
	if err != nil {
		return 0, err
	}

	return decimal, nil
}
