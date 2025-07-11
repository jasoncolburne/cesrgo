package test

import (
	"bytes"
	"testing"

	"github.com/jasoncolburne/cesrgo/common"
	cesr "github.com/jasoncolburne/cesrgo/core"
	codex "github.com/jasoncolburne/cesrgo/core/matter"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
	"github.com/jasoncolburne/cesrgo/core/types"
)

func TestSalterCompatibility(t *testing.T) {
	// >>> from keri.core.signing import Salter
	// >>> raw = b"\x19?\xfa\xc7\x8f\x8b\x7f\x8b\xdbS\"$\xd7[\x85\x87"
	// >>> s = Salter(raw=raw)
	// >>> s.stretch()
	// b'\xcd\xe0\xe2\xf4+<\xda\xe3\x9b5+\x1e\\\x87*=\xa03\x81t\x7f\xec\xcd\xca>D\xe1D\xc2\x94\xa1\x82'
	// >>> s.stretch(tier='med')
	// b'g\x82\xb1\x93\xc6\xd7\x1e\xc2=\xe3\xca\xd2\x08\x1b\xd1\xcc\xfb#\x01\x10\x8bZ[\x8d\x13\x83h\x02\xea\xf8\xa8\xbb'
	// >>> s.stretch(tier='high')
	// b"&\xfd\xdd\xa4\xda\xc6\xa4|\x85.9\xe2\x1a\xaf\x1cZ\xeb\xe0\xe2J'E;[\x1a\xcd\xee\xb1\xa9Q\xf3z"

	raw := types.Raw("\x19?\xfa\xc7\x8f\x8b\x7f\x8b\xdbS\"$\xd7[\x85\x87")
	salter, err := cesr.NewSalter(nil, options.WithCode(codex.Salt_128), options.WithRaw(raw))
	if err != nil {
		t.Fatalf("failed to create salter: %v", err)
	}

	seed, err := salter.Stretch(nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("failed to stretch: %v", err)
	}

	if !bytes.Equal(seed, []byte("\xcd\xe0\xe2\xf4+<\xda\xe3\x9b5+\x1e\\\x87*=\xa03\x81t\x7f\xec\xcd\xca>D\xe1D\xc2\x94\xa1\x82")) {
		t.Fatalf("seed mismatch")
	}

	seed, err = salter.Stretch(nil, nil, &common.TIER_MED, nil)
	if err != nil {
		t.Fatalf("failed to stretch: %v", err)
	}

	if !bytes.Equal(seed, []byte("g\x82\xb1\x93\xc6\xd7\x1e\xc2=\xe3\xca\xd2\x08\x1b\xd1\xcc\xfb#\x01\x10\x8bZ[\x8d\x13\x83h\x02\xea\xf8\xa8\xbb")) {
		t.Fatalf("seed mismatch")
	}

	seed, err = salter.Stretch(nil, nil, &common.TIER_HIGH, nil)
	if err != nil {
		t.Fatalf("failed to stretch: %v", err)
	}

	if !bytes.Equal(seed, []byte("&\xfd\xdd\xa4\xda\xc6\xa4|\x85.9\xe2\x1a\xaf\x1cZ\xeb\xe0\xe2J'E;[\x1a\xcd\xee\xb1\xa9Q\xf3z")) {
		t.Fatalf("seed mismatch")
	}
}
