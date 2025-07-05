package cesr

import (
	"fmt"
	"math/big"
	"slices"
	"strings"

	"github.com/jasoncolburne/cesrgo/common"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

type Number struct {
	matter
	number big.Int
}

func NewNumber(number *big.Int, hex *string, opts ...options.MatterOption) (*Number, error) {
	n := &Number{}

	config := options.MatterOptions{}
	for _, opt := range opts {
		opt(&config)
	}

	if config.Qb2 != nil {
		if number != nil || hex != nil || config.Raw != nil || config.Qb64 != nil || config.Qb64b != nil {
			return nil, fmt.Errorf("qb2 cannot be provided with other options")
		}

		err := NewMatter(
			n,
			options.WithQb2(*config.Qb2),
		)
		if err != nil {
			return nil, err
		}
	}

	if config.Qb64 != nil {
		if number != nil || hex != nil || config.Raw != nil || config.Qb2 != nil || config.Qb64b != nil {
			return nil, fmt.Errorf("qb2 cannot be provided with other options")
		}

		err := NewMatter(
			n,
			options.WithQb64(*config.Qb64),
		)
		if err != nil {
			return nil, err
		}
	}

	if config.Qb64b != nil {
		if number != nil || hex != nil || config.Raw != nil || config.Qb2 != nil || config.Qb64 != nil {
			return nil, fmt.Errorf("qb2 cannot be provided with other options")
		}

		err := NewMatter(
			n,
			options.WithQb64b(*config.Qb64b),
		)
		if err != nil {
			return nil, err
		}
	}

	// if we've processed qualified data, return now
	if config.Qb2 != nil || config.Qb64 != nil || config.Qb64b != nil {
		if !slices.Contains(codex.NumCodex, n.GetCode()) {
			return nil, fmt.Errorf("invalid number code")
		}

		n.number.SetBytes(n.GetRaw())
		if !common.LessThanMaxON(n.number) {
			return nil, fmt.Errorf("number is too large to be represented by 16 octets")
		}

		return n, nil
	}

	if config.Raw != nil {
		if !slices.Contains([]int{2, 5, 8, 11, 14, 17}, len(*config.Raw)) {
			return nil, fmt.Errorf("raw must be a valid size for typed numbers")
		}

		bigNum := big.Int{}
		bigNum.SetBytes(*config.Raw)

		n.number = bigNum
		if !common.LessThanMaxON(n.number) {
			return nil, fmt.Errorf("number is too large to be represented by 16 octets")
		}

		var code types.Code
		switch len(*config.Raw) {
		case 2:
			code = codex.Short
		case 5:
			code = codex.Tall
		case 8:
			code = codex.Big
		case 11:
			code = codex.Large
		case 14:
			code = codex.Great
		case 17:
			code = codex.Vast
		default:
			return nil, fmt.Errorf("raw must be a valid size for typed numbers")
		}

		err := NewMatter(
			n,
			options.WithCode(code),
			options.WithRaw(*config.Raw),
		)
		if err != nil {
			return nil, err
		}

		return n, nil
	}

	if hex != nil {
		if number != nil {
			return nil, fmt.Errorf("number and hex cannot both be provided")
		}

		bigNum := big.Int{}
		bigNum.SetString(*hex, 16)

		number = &bigNum
	}

	if number == nil {
		number = big.NewInt(0)
	}

	var code types.Code

	bitLen := number.BitLen()
	byteLen := (bitLen + 7) / 8

	if config.Code != nil {
		code = *config.Code

		if !slices.Contains(codex.NumCodex, code) {
			return nil, fmt.Errorf("invalid number code")
		}
	} else {
		if byteLen <= 2 {
			code = codex.Short
		} else if byteLen <= 5 {
			code = codex.Tall
		} else if byteLen <= 8 {
			code = codex.Big
		} else if byteLen <= 11 {
			code = codex.Large
		} else if byteLen <= 14 {
			code = codex.Great
		} else if byteLen <= 17 {
			code = codex.Vast
		} else {
			return nil, fmt.Errorf("number is too large to be represented in 17 bytes")
		}
	}

	var requiredBytes int
	switch code {
	case codex.Short:
		requiredBytes = 2
	case codex.Tall:
		requiredBytes = 5
	case codex.Big:
		requiredBytes = 8
	case codex.Large:
		requiredBytes = 11
	case codex.Great:
		requiredBytes = 14
	case codex.Vast:
		requiredBytes = 17
	}

	if byteLen > requiredBytes {
		return nil, fmt.Errorf("number is too large to be represented in %d bytes", requiredBytes)
	}

	n.number = *number

	raw := make([]byte, requiredBytes)
	n.number.FillBytes(raw)
	if !common.LessThanMaxON(n.number) {
		return nil, fmt.Errorf("number is too large to be represented by 16 octets")
	}

	err := NewMatter(
		n,
		options.WithCode(code),
		options.WithRaw(raw),
	)
	if err != nil {
		return nil, err
	}

	return n, nil
}

func (n *Number) Number() big.Int {
	return n.number
}

func (n *Number) Hex() string {
	var requiredChars int
	switch n.GetCode() {
	case codex.Short:
		requiredChars = 4
	case codex.Tall:
		requiredChars = 10
	case codex.Big:
		requiredChars = 16
	case codex.Large:
		requiredChars = 22
	case codex.Great:
		requiredChars = 28
	case codex.Vast:
		requiredChars = 34
	}

	hex := strings.ToUpper(n.number.Text(16))

	if len(hex) < requiredChars {
		hex = strings.Repeat("0", requiredChars-len(hex)) + hex
	}

	return hex
}
