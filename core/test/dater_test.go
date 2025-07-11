package test

import (
	"testing"
	"time"

	cesr "github.com/jasoncolburne/cesrgo/core"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

func TestDaterRoundTrip(t *testing.T) {
	dts := types.DateTime("2006-01-02T15:04:05.000000+07:00")
	dtsDater, err := cesr.NewDater(&dts)
	if err != nil {
		t.Fatalf("failed to create dater: %v", err)
	}

	qb2, err := dtsDater.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Dater, err := cesr.NewDater(nil, options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create dater: %v", err)
	}

	qb64, err := qb2Dater.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Dater, err := cesr.NewDater(nil, options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create dater: %v", err)
	}

	qb64b, err := qb64Dater.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bDater, err := cesr.NewDater(nil, options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create dater: %v", err)
	}

	qb64bDts, err := qb64bDater.DTS()
	if err != nil {
		t.Fatalf("failed to get dts: %v", err)
	}

	if qb64bDts != types.DateTime(dts) {
		t.Fatalf("qb64b dts %s does not match dts %s", qb64bDts, dts)
	}
}

func TestDaterQb64Represenatation(t *testing.T) {
	dts := types.DateTime("2006-01-02T15:04:05.000000+07:00")
	dtsDater, err := cesr.NewDater(&dts)
	if err != nil {
		t.Fatalf("failed to create dater: %v", err)
	}

	qb64, err := dtsDater.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	if qb64 != "1AAG2006-01-02T15c04c05d000000p07c00" {
		t.Fatalf("qb64 %s does not match expected %s", qb64, "1AAG2006-01-02T15c04c05d000000p07c00")
	}
}

func TestDaterDefaultTemporality(t *testing.T) {
	then := time.Now().UTC()
	time.Sleep(1 * time.Microsecond)

	dtsDater, err := cesr.NewDater(nil)
	if err != nil {
		t.Fatalf("failed to create dater: %v", err)
	}

	time.Sleep(1 * time.Microsecond)
	now := time.Now().UTC()

	dts, err := dtsDater.DTS()
	if err != nil {
		t.Fatalf("failed to get dts: %v", err)
	}

	dtsTime, err := time.Parse(time.RFC3339, string(dts))
	if err != nil {
		t.Fatalf("failed to parse dts: %v", err)
	}

	if dtsTime.Before(then) || dtsTime.After(now) {
		t.Fatalf("dts time %s is not within the range of then %s and now %s", dtsTime, then, now)
	}
}
