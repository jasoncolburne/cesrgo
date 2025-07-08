package test

import (
	"bytes"
	"fmt"
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
)

func TestNoncerRoundTrip(t *testing.T) {
	noncer, err := cesr.NewNoncer(nil, options.WithCode(codex.Salt_128))
	if err != nil {
		t.Fatalf("failed to create noncer: %v", err)
	}

	qb2, err := noncer.Qb2()
	if err != nil {
		t.Fatalf("failed to get nonce: %v", err)
	}

	qb2Noncer, err := cesr.NewNoncer(nil, options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create noncer: %v", err)
	}

	qb64, err := qb2Noncer.Qb64()
	if err != nil {
		t.Fatalf("failed to get nonce: %v", err)
	}

	qb64Noncer, err := cesr.NewNoncer(nil, options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create noncer: %v", err)
	}

	qb64b, err := qb64Noncer.Qb64b()
	if err != nil {
		t.Fatalf("failed to get nonce: %v", err)
	}

	qb64bNoncer, err := cesr.NewNoncer(nil, options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create noncer: %v", err)
	}

	noncerNonce, err := noncer.Nonceb()
	if err != nil {
		t.Fatalf("failed to get nonce: %v", err)
	}

	qb64bNonce, err := qb64bNoncer.Nonceb()
	if err != nil {
		t.Fatalf("failed to get nonce: %v", err)
	}

	fmt.Printf("nonce: %q", noncerNonce)
	if !bytes.Equal(noncerNonce, qb64bNonce) {
		t.Fatalf("nonce mismatch: %v != %v", noncerNonce, qb64bNonce)
	}
}
