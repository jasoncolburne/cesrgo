package test

import (
	"bytes"
	"crypto/rand"
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
)

func TestPrefixerRoundTrip(t *testing.T) {
	raw := [32]byte{}
	_, err := rand.Read(raw[:])
	if err != nil {
		t.Fatalf("failed to read random bytes: %v", err)
	}

	prefixer, err := cesr.NewPrefixer(options.WithCode(codex.Blake3_256), options.WithRaw(raw[:]))
	if err != nil {
		t.Fatalf("failed to create prefixer: %v", err)
	}

	qb2, err := prefixer.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Prefixer, err := cesr.NewPrefixer(options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create prefixer: %v", err)
	}

	qb64, err := qb2Prefixer.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Prefixer, err := cesr.NewPrefixer(options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create prefixer: %v", err)
	}

	qb64b, err := qb64Prefixer.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bPrefixer, err := cesr.NewPrefixer(options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create prefixer: %v", err)
	}

	qb64bRaw := qb64bPrefixer.GetRaw()
	if !bytes.Equal(qb64bRaw, raw[:]) {
		t.Fatalf("raw mismatch: %v != %v", qb64bRaw, raw)
	}
}
