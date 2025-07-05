package test

import (
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
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
