package test

import (
	"fmt"
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

func TestDigerCodesAndSizes(t *testing.T) {
	var testCases = []struct {
		DigerCode types.Code
		DigerSize types.Size
	}{
		{
			DigerCode: codex.Blake3_256,
			DigerSize: 32,
		},
		{
			DigerCode: codex.Blake3_512,
			DigerSize: 64,
		},
		{
			DigerCode: codex.Blake2b_256,
			DigerSize: 32,
		},
		{
			DigerCode: codex.Blake2b_512,
			DigerSize: 64,
		},
		{
			DigerCode: codex.Blake2s_256,
			DigerSize: 32,
		},
		{
			DigerCode: codex.SHA2_256,
			DigerSize: 32,
		},
		{
			DigerCode: codex.SHA2_512,
			DigerSize: 64,
		},
		{
			DigerCode: codex.SHA3_256,
			DigerSize: 32,
		},
		{
			DigerCode: codex.SHA3_512,
			DigerSize: 64,
		},
	}

	for _, testVector := range testCases {
		label := fmt.Sprintf("%s", testVector.DigerCode)
		t.Run(label, func(t *testing.T) {
			Diger, err := cesr.NewDiger([]byte{}, options.WithCode(testVector.DigerCode))
			if err != nil {
				t.Fatalf("failed to create Diger: %v", err)
			}

			if Diger.GetCode() != testVector.DigerCode {
				t.Fatalf("Diger code mismatch: %s != %s", Diger.GetCode(), testVector.DigerCode)
			}

			if Diger.GetSize() != testVector.DigerSize {
				t.Fatalf("Diger size mismatch: %d != %d", Diger.GetSize(), testVector.DigerSize)
			}
		})
	}
}

func TestDigerVerification(t *testing.T) {
	testCases := []struct {
		Code types.Code
	}{
		{
			Code: codex.Blake3_256,
		},
		{
			Code: codex.Blake3_512,
		},
		{
			Code: codex.Blake2b_256,
		},
		{
			Code: codex.Blake2b_512,
		},
		{
			Code: codex.Blake2s_256,
		},
		{
			Code: codex.SHA2_256,
		},
		{
			Code: codex.SHA2_512,
		},
		{
			Code: codex.SHA3_256,
		},
		{
			Code: codex.SHA3_512,
		},
	}

	for _, testVector := range testCases {
		label := fmt.Sprintf("%s[+]", testVector.Code)
		t.Run(label, func(t *testing.T) {
			diger, err := cesr.NewDiger([]byte{}, options.WithCode(testVector.Code))
			if err != nil {
				t.Fatalf("failed to create Diger: %v", err)
			}

			verified, err := diger.Verify([]byte{})
			if err != nil {
				t.Fatalf("failed to verify: %v", err)
			}

			if !verified {
				t.Fatalf("invalid digest")
			}
		})
	}

	for _, testVector := range testCases {
		label := fmt.Sprintf("%s[-]", testVector.Code)
		t.Run(label, func(t *testing.T) {
			diger, err := cesr.NewDiger([]byte{}, options.WithCode(testVector.Code))
			if err != nil {
				t.Fatalf("failed to create Diger: %v", err)
			}

			verified, err := diger.Verify([]byte{1})
			if verified {
				t.Fatalf("unexpected valid digest")
			}
		})
	}
}
