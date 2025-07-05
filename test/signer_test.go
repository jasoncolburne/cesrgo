package test

import (
	"fmt"
	"testing"

	"github.com/jasoncolburne/cesrgo"
	mdex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
)

func TestSignerCodesAndSizes(t *testing.T) {
	testVectors := []struct {
		Code types.Code
		Size types.Size
	}{
		{
			Code: mdex.Ed25519_Seed,
			Size: 32,
		},
		{
			Code: mdex.ECDSA_256k1_Seed,
			Size: 32,
		},
		{
			Code: mdex.ECDSA_256r1_Seed,
			Size: 32,
		},
	}

	for _, testVector := range testVectors {
		label := fmt.Sprintf("%s", testVector.Code)
		t.Run(label, func(t *testing.T) {
			signer, err := cesrgo.NewSigner(true, options.WithCode(testVector.Code))
			if err != nil {
				t.Fatalf("failed to create signer: %v", err)
			}

			if signer.GetCode() != testVector.Code {
				t.Fatalf("signer code mismatch: %s != %s", signer.GetCode(), testVector.Code)
			}

			if signer.GetSize() != testVector.Size {
				t.Fatalf("signer size mismatch: %d != %d", signer.GetSize(), testVector.Size)
			}
		})
	}
}

func TestSignerDeterminism(t *testing.T) {
	testVectors := []struct {
		Code          types.Code
		Raw           []byte
		NullCigarQb64 types.Qb64
	}{
		{
			Code:          mdex.Ed25519_Seed,
			Raw:           []byte{73, 101, 58, 239, 142, 157, 136, 49, 137, 118, 128, 4, 230, 235, 118, 214, 151, 48, 147, 53, 222, 192, 90, 217, 81, 203, 48, 204, 29, 249, 88, 187},
			NullCigarQb64: types.Qb64("0BANAYU1Rk-fC1Xc26rbjfbh4Wwm0Ghnhe_oJ3Gc8ka10DCAUwTd_iscBGzX0ppRmi2jeizchGwJiO0jli8zbckF"),
		},
		{
			Code:          mdex.ECDSA_256k1_Seed,
			Raw:           []byte{180, 254, 124, 246, 186, 56, 94, 128, 161, 224, 100, 126, 79, 97, 198, 211, 228, 84, 105, 175, 28, 77, 33, 245, 29, 156, 241, 90, 87, 104, 189, 156},
			NullCigarQb64: types.Qb64("0CAAeEWPU1gVoz3p804jwDbJiPDveINH3oNf3oPjJydiO2nM4V0qpB6tYgQ8LAVnItykoTC9xXJlHwNTjKVWBHNq"),
		},
		{
			Code:          mdex.ECDSA_256r1_Seed,
			Raw:           []byte{144, 11, 61, 132, 63, 222, 90, 96, 110, 254, 120, 210, 211, 210, 189, 21, 4, 3, 68, 77, 226, 196, 142, 2, 42, 115, 37, 125, 244, 132, 176, 2},
			NullCigarQb64: types.Qb64("0IDp2liCC2HmkE3TaemOVpL2snwPoCYON58LIDbMOTKYPhDxPW6181HZLdyt6f03pcrtb-f38KLQLjHoCbqR7NIA"),
		},
	}

	for _, testVector := range testVectors {
		label := fmt.Sprintf("%s", testVector.Code)
		t.Run(label, func(t *testing.T) {
			signer, err := cesrgo.NewSigner(true, options.WithCode(testVector.Code), options.WithRaw(testVector.Raw))
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

			if qb64 != testVector.NullCigarQb64 {
				t.Fatalf("cigar qb64 mismatch: %s != %s", qb64, testVector.NullCigarQb64)
			}
		})
	}

}

func TestSignerUnindexedSignatureCreation(t *testing.T) {
	testVectors := []struct {
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

	for _, testVector := range testVectors {
		label := fmt.Sprintf("%s", testVector.Code)
		t.Run(label, func(t *testing.T) {
			signer, err := cesrgo.NewSigner(true, options.WithCode(testVector.Code))
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
	testVectors := []struct {
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

	for _, testVector := range testVectors {
		label := fmt.Sprintf("%s", testVector.Code)
		t.Run(label, func(t *testing.T) {
			signer, err := cesrgo.NewSigner(true, options.WithCode(testVector.Code))
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
