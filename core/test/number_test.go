package test

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"slices"
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
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
		number, err := cesr.NewNumber(test.number, nil)
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
	number, err := cesr.NewNumber(nil, nil)
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

	number, err = cesr.NewNumber(nil, nil, options.WithCode(codex.Vast))
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
	number, err := cesr.NewNumber(nil, &hex)
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
	number, err = cesr.NewNumber(nil, &hex)
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

func TestNumberRoundTrip(t *testing.T) {
	raw := [15]byte{}
	_, err := rand.Read(raw[:])
	if err != nil {
		t.Fatalf("failed to read random bytes: %v", err)
	}

	number, err := cesr.NewNumber(nil, nil, options.WithRaw(raw[:]))
	if err != nil {
		t.Fatalf("failed to create number: %v", err)
	}

	qb2, err := number.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Number, err := cesr.NewNumber(nil, nil, options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create number from qb2: %v", err)
	}

	qb64, err := qb2Number.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Number, err := cesr.NewNumber(nil, nil, options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create number from qb64: %v", err)
	}

	qb64b, err := qb64Number.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bNumber, err := cesr.NewNumber(nil, nil, options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create number from qb64b: %v", err)
	}

	padded := make([]byte, 17)
	copy(padded[2:], raw[:])
	qb64bRaw := qb64bNumber.GetRaw()
	if !bytes.Equal(qb64bRaw, padded) {
		t.Fatalf("qb64b raw mismatch: %x != %x", qb64bRaw, padded)
	}
}
