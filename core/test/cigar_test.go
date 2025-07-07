package test

import (
	"fmt"
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

func TestCigarCodesAndSizes(t *testing.T) {
	var testCases = []struct {
		SignerCode types.Code
		CigarCode  types.Code
	}{
		{
			SignerCode: codex.Ed25519_Seed,
			CigarCode:  codex.Ed25519_Sig,
		},
		{
			SignerCode: codex.ECDSA_256k1_Seed,
			CigarCode:  codex.ECDSA_256k1_Sig,
		},
		{
			SignerCode: codex.ECDSA_256r1_Seed,
			CigarCode:  codex.ECDSA_256r1_Sig,
		},
	}

	for _, testVector := range testCases {
		label := fmt.Sprintf("%s->%s", testVector.SignerCode, testVector.CigarCode)
		t.Run(label, func(t *testing.T) {
			signer, err := cesr.NewSigner(true, options.WithCode(testVector.SignerCode))
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
		})
	}
}
