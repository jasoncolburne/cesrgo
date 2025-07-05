package test

import (
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	"github.com/jasoncolburne/cesrgo/core/types"
)

func TestSaiderSaidification(t *testing.T) {
	sad := types.NewMap()
	sad.Set("d", "")

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
