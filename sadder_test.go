package cesrgo_test

import (
	"testing"

	"github.com/jasoncolburne/cesrgo"
	"github.com/jasoncolburne/cesrgo/types"
)

func TestNewSadder(t *testing.T) {
	ked := types.NewMap()

	ked.Set("vs", "KERICAACAAJSONAAAA.")
	ked.Set("d", "")

	sadder, err := cesrgo.NewSadder(&ked, nil, true)
	if err != nil {
		t.Fatalf("failed to create sadder: %v", err)
	}

	json, err := sadder.GetKed().MarshalJSON()
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	if string(json) != "{\"vs\":\"KERICAACAAJSONAABP.\",\"d\":\"EPyO5MrDqLBmvskDfXBnwWHBv27pAPcypkDfvn9468Tr\"}" {
		t.Fatalf("json mismatch: %s", string(json))
	}
}
