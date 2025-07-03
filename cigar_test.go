package cesrgo_test

import (
	"testing"

	"github.com/jasoncolburne/cesrgo"
	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
)

func Test256r1Cigar(t *testing.T) {
	signer, err := cesrgo.NewSigner(true, options.WithCode(codex.ECDSA_256r1_Seed))
	if err != nil {
		t.Fatalf("failed to create signer: %v", err)
	}

	cigar, err := signer.SignUnindexed([]byte{})
	if err != nil {
		t.Fatalf("failed to sign unindexed: %v", err)
	}

	sVerfer := signer.GetVerfer()
	cVerfer := cigar.GetVerfer()

	svQb64, err := sVerfer.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	cvQb64, err := cVerfer.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	if svQb64 != cvQb64 {
		t.Fatalf("signer verfer qb64 mismatch: %s != %s", svQb64, cvQb64)
	}

	err = sVerfer.Verify(cigar.GetRaw(), []byte{})
	if err != nil {
		t.Fatalf("failed to verify: %v", err)
	}

	badRaw := cigar.GetRaw()
	badRaw[len(badRaw)-1] ^= 0x01
	err = sVerfer.Verify(badRaw, []byte{})
	if err == nil {
		t.Fatalf("should have failed to verify")
	}
}

func Test256k1Cigar(t *testing.T) {
	signer, err := cesrgo.NewSigner(true, options.WithCode(codex.ECDSA_256k1_Seed))
	if err != nil {
		t.Fatalf("failed to create signer: %v", err)
	}

	cigar, err := signer.SignUnindexed([]byte{})
	if err != nil {
		t.Fatalf("failed to sign unindexed: %v", err)
	}

	sVerfer := signer.GetVerfer()
	cVerfer := cigar.GetVerfer()

	svQb64, err := sVerfer.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	cvQb64, err := cVerfer.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	if svQb64 != cvQb64 {
		t.Fatalf("signer verfer qb64 mismatch: %s != %s", svQb64, cvQb64)
	}

	err = sVerfer.Verify(cigar.GetRaw(), []byte{})
	if err != nil {
		t.Fatalf("failed to verify: %v", err)
	}

	badRaw := cigar.GetRaw()
	badRaw[len(badRaw)-1] ^= 0x01
	err = sVerfer.Verify(badRaw, []byte{})
	if err == nil {
		t.Fatalf("should have failed to verify")
	}
}
