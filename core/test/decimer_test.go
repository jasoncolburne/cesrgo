package test

import (
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
)

func TestDecimerRoundTrip(t *testing.T) {
	dns := "1234567812.123"
	decimer, err := cesr.NewDecimer(&dns, nil, options.WithCode(codex.Decimal_L2))
	if err != nil {
		t.Fatalf("failed to create decimer: %v", err)
	}

	qb2, err := decimer.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Decimer, err := cesr.NewDecimer(nil, nil, options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create decimer from qb2: %v", err)
	}

	qb64, err := qb2Decimer.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64 from qb2 decimer: %v", err)
	}

	qb64Decimer, err := cesr.NewDecimer(nil, nil, options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create decimer from qb64: %v", err)
	}

	qb64b, err := qb64Decimer.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b from qb64 decimer: %v", err)
	}

	qb64bDecimer, err := cesr.NewDecimer(nil, nil, options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create decimer from qb64b: %v", err)
	}

	qb64bDns, err := qb64bDecimer.Dns()
	if err != nil {
		t.Fatalf("failed to get decimal from qb64b decimer: %v", err)
	}

	if qb64bDns != dns {
		t.Fatalf("dns mismatch: %v != %v", qb64bDns, dns)
	}
}
