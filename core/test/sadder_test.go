package test

import (
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	"github.com/jasoncolburne/cesrgo/core/types"
)

func TestSadderMarshalling(t *testing.T) {
	ked := types.NewMap(
		[]string{"v", "d"},
		[]any{"KERICAACAAJSONAAAA.", ""},
	)

	sadder, err := cesr.NewSadder(nil, nil, &ked, nil, true)
	if err != nil {
		t.Fatalf("failed to create sadder: %v", err)
	}

	json, err := sadder.GetKed().MarshalJSON()
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	if string(json) != "{\"v\":\"KERICAACAAJSONAABO.\",\"d\":\"EHJq2PWESIo1D4z3ca3ve7UKpwZ4uzmp-LCV5VO9v7OU\"}" {
		t.Fatalf("json mismatch: %s", string(json))
	}
}
