package test

import (
	"testing"

	"github.com/jasoncolburne/cesrgo"
	cesr "github.com/jasoncolburne/cesrgo/core"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
)

func TestIlks(t *testing.T) {
	for _, ilk := range cesrgo.ILKS {
		label := string(ilk)
		t.Run(label, func(t *testing.T) {
			ilker, err := cesr.NewIlker(&ilk)
			if err != nil {
				t.Fatalf("failed to create ilker for %s: %v", ilk, err)
			}

			if ilker.GetSoft() != string(ilk) {
				t.Fatalf("ilker.GetSoft() != %s", ilk)
			}
		})
	}
}

func TestRoundTrip(t *testing.T) {
	ilk := cesrgo.Ilk_ICP
	ilkIlker, err := cesr.NewIlker(&ilk)
	if err != nil {
		t.Fatalf("failed to create ilker: %v", err)
	}

	qb2, err := ilkIlker.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Ilker, err := cesr.NewIlker(nil, options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create ilker from qb2: %v", err)
	}

	qb64, err := qb2Ilker.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Ilker, err := cesr.NewIlker(nil, options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create ilker from qb64: %v", err)
	}

	qb64b, err := qb64Ilker.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bIlker, err := cesr.NewIlker(nil, options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create ilker from qb64b: %v", err)
	}

	if qb64bIlker.GetSoft() != string(ilk) {
		t.Fatalf("qb64bIlker.GetSoft() != %s", ilk)
	}
}
