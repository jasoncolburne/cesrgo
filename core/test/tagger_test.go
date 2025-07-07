package test

import (
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

func TestTaggerCodesAndTags(t *testing.T) {
	testCases := []struct {
		tag  string
		code types.Code
	}{
		{
			tag:  "0",
			code: codex.Tag1,
		},
		{
			tag:  "00",
			code: codex.Tag2,
		},
		{
			tag:  "000",
			code: codex.Tag3,
		},
		{
			tag:  "0000",
			code: codex.Tag4,
		},
		{
			tag:  "00000",
			code: codex.Tag5,
		},
		{
			tag:  "000000",
			code: codex.Tag6,
		},
		{
			tag:  "0000000",
			code: codex.Tag7,
		},
		{
			tag:  "00000000",
			code: codex.Tag8,
		},
		{
			tag:  "000000000",
			code: codex.Tag9,
		},
		{
			tag:  "0000000000",
			code: codex.Tag10,
		},
		{
			tag:  "00000000000",
			code: codex.Tag11,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.tag, func(t *testing.T) {
			tagger, err := cesr.NewTagger(&testCase.tag)
			if err != nil {
				t.Fatalf("error creating tagger: %v", err)
			}

			if tagger.GetCode() != testCase.code {
				t.Fatalf("expected code %s, got %s", testCase.code, tagger.GetCode())
			}
		})
	}
}

func TestTaggerRoundTrip(t *testing.T) {
	tag := "fakelongtag"
	tagTagger, err := cesr.NewTagger(&tag)
	if err != nil {
		t.Fatalf("failed to create tagger: %v", err)
	}

	qb2, err := tagTagger.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Tagger, err := cesr.NewTagger(nil, options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create tagger from qb2: %v", err)
	}

	qb64, err := qb2Tagger.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Tagger, err := cesr.NewTagger(nil, options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create tagger from qb64: %v", err)
	}

	qb64b, err := qb64Tagger.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bTagger, err := cesr.NewTagger(nil, options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create tagger from qb64b: %v", err)
	}

	qb64bTag, err := qb64bTagger.Tag()
	if err != nil {
		t.Fatalf("failed to get tag: %v", err)
	}

	if qb64bTag != tag {
		t.Fatalf("qb64b tag mismatch: %s != %s", qb64bTag, tag)
	}
}
