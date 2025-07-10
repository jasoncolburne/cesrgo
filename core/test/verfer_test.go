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

func TestVerferCodesAndSizes(t *testing.T) {
	var testCases = []struct {
		Code types.Code
		Size types.Size
	}{
		{
			Code: codex.Ed25519,
			Size: 32,
		},
		{
			Code: codex.Ed25519N,
			Size: 32,
		},
		{
			Code: codex.ECDSA_256k1,
			Size: 33,
		},
		{
			Code: codex.ECDSA_256k1N,
			Size: 33,
		},
		{
			Code: codex.ECDSA_256r1,
			Size: 33,
		},
		{
			Code: codex.ECDSA_256r1N,
			Size: 33,
		},
	}

	for _, testCase := range testCases {
		label := fmt.Sprintf("%s[+]", testCase.Code)
		t.Run(label, func(t *testing.T) {
			raw := make(types.Raw, testCase.Size)
			verfer, err := cesr.NewVerfer(options.WithCode(testCase.Code), options.WithRaw(raw))
			if err != nil {
				t.Fatalf("failed to create verfer: %v", err)
			}

			if verfer.GetCode() != testCase.Code {
				t.Fatalf("verfer code mismatch: %s != %s", verfer.GetCode(), testCase.Code)
			}
		})
	}

	for _, testCase := range testCases {
		label := fmt.Sprintf("%s[-]", testCase.Code)
		t.Run(label, func(t *testing.T) {
			raw := make(types.Raw, testCase.Size+1)
			_, err := cesr.NewVerfer(options.WithCode(testCase.Code), options.WithRaw(raw))
			if err == nil {
				t.Fatalf("created verfer with invalid raw")
			}
		})
	}
}

func TestVerferVerification(t *testing.T) {
	testCases := []struct {
		SignerCode types.Code
	}{
		{
			SignerCode: codex.Ed25519_Seed,
		},
		{
			SignerCode: codex.ECDSA_256k1_Seed,
		},
		{
			SignerCode: codex.ECDSA_256r1_Seed,
		},
	}

	for _, testCase := range testCases {
		label := fmt.Sprintf("%s[+]", testCase.SignerCode)
		t.Run(label, func(t *testing.T) {
			signer, err := cesr.NewSigner(false, options.WithCode(testCase.SignerCode))
			if err != nil {
				t.Fatalf("failed to create signer: %v", err)
			}

			cigar, err := signer.SignUnindexed([]byte{})
			if err != nil {
				t.Fatalf("failed to sign: %v", err)
			}

			verfer := signer.GetVerfer()
			if verfer == nil {
				t.Fatalf("verfer is nil")
			}

			verified, err := verfer.Verify(cigar.GetRaw(), []byte{})
			if err != nil {
				t.Fatalf("invalid signature: %v", err)
			}

			if !verified {
				t.Fatalf("invalid signature")
			}
		})
	}

	for _, testCase := range testCases {
		label := fmt.Sprintf("%s[-]", testCase.SignerCode)
		t.Run(label, func(t *testing.T) {
			signer, err := cesr.NewSigner(false, options.WithCode(testCase.SignerCode))
			if err != nil {
				t.Fatalf("failed to create signer: %v", err)
			}

			cigar, err := signer.SignUnindexed([]byte{})
			if err != nil {
				t.Fatalf("failed to sign: %v", err)
			}

			verfer := signer.GetVerfer()
			if verfer == nil {
				t.Fatalf("verfer is nil")
			}

			verified, err := verfer.Verify(cigar.GetRaw(), []byte{1})
			if err == nil {
				t.Fatalf("unexpected valid signature")
			}

			if verified {
				t.Fatalf("unexpected valid signature")
			}
		})
	}
}

func TestVerferRoundTrip(t *testing.T) {
	raw := [32]byte{}
	_, err := rand.Read(raw[:])
	if err != nil {
		t.Fatalf("failed to read random bytes: %v", err)
	}

	verfer, err := cesr.NewVerfer(options.WithCode(codex.Ed25519), options.WithRaw(raw[:]))
	if err != nil {
		t.Fatalf("failed to create verfer: %v", err)
	}

	qb2, err := verfer.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Verfer, err := cesr.NewVerfer(options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create verfer: %v", err)
	}

	qb64, err := qb2Verfer.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Verfer, err := cesr.NewVerfer(options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create verfer: %v", err)
	}

	qb64b, err := qb64Verfer.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bVerfer, err := cesr.NewVerfer(options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create verfer: %v", err)
	}

	qb64bRaw := qb64bVerfer.GetRaw()
	if !bytes.Equal(qb64bRaw, raw[:]) {
		t.Fatalf("qb64b raw mismatch: %x != %x", qb64bRaw, raw[:])
	}
}
