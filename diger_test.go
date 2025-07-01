package cesrgo_test

import (
	"testing"

	"github.com/jasoncolburne/cesrgo"
	codex "github.com/jasoncolburne/cesrgo/matter"
	"github.com/jasoncolburne/cesrgo/matter/options"
)

func TestNewDiger(t *testing.T) {
	dig, err := cesrgo.NewDiger([]byte{}, options.WithCode(codex.Blake3_256))
	if err != nil {
		t.Fatalf("ser + code should not fail: %v", err)
	}

	dig2, err := cesrgo.NewDiger(nil, options.WithCode(codex.Blake3_256), options.WithRaw(dig.GetRaw()))
	if err != nil {
		t.Fatalf("code + raw should not fail: %v", err)
	}

	verified, err := dig2.Verify([]byte{})
	if err != nil {
		t.Fatalf("verify should not fail: %v", err)
	}

	if !verified {
		t.Fatalf("value should verify")
	}

	verified, err = dig2.Verify([]byte{1})
	if err != nil {
		t.Fatalf("verify should not fail: %v", err)
	}

	if verified {
		t.Fatalf("value should not verify")
	}

	qb64, err := dig2.Qb64()
	if err != nil {
		t.Fatalf("qb64 should not fail: %v", err)
	}

	if qb64 == "" {
		t.Fatalf("qb64 should not be empty")
	}

	dig3, err := cesrgo.NewDiger(nil, options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("qb64 should not fail: %v", err)
	}

	verified, err = dig3.Verify([]byte{})
	if err != nil {
		t.Fatalf("verify should not fail: %v", err)
	}
}
