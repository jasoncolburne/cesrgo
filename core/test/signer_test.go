package test

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	mdex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

func TestSignerCodesAndSizes(t *testing.T) {
	testCases := []struct {
		Code types.Code
	}{
		{Code: mdex.Ed25519_Seed},
		{Code: mdex.ECDSA_256k1_Seed},
		{Code: mdex.ECDSA_256r1_Seed},
	}

	for _, testCase := range testCases {
		label := fmt.Sprintf("%s", testCase.Code)
		t.Run(label, func(t *testing.T) {
			signer, err := cesr.NewSigner(true, options.WithCode(testCase.Code))
			if err != nil {
				t.Fatalf("failed to create signer: %v", err)
			}

			if signer.GetCode() != testCase.Code {
				t.Fatalf("signer code mismatch: %s != %s", signer.GetCode(), testCase.Code)
			}
		})
	}
}

func TestSignerDeterminism(t *testing.T) {
	testCases := []struct {
		Code          types.Code
		Raw           []byte
		NullCigarQb64 types.Qb64
	}{
		{
			Code:          mdex.Ed25519_Seed,
			Raw:           types.Raw("Ie:\uf39d\x881\x89v\x80\x04\xe6\xebv֗0\x935\xde\xc0Z\xd9Q\xcb0\xcc\x1d\xf9X\xbb"),
			NullCigarQb64: types.Qb64("0BANAYU1Rk-fC1Xc26rbjfbh4Wwm0Ghnhe_oJ3Gc8ka10DCAUwTd_iscBGzX0ppRmi2jeizchGwJiO0jli8zbckF"),
		},
		{
			Code:          mdex.ECDSA_256k1_Seed,
			Raw:           types.Raw("\xb4\xfe|\xf6\xba8^\x80\xa1\xe0d~Oa\xc6\xd3\xe4Ti\xaf\x1cM!\xf5\x1d\x9c\xf1ZWh\xbd\x9c"),
			NullCigarQb64: types.Qb64("0CAAeEWPU1gVoz3p804jwDbJiPDveINH3oNf3oPjJydiO2nM4V0qpB6tYgQ8LAVnItykoTC9xXJlHwNTjKVWBHNq"),
		},
		{
			Code:          mdex.ECDSA_256r1_Seed,
			Raw:           types.Raw("\x90\v=\x84?\xdeZ`n\xfex\xd2\xd3ҽ\x15\x04\x03DM\xe2Ď\x02*s%}\xf4\x84\xb0\x02"),
			NullCigarQb64: types.Qb64("0IDp2liCC2HmkE3TaemOVpL2snwPoCYON58LIDbMOTKYPhDxPW6181HZLdyt6f03pcrtb-f38KLQLjHoCbqR7NIA"),
		},
	}

	for _, testCase := range testCases {
		label := fmt.Sprintf("%s", testCase.Code)
		t.Run(label, func(t *testing.T) {
			signer, err := cesr.NewSigner(true, options.WithCode(testCase.Code), options.WithRaw(testCase.Raw))
			if err != nil {
				t.Fatalf("failed to create signer: %v", err)
			}

			cigar, err := signer.SignUnindexed([]byte{})
			if err != nil {
				t.Fatalf("failed to sign unindexed: %v", err)
			}

			qb64, err := cigar.Qb64()
			if err != nil {
				t.Fatalf("failed to get qb64: %v", err)
			}

			if qb64 != testCase.NullCigarQb64 {
				t.Fatalf("cigar qb64 mismatch: %s != %s", qb64, testCase.NullCigarQb64)
			}
		})
	}

}

func TestSignerUnindexedSignatureCreation(t *testing.T) {
	testCases := []struct {
		Code types.Code
	}{
		{
			Code: mdex.Ed25519_Seed,
		},
		{
			Code: mdex.ECDSA_256k1_Seed,
		},
		{
			Code: mdex.ECDSA_256r1_Seed,
		},
	}

	for _, testCase := range testCases {
		label := fmt.Sprintf("%s", testCase.Code)
		t.Run(label, func(t *testing.T) {
			signer, err := cesr.NewSigner(true, options.WithCode(testCase.Code))
			if err != nil {
				t.Fatalf("failed to create signer: %v", err)
			}

			_, err = signer.SignUnindexed([]byte{})
			if err != nil {
				t.Fatalf("failed to sign unindexed: %v", err)
			}
		})
	}
}

func TestSignerIndexedSignatureCreation(t *testing.T) {
	testCases := []struct {
		Code types.Code
	}{
		{
			Code: mdex.Ed25519_Seed,
		},
		{
			Code: mdex.ECDSA_256k1_Seed,
		},
		{
			Code: mdex.ECDSA_256r1_Seed,
		},
	}

	for _, testCase := range testCases {
		label := fmt.Sprintf("%s", testCase.Code)
		t.Run(label, func(t *testing.T) {
			signer, err := cesr.NewSigner(true, options.WithCode(testCase.Code))
			if err != nil {
				t.Fatalf("failed to create signer: %v", err)
			}

			_, err = signer.SignIndexed([]byte{}, false, types.Index(0), nil)
			if err != nil {
				t.Fatalf("failed to sign indexed: %v", err)
			}
		})
	}
}

func TestSignerRoundTrip(t *testing.T) {
	raw := [32]byte{}
	_, err := rand.Read(raw[:])
	if err != nil {
		t.Fatalf("failed to read random bytes: %v", err)
	}

	codeSigner, err := cesr.NewSigner(true, options.WithCode(mdex.Ed25519_Seed), options.WithRaw(raw[:]))
	if err != nil {
		t.Fatalf("failed to create signer: %v", err)
	}

	qb2, err := codeSigner.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Signer, err := cesr.NewSigner(true, options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create signer from qb2: %v", err)
	}

	qb64, err := qb2Signer.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Signer, err := cesr.NewSigner(true, options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create signer from qb64: %v", err)
	}

	qb64b, err := qb64Signer.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bSigner, err := cesr.NewSigner(true, options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create signer from qb64b: %v", err)
	}

	qb64bRaw := qb64bSigner.GetRaw()
	if !bytes.Equal(qb64bRaw, raw[:]) {
		t.Fatalf("qb64b raw mismatch: %x != %x", qb64bRaw, raw[:])
	}
}
