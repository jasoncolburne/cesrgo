package test

import (
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
)

func TestTexterRoundTrip(t *testing.T) {
	text := "KERI/ACDC"
	texter, err := cesr.NewTexter(&text, options.WithCode(codex.Bytes_L0))
	if err != nil {
		t.Fatalf("failed to create texter: %v", err)
	}

	qb2, err := texter.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Texter, err := cesr.NewTexter(nil, options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create texter from qb2: %v", err)
	}

	qb64, err := qb2Texter.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Texter, err := cesr.NewTexter(nil, options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create texter from qb64: %v", err)
	}

	qb64b, err := qb64Texter.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bTexter, err := cesr.NewTexter(nil, options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create texter from qb64b: %v", err)
	}

	qb64bText := qb64bTexter.Text()
	if qb64bText != text {
		t.Fatalf("qb64b text mismatch: %s != %s", qb64bText, text)
	}
}
