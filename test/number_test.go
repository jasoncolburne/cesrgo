package test

import (
	"math/big"
	"slices"
	"testing"

	"github.com/jasoncolburne/cesrgo"
	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
)

func TestNumberCodesAndSizes(t *testing.T) {
	shortNum := big.Int{}
	tallNum := big.Int{}
	bigNum := big.Int{}
	largeNum := big.Int{}
	greatNum := big.Int{}
	vastNum := big.Int{}

	shortNum.SetString("FFFF", 16)
	tallNum.SetString("FFFFFFFFFF", 16)
	bigNum.SetString("FFFFFFFFFFFFFFFF", 16)
	largeNum.SetString("FFFFFFFFFFFFFFFFFFFFFF", 16)
	greatNum.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
	vastNum.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)

	testCases := []struct {
		code   types.Code
		number *big.Int
	}{
		{code: codex.Short, number: &shortNum},
		{code: codex.Tall, number: &tallNum},
		{code: codex.Big, number: &bigNum},
		{code: codex.Large, number: &largeNum},
		{code: codex.Great, number: &greatNum},
		{code: codex.Vast, number: &vastNum},
	}

	for _, test := range testCases {
		number, err := cesrgo.NewNumber(test.number, nil)
		if err != nil {
			t.Fatalf("failed to create number: %v", err)
		}

		if number.GetCode() != test.code {
			t.Fatalf("expected code %v, got %v", test.code, number.GetCode())
		}

		// byteLen/rawLen account for padding, but make the test a bit unsafe.
		// this only happens for one case, and the delta is only 1, so we can ensure
		// the delta is in a narrow window to add a bit of safety back
		byteLen := (test.number.BitLen() + 7) / 8
		rawLen := len(number.GetRaw())

		// it's okay for the vast case to have 1 leading pad byte, not okay for others
		if (rawLen > byteLen && test.code != codex.Vast) || rawLen-1 > byteLen || byteLen > rawLen {
			t.Fatalf("raw does not match byte length. raw=%v, bytes=%v", number.GetRaw(), test.number.Bytes())
		}

		if slices.Compare(number.GetRaw()[rawLen-byteLen:], test.number.Bytes()) != 0 {
			t.Fatalf("expected raw %v, got %v", test.number.Bytes(), number.GetRaw())
		}
	}
}

func TestNumberDefaults(t *testing.T) {
	number, err := cesrgo.NewNumber(nil, nil)
	if err != nil {
		t.Fatalf("failed to create number: %v", err)
	}

	if number.GetCode() != codex.Short {
		t.Fatalf("expected code %v, got %v", codex.Short, number.GetCode())
	}

	bigNum := number.Number()
	if bigNum.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("expected number to be 0, got %v", number.Number())
	}

	number, err = cesrgo.NewNumber(nil, nil, options.WithCode(codex.Vast))
	if err != nil {
		t.Fatalf("failed to create number: %v", err)
	}

	if number.GetCode() != codex.Vast {
		t.Fatalf("expected code %v, got %v", codex.Vast, number.GetCode())
	}

	bigNum = number.Number()
	if bigNum.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("expected number to be 0, got %v", number.Number())
	}
}

func TestNumberHexEncoding(t *testing.T) {
	hex := "FFFF"
	number, err := cesrgo.NewNumber(nil, &hex)
	if err != nil {
		t.Fatalf("failed to create number: %v", err)
	}

	if number.GetCode() != codex.Short {
		t.Fatalf("expected code %v, got %v", codex.Short, number.GetCode())
	}

	if number.Hex() != hex {
		t.Fatalf("expected hex %v, got %v", hex, number.Hex())
	}

	// should this bail if we create an unpadded number like this?
	hex = "0123456789ABCDEFFEDBCA9876543210"
	number, err = cesrgo.NewNumber(nil, &hex)
	if err != nil {
		t.Fatalf("failed to create number: %v", err)
	}

	if number.GetCode() != codex.Vast {
		t.Fatalf("expected code %v, got %v", codex.Vast, number.GetCode())
	}

	if number.Hex() != "00"+hex {
		t.Fatalf("expected hex %v, got %v", "00"+hex, number.Hex())
	}
}
