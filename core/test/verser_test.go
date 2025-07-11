package test

import (
	"testing"

	"github.com/jasoncolburne/cesrgo"
	cesr "github.com/jasoncolburne/cesrgo/core"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

func TestVerserRoundTrip(t *testing.T) {
	proto := cesrgo.Proto_KERI
	pvrsn := cesrgo.VERSION_2_0
	gvrsn := &cesrgo.VERSION_2_0

	atomicVerser, err := cesr.NewVerser(nil, &proto, &pvrsn, gvrsn)
	if err != nil {
		t.Fatalf("failed to create verser: %v", err)
	}

	versage, err := atomicVerser.Versage()
	if err != nil {
		t.Fatalf("failed to get versage: %v", err)
	}

	versageVerser, err := cesr.NewVerser(&versage, nil, nil, nil)
	if err != nil {
		t.Fatalf("failed to create verser: %v", err)
	}

	qb2, err := versageVerser.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Verser, err := cesr.NewVerser(nil, nil, nil, nil, options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create verser: %v", err)
	}

	qb64, err := qb2Verser.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Verser, err := cesr.NewVerser(nil, nil, nil, nil, options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create verser: %v", err)
	}

	qb64b, err := qb64Verser.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bVerser, err := cesr.NewVerser(nil, nil, nil, nil, options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create verser: %v", err)
	}

	expectedVersage := types.Versage{
		Proto: proto,
		Pvrsn: pvrsn,
		Gvrsn: gvrsn,
	}

	derivedVersage, err := qb64bVerser.Versage()
	if err != nil {
		t.Fatalf("failed to get versage: %v", err)
	}

	if derivedVersage.Proto != expectedVersage.Proto || derivedVersage.Pvrsn != expectedVersage.Pvrsn || *derivedVersage.Gvrsn != *expectedVersage.Gvrsn {
		t.Fatalf("versage mismatch: got %v (gv=%v), expected %v (gv=%v)", derivedVersage, *derivedVersage.Gvrsn, expectedVersage, *expectedVersage.Gvrsn)
	}
}
