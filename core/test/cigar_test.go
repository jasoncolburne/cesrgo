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

	for _, testCase := range testCases {
		label := fmt.Sprintf("%s->%s", testCase.SignerCode, testCase.CigarCode)
		t.Run(label, func(t *testing.T) {
			signer, err := cesr.NewSigner(true, options.WithCode(testCase.SignerCode))
			if err != nil {
				t.Fatalf("failed to create signer: %v", err)
			}

			cigar, err := signer.SignUnindexed([]byte{})
			if err != nil {
				t.Fatalf("failed to sign unindexed: %v", err)
			}

			if cigar.GetCode() != testCase.CigarCode {
				t.Fatalf("cigar code mismatch: %s != %s", cigar.GetCode(), testCase.CigarCode)
			}
		})
	}
}

func TestCigarRoundTrip(t *testing.T) {
	raw := [64]byte{}
	_, err := rand.Read(raw[:])
	if err != nil {
		t.Fatalf("failed to read random bytes: %v", err)
	}

	cigar, err := cesr.NewCigar(nil, options.WithCode(codex.Ed25519_Sig), options.WithRaw(raw[:]))
	if err != nil {
		t.Fatalf("failed to create cigar: %v", err)
	}

	qb2, err := cigar.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Cigar, err := cesr.NewCigar(nil, options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create cigar from qb2: %v", err)
	}

	qb64, err := qb2Cigar.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Cigar, err := cesr.NewCigar(nil, options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create cigar from qb64: %v", err)
	}

	qb64b, err := qb64Cigar.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bCigar, err := cesr.NewCigar(nil, options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create cigar from qb64b: %v", err)
	}

	qb64bRaw := qb64bCigar.GetRaw()
	if !bytes.Equal(qb64bRaw, raw[:]) {
		t.Fatalf("qb64b raw mismatch: %x != %x", qb64bRaw, raw[:])
	}
}
