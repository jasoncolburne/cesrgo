package test

import (
	"testing"

	"github.com/jasoncolburne/cesrgo"
	"github.com/jasoncolburne/cesrgo/types"
)

func TestSadderMarshalling(t *testing.T) {
	ked := types.NewMap()

	ked.Set("v", "KERICAACAAJSONAAAA.")
	ked.Set("d", "")

	sadder, err := cesrgo.NewSadder(nil, nil, &ked, nil, true)
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
