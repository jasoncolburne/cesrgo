package test

import (
	"bytes"
	"crypto/rand"
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

func TestSaiderSaidification(t *testing.T) {
	sad := types.NewMap(
		[]string{"d"},
		[]any{""},
	)

	saider, err := cesr.NewSaider(&sad, nil, nil)
	if err != nil {
		t.Fatalf("failed to create saider: %v", err)
	}

	qb64, err := saider.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	// echo -n '{"d":""}' > file.json && kli saidify --file=file.json && cat file.json | jq -r .d && rm file.json
	if string(qb64) != "EIeKlm9B5ul5vsHu_-OpjNmSf1kn1iMsyTb7rpuE4Ylc" {
		t.Fatalf("qb64 mismatch: %s != %s", string(qb64), "EIeKlm9B5ul5vsHu_-OpjNmSf1kn1iMsyTb7rpuE4Ylc")
	}
}

func TestSaiderRoundTrip(t *testing.T) {
	raw := [32]byte{}
	_, err := rand.Read(raw[:])
	if err != nil {
		t.Fatalf("failed to read random bytes: %v", err)
	}

	saider, err := cesr.NewSaider(nil, nil, nil, options.WithRaw(raw[:]), options.WithCode(codex.Blake3_256))
	if err != nil {
		t.Fatalf("failed to create saider: %v", err)
	}

	qb2, err := saider.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Saider, err := cesr.NewSaider(nil, nil, nil, options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create saider from qb2: %v", err)
	}

	qb64, err := qb2Saider.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Saider, err := cesr.NewSaider(nil, nil, nil, options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create saider from qb64: %v", err)
	}

	qb64b, err := qb64Saider.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bSaider, err := cesr.NewSaider(nil, nil, nil, options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create saider from qb64b: %v", err)
	}

	qb64bRaw := qb64bSaider.GetRaw()
	if !bytes.Equal(qb64bRaw, raw[:]) {
		t.Fatalf("qb64b raw mismatch: %x != %x", qb64bRaw, raw[:])
	}
}
