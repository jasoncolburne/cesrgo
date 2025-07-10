package test

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	idex "github.com/jasoncolburne/cesrgo/core/indexer"
	iopts "github.com/jasoncolburne/cesrgo/core/indexer/options"
	mdex "github.com/jasoncolburne/cesrgo/core/matter"
	mopts "github.com/jasoncolburne/cesrgo/core/matter/options"
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

	for _, testCase := range testCases {
		label := fmt.Sprintf("%s->%s[%t]", testCase.SignerCode, testCase.SigerCode, testCase.Only)
		t.Run(label, func(t *testing.T) {
			signer, err := cesr.NewSigner(true, mopts.WithCode(testCase.SignerCode))
			if err != nil {
				t.Fatalf("failed to create signer: %v", err)
			}

			siger, err := signer.SignIndexed([]byte{}, testCase.Only, testCase.Index, testCase.Ondex)
			if err != nil {
				t.Fatalf("failed to sign indexed: %v", err)
			}

			if siger.GetCode() != testCase.SigerCode {
				t.Fatalf("siger code mismatch: %s != %s", siger.GetCode(), testCase.SigerCode)
			}

			if siger.GetIndex() != testCase.Index {
				t.Fatalf("siger index mismatch: %d != %d", siger.GetIndex(), testCase.Index)
			}

			ondex := siger.GetOndex()
			if testCase.Only && ondex == nil {
				return
			}

			if testCase.Ondex == nil && *ondex == types.Ondex(testCase.Index) {
				return
			}

			if testCase.Ondex != nil && *ondex == *testCase.Ondex {
				return
			}

			t.Fatalf("siger ondex mismatch: %d", *ondex)
		})
	}
}

func TestSigerRoundTrip(t *testing.T) {
	raw := [64]byte{}
	_, err := rand.Read(raw[:])
	if err != nil {
		t.Fatalf("failed to read random bytes: %v", err)
	}

	siger, err := cesr.NewSiger(
		nil,
		iopts.WithCode(idex.Ed25519_Big),
		iopts.WithRaw(raw[:]),
		iopts.WithIndex(37),
		iopts.WithOndex(581),
	)
	if err != nil {
		t.Fatalf("failed to create signer: %v", err)
	}

	qb2, err := siger.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Siger, err := cesr.NewSiger(nil, iopts.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create siger from qb2: %v", err)
	}

	qb64, err := qb2Siger.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Siger, err := cesr.NewSiger(nil, iopts.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create siger from qb64: %v", err)
	}

	qb64b, err := qb64Siger.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bSiger, err := cesr.NewSiger(nil, iopts.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create siger from qb64b: %v", err)
	}

	qb64bRaw := qb64bSiger.GetRaw()
	if !bytes.Equal(qb64bRaw, raw[:]) {
		t.Fatalf("qb64b raw mismatch: %x != %x", qb64bRaw, raw[:])
	}
}
