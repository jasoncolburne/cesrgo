package test

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

func TestDigerCodes(t *testing.T) {
	var testCases = []struct {
		DigerCode types.Code
	}{
		{DigerCode: codex.Blake3_256},
		{DigerCode: codex.Blake3_512},
		{DigerCode: codex.Blake2b_256},
		{DigerCode: codex.Blake2b_512},
		{DigerCode: codex.Blake2s_256},
		{DigerCode: codex.SHA2_256},
		{DigerCode: codex.SHA2_512},
		{DigerCode: codex.SHA3_256},
		{DigerCode: codex.SHA3_512},
	}

	for _, testCase := range testCases {
		label := fmt.Sprintf("%s", testCase.DigerCode)
		t.Run(label, func(t *testing.T) {
			Diger, err := cesr.NewDiger([]byte{}, options.WithCode(testCase.DigerCode))
			if err != nil {
				t.Fatalf("failed to create Diger: %v", err)
			}

			if Diger.GetCode() != testCase.DigerCode {
				t.Fatalf("Diger code mismatch: %s != %s", Diger.GetCode(), testCase.DigerCode)
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

	for _, testCase := range testCases {
		label := fmt.Sprintf("%s[+]", testCase.Code)
		t.Run(label, func(t *testing.T) {
			diger, err := cesr.NewDiger([]byte{}, options.WithCode(testCase.Code))
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

	for _, testCase := range testCases {
		label := fmt.Sprintf("%s[-]", testCase.Code)
		t.Run(label, func(t *testing.T) {
			diger, err := cesr.NewDiger([]byte{}, options.WithCode(testCase.Code))
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

func TestDigerRoundTrip(t *testing.T) {
	raw := [32]byte{}
	_, err := rand.Read(raw[:])
	if err != nil {
		t.Fatalf("failed to read random bytes: %v", err)
	}

	diger, err := cesr.NewDiger(nil, options.WithCode(codex.Blake3_256), options.WithRaw(raw[:]))
	if err != nil {
		t.Fatalf("failed to create diger: %v", err)
	}

	qb2, err := diger.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Diger, err := cesr.NewDiger(nil, options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create diger from qb2: %v", err)
	}

	qb64, err := qb2Diger.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Diger, err := cesr.NewDiger(nil, options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create diger from qb64: %v", err)
	}

	qb64b, err := qb64Diger.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bDiger, err := cesr.NewDiger(nil, options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create diger from qb64b: %v", err)
	}

	qb64bRaw := qb64bDiger.GetRaw()
	if !bytes.Equal(qb64bRaw, raw[:]) {
		t.Fatalf("qb64b raw mismatch: %x != %x", qb64bRaw, raw[:])
	}
}
