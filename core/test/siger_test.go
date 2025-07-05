package test

import (
	"fmt"
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	idex "github.com/jasoncolburne/cesrgo/core/indexer"
	mdex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

func TestSigerCodesAndIndicies(t *testing.T) {
	_58 := types.Ondex(58)
	_73 := types.Ondex(73)
	_81 := types.Ondex(81)

	testCases := []struct {
		SignerCode types.Code
		SigerCode  types.Code
		Only       bool
		Index      types.Index
		Ondex      *types.Ondex
	}{
		{
			SignerCode: mdex.Ed25519_Seed,
			SigerCode:  idex.Ed25519,
			Only:       false,
			Index:      37,
		},
		{
			SignerCode: mdex.ECDSA_256k1_Seed,
			SigerCode:  idex.ECDSA_256k1,
			Only:       false,
			Index:      23,
		},
		{
			SignerCode: mdex.ECDSA_256r1_Seed,
			SigerCode:  idex.ECDSA_256r1,
			Only:       false,
			Index:      63,
		},
		{
			SignerCode: mdex.Ed25519_Seed,
			SigerCode:  idex.Ed25519_Big,
			Only:       false,
			Index:      37,
			Ondex:      &_58,
		},
		{
			SignerCode: mdex.ECDSA_256k1_Seed,
			SigerCode:  idex.ECDSA_256k1_Big,
			Only:       false,
			Index:      23,
			Ondex:      &_73,
		},
		{
			SignerCode: mdex.ECDSA_256r1_Seed,
			SigerCode:  idex.ECDSA_256r1_Big,
			Only:       false,
			Index:      64,
			Ondex:      &_81,
		},
		{
			SignerCode: mdex.Ed25519_Seed,
			SigerCode:  idex.Ed25519_Crt,
			Only:       true,
			Index:      37,
		},
		{
			SignerCode: mdex.ECDSA_256k1_Seed,
			SigerCode:  idex.ECDSA_256k1_Crt,
			Only:       true,
			Index:      23,
		},
		{
			SignerCode: mdex.ECDSA_256r1_Seed,
			SigerCode:  idex.ECDSA_256r1_Crt,
			Only:       true,
			Index:      63,
		},
		{
			SignerCode: mdex.Ed25519_Seed,
			SigerCode:  idex.Ed25519_Big_Crt,
			Only:       true,
			Index:      64,
		},
		{
			SignerCode: mdex.ECDSA_256k1_Seed,
			SigerCode:  idex.ECDSA_256k1_Big_Crt,
			Only:       true,
			Index:      65,
		},
		{
			SignerCode: mdex.ECDSA_256r1_Seed,
			SigerCode:  idex.ECDSA_256r1_Big_Crt,
			Only:       true,
			Index:      66,
		},
	}

	for _, testVector := range testCases {
		label := fmt.Sprintf("%s->%s[%t]", testVector.SignerCode, testVector.SigerCode, testVector.Only)
		t.Run(label, func(t *testing.T) {
			signer, err := cesr.NewSigner(true, options.WithCode(testVector.SignerCode))
			if err != nil {
				t.Fatalf("failed to create signer: %v", err)
			}

			siger, err := signer.SignIndexed([]byte{}, testVector.Only, testVector.Index, testVector.Ondex)
			if err != nil {
				t.Fatalf("failed to sign indexed: %v", err)
			}

			if siger.GetCode() != testVector.SigerCode {
				t.Fatalf("siger code mismatch: %s != %s", siger.GetCode(), testVector.SigerCode)
			}

			if siger.GetIndex() != testVector.Index {
				t.Fatalf("siger index mismatch: %d != %d", siger.GetIndex(), testVector.Index)
			}

			ondex := siger.GetOndex()
			if testVector.Only && ondex == nil {
				return
			}

			if testVector.Ondex == nil && *ondex == types.Ondex(testVector.Index) {
				return
			}

			if testVector.Ondex != nil && *ondex == *testVector.Ondex {
				return
			}

			t.Fatalf("siger ondex mismatch: %d", *ondex)
		})
	}
}
