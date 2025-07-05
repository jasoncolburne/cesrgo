package test

import (
	"fmt"
	"testing"

	"github.com/jasoncolburne/cesrgo"
	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
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

	for _, testVector := range testCases {
		label := fmt.Sprintf("%s[+]", testVector.Code)
		t.Run(label, func(t *testing.T) {
			raw := make(types.Raw, testVector.Size)
			verfer, err := cesrgo.NewVerfer(options.WithCode(testVector.Code), options.WithRaw(raw))
			if err != nil {
				t.Fatalf("failed to create verfer: %v", err)
			}

			if verfer.GetCode() != testVector.Code {
				t.Fatalf("verfer code mismatch: %s != %s", verfer.GetCode(), testVector.Code)
			}

			if verfer.GetSize() != testVector.Size {
				t.Fatalf("verfer size mismatch: %d != %d", verfer.GetSize(), testVector.Size)
			}
		})
	}

	for _, testVector := range testCases {
		label := fmt.Sprintf("%s[-]", testVector.Code)
		t.Run(label, func(t *testing.T) {
			raw := make(types.Raw, testVector.Size+1)
			_, err := cesrgo.NewVerfer(options.WithCode(testVector.Code), options.WithRaw(raw))
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

	for _, testVector := range testCases {
		label := fmt.Sprintf("%s[+]", testVector.SignerCode)
		t.Run(label, func(t *testing.T) {
			signer, err := cesrgo.NewSigner(false, options.WithCode(testVector.SignerCode))
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

	for _, testVector := range testCases {
		label := fmt.Sprintf("%s[-]", testVector.SignerCode)
		t.Run(label, func(t *testing.T) {
			signer, err := cesrgo.NewSigner(false, options.WithCode(testVector.SignerCode))
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
