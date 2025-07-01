package cesrgo_test

import (
	"testing"

	"github.com/jasoncolburne/cesrgo"
)

func TestNewSigner(t *testing.T) {
	signer, err := cesrgo.NewSigner(true)
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
}
