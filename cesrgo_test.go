package cesrgo_test

import (
	"testing"

	"github.com/jasoncolburne/cesrgo"
)

func TestHello(t *testing.T) {
	version := cesrgo.Version
	expected := "0.1.0"

	if version != expected {
		t.Errorf("Version = %q, want %q", version, expected)
	}
}
