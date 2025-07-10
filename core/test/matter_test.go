package test

import (
	"slices"
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

func TestMatterInvalidOptionCombinations(t *testing.T) {
	code := codex.Blake3_256
	raw := types.Raw{}
	qb2 := types.Qb2{}
	qb64 := types.Qb64("")
	qb64b := types.Qb64b{}

	testCases := []options.MatterOptions{
		// missing data
		{
			Code:  &code,
			Raw:   nil,
			Qb2:   nil,
			Qb64:  nil,
			Qb64b: nil,
		},
		{
			Code:  nil,
			Raw:   &raw,
			Qb2:   nil,
			Qb64:  nil,
			Qb64b: nil,
		},

		// conflicting inputs
		{
			Code:  &code,
			Raw:   &raw,
			Qb2:   &qb2,
			Qb64:  nil,
			Qb64b: nil,
		},
		{
			Code:  &code,
			Raw:   &raw,
			Qb2:   nil,
			Qb64:  &qb64,
			Qb64b: nil,
		},
		{
			Code:  &code,
			Raw:   &raw,
			Qb2:   nil,
			Qb64:  nil,
			Qb64b: &qb64b,
		},
		{
			Code:  nil,
			Raw:   nil,
			Qb2:   &qb2,
			Qb64:  &qb64,
			Qb64b: nil,
		},
		{
			Code:  nil,
			Raw:   nil,
			Qb2:   &qb2,
			Qb64:  nil,
			Qb64b: &qb64b,
		},
		{
			Code:  nil,
			Raw:   nil,
			Qb2:   nil,
			Qb64:  &qb64,
			Qb64b: &qb64b,
		},
	}

	for _, testCase := range testCases {
		m := &cesr.UndifferentiatedMatter{}

		args := []options.MatterOption{}

		if testCase.Code != nil {
			args = append(args, options.WithCode(*testCase.Code))
		}

		if testCase.Raw != nil {
			args = append(args, options.WithRaw(*testCase.Raw))
		}

		if testCase.Qb2 != nil {
			args = append(args, options.WithQb2(*testCase.Qb2))
		}

		if testCase.Qb64 != nil {
			args = append(args, options.WithQb64(*testCase.Qb64))
		}

		if testCase.Qb64b != nil {
			args = append(args, options.WithQb64b(*testCase.Qb64b))
		}

		if err := cesr.NewMatter(m, args...); err == nil {
			t.Fatalf("creation did not fail")
		}
	}
}

func TestMatterRoundTrip(t *testing.T) {
	codeAndRawMatter := &cesr.UndifferentiatedMatter{}
	qb2Matter := &cesr.UndifferentiatedMatter{}
	qb64Matter := &cesr.UndifferentiatedMatter{}
	qb64bMatter := &cesr.UndifferentiatedMatter{}

	// from crypto.rand.Read
	raw := types.Raw("\xeaz5\x17\xfeQ!yʹH\x9b\x1aօXO\x1a\x1aq\x17\xd7r_$9\xfaҺණ")

	if err := cesr.NewMatter(
		codeAndRawMatter,
		options.WithCode(codex.Blake3_256),
		options.WithRaw(raw),
	); err != nil {
		t.Fatalf("failed to create matter: %v", err)
	}

	qb2, err := codeAndRawMatter.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	if err := cesr.NewMatter(
		qb2Matter,
		options.WithQb2(qb2),
	); err != nil {
		t.Fatalf("failed to create matter: %v", err)
	}

	qb64, err := qb2Matter.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	if err := cesr.NewMatter(
		qb64Matter,
		options.WithQb64(qb64),
	); err != nil {
		t.Fatalf("failed to create matter: %v", err)
	}

	qb64b, err := qb64Matter.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	if err := cesr.NewMatter(
		qb64bMatter,
		options.WithQb64b(qb64b),
	); err != nil {
		t.Fatalf("failed to create matter: %v", err)
	}

	if slices.Compare(raw, qb64bMatter.GetRaw()) != 0 {
		t.Fatalf("raw mismatch after round trip: %x != %x", raw, qb64bMatter.GetRaw())
	}
}
