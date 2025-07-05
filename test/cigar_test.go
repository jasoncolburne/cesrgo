package test

import (
	"fmt"
	"testing"

	"github.com/jasoncolburne/cesrgo"
	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
)

func TestCigarCodesAndSizes(t *testing.T) {
	var testVectors = []struct {
		SignerCode types.Code
		CigarCode  types.Code
		CigarSize  types.Size
	}{
		{
			SignerCode: codex.Ed25519_Seed,
			CigarCode:  codex.Ed25519_Sig,
			CigarSize:  64,
		},
		{
			SignerCode: codex.ECDSA_256k1_Seed,
			CigarCode:  codex.ECDSA_256k1_Sig,
			CigarSize:  64,
		},
		{
			SignerCode: codex.ECDSA_256r1_Seed,
			CigarCode:  codex.ECDSA_256r1_Sig,
			CigarSize:  64,
		},
	}

	for _, testVector := range testVectors {
		label := fmt.Sprintf("%s->%s", testVector.SignerCode, testVector.CigarCode)
		t.Run(label, func(t *testing.T) {
			signer, err := cesrgo.NewSigner(true, options.WithCode(testVector.SignerCode))
			if err != nil {
				t.Fatalf("failed to create signer: %v", err)
			}

			cigar, err := signer.SignUnindexed([]byte{})
			if err != nil {
				t.Fatalf("failed to sign unindexed: %v", err)
			}

			if cigar.GetCode() != testVector.CigarCode {
				t.Fatalf("cigar code mismatch: %s != %s", cigar.GetCode(), testVector.CigarCode)
			}

			if cigar.GetSize() != testVector.CigarSize {
				t.Fatalf("cigar size mismatch: %d != %d", cigar.GetSize(), testVector.CigarSize)
			}
		})
	}
}
