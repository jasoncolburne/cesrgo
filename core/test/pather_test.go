package test

import (
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
)

func TestPatherRoundTrip(t *testing.T) {
	path := "AB/CD/EF/GH/IJ/KL/MN/OP/QR/ST/UV/WX/YZ/01/23/45/67/89/_"

	pather, err := cesr.NewPather(&path, nil, true, true)
	if err != nil {
		t.Fatalf("failed to create pather: %v", err)
	}

	qb2, err := pather.Qb2()
	if err != nil {
		t.Fatalf("failed to create qb2: %v", err)
	}

	qb2Pather, err := cesr.NewPather(nil, nil, true, true, options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to get qb2 path: %v", err)
	}

	qb64, err := qb2Pather.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Pather, err := cesr.NewPather(nil, nil, true, true, options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create qb64 pather: %v", err)
	}

	qb64b, err := qb64Pather.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bPather, err := cesr.NewPather(nil, nil, true, true, options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create qb64b pather: %v", err)
	}

	qb64bPath, err := qb64bPather.Path()
	if err != nil {
		t.Fatalf("failed to get qb64b path: %v", err)
	}

	if qb64bPath != path {
		t.Fatalf("qb64b path mismatch: %s != %s", qb64bPath, path)
	}
}
