package test

import (
	"math/big"
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
)

func TestSeqnerDefaults(t *testing.T) {
	seqner, err := cesr.NewSeqner(nil, nil)
	if err != nil {
		t.Fatalf("failed to create seqner: %v", err)
	}

	sn := seqner.Sn()
	if sn.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("sn is not 0: %X", sn.Bytes())
	}

	if seqner.GetCode() != codex.Huge {
		t.Fatalf("seqner is not coded as huge")
	}
}

func TestSeqnerRoundTrip(t *testing.T) {
	sn := big.NewInt(57284942893)
	snSeqner, err := cesr.NewSeqner(sn, nil)
	if err != nil {
		t.Fatalf("failed to create seqner: %v", err)
	}

	qb2, err := snSeqner.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Seqner, err := cesr.NewSeqner(nil, nil, options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create seqner from qb2: %v", err)
	}

	qb64, err := qb2Seqner.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Seqner, err := cesr.NewSeqner(nil, nil, options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create seqner from qb64: %v", err)
	}

	qb64b, err := qb64Seqner.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bSeqner, err := cesr.NewSeqner(nil, nil, options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create seqner from qb64b: %v", err)
	}

	qb64bSn := qb64bSeqner.Sn()

	if qb64bSn.Cmp(sn) != 0 {
		t.Fatalf("qb64b sn is not the same as sn: %X != %X", qb64bSn.Bytes(), sn.Bytes())
	}
}
