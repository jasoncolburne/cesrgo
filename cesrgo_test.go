package cesrgo_test

import (
	"testing"

	"github.com/jasoncolburne/cesrgo"
)

func TestVersion(t *testing.T) {
	version := cesrgo.Version
	expected := "2.0.0"

	if version != expected {
		t.Errorf("Version = %q, want %q", version, expected)
	}
}
