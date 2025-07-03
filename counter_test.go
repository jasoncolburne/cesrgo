package cesrgo_test

import (
	"testing"

	"github.com/jasoncolburne/cesrgo"
	"github.com/jasoncolburne/cesrgo/counter/options"
	"github.com/jasoncolburne/cesrgo/counter/two"
	"github.com/jasoncolburne/cesrgo/types"
)

func TestNewCounter(t *testing.T) {
	initialCount := types.Count(128)

	counter, err := cesrgo.NewCounter(
		options.WithCode(two.BigAttachmentGroup),
		options.WithCount(initialCount),
	)
	if err != nil {
		t.Fatalf("failed to create counter: %v", err)
	}

	if counter.GetCount() != initialCount {
		t.Fatalf("counter count mismatch: expected %d, got %d", initialCount, counter.GetCount())
	}

	if counter.GetCode() != two.BigAttachmentGroup {
		t.Fatalf("counter code mismatch: expected %s, got %s", two.BigAttachmentGroup, counter.GetCode())
	}

	qb2, err := counter.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	counter2, err := cesrgo.NewCounter(
		options.WithQb2(qb2),
	)
	if err != nil {
		t.Fatalf("failed to create counter: %v", err)
	}

	if counter2.GetCount() != initialCount {
		t.Fatalf("counter count mismatch: expected %d, got %d", initialCount, counter2.GetCount())
	}

	if counter2.GetCode() != two.BigAttachmentGroup {
		t.Fatalf("counter code mismatch: expected %s, got %s", two.BigAttachmentGroup, counter2.GetCode())
	}

	qb64, err := counter2.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	counter3, err := cesrgo.NewCounter(
		options.WithQb64(qb64),
	)
	if err != nil {
		t.Fatalf("failed to create counter: %v", err)
	}

	if counter3.GetCount() != initialCount {
		t.Fatalf("counter count mismatch: expected %d, got %d", initialCount, counter3.GetCount())
	}

	qb64b, err := counter2.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	counter4, err := cesrgo.NewCounter(
		options.WithQb64b(qb64b),
	)
	if err != nil {
		t.Fatalf("failed to create counter: %v", err)
	}

	if counter4.GetCount() != initialCount {
		t.Fatalf("counter count mismatch: expected %d, got %d", initialCount, counter4.GetCount())
	}

	if counter4.GetCode() != two.BigAttachmentGroup {
		t.Fatalf("counter code mismatch: expected %s, got %s", two.BigAttachmentGroup, counter4.GetCode())
	}
}
