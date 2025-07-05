package test

import (
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	"github.com/jasoncolburne/cesrgo/core/counter/options"
	"github.com/jasoncolburne/cesrgo/core/counter/two"
	"github.com/jasoncolburne/cesrgo/core/types"
)

func TestCounterRoundTrip(t *testing.T) {
	count := types.Count(128)

	codeAndCountCounter, err := cesr.NewCounter(
		options.WithCode(two.BigAttachmentGroup),
		options.WithCount(count),
	)
	if err != nil {
		t.Fatalf("failed to create counter: %v", err)
	}

	qb2, err := codeAndCountCounter.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Counter, err := cesr.NewCounter(
		options.WithQb2(qb2),
	)
	if err != nil {
		t.Fatalf("failed to create counter: %v", err)
	}

	qb64, err := qb2Counter.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Counter, err := cesr.NewCounter(
		options.WithQb64(qb64),
	)
	if err != nil {
		t.Fatalf("failed to create counter: %v", err)
	}

	qb64b, err := qb64Counter.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bCounter, err := cesr.NewCounter(
		options.WithQb64b(qb64b),
	)
	if err != nil {
		t.Fatalf("failed to create counter: %v", err)
	}

	if qb64bCounter.GetCount() != count {
		t.Fatalf("counter count mismatch: expected %d, got %d", count, qb64bCounter.GetCount())
	}

	if qb64bCounter.GetCode() != two.BigAttachmentGroup {
		t.Fatalf("counter code mismatch: expected %s, got %s", two.BigAttachmentGroup, qb64bCounter.GetCode())
	}
}
