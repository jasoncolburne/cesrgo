package cesrgo_test

import (
	"testing"

	"github.com/jasoncolburne/cesrgo"
	iopts "github.com/jasoncolburne/cesrgo/indexer/options"
	mdex "github.com/jasoncolburne/cesrgo/matter"
	mopts "github.com/jasoncolburne/cesrgo/matter/options"
	"github.com/jasoncolburne/cesrgo/types"
)

func TestNewSiger(t *testing.T) {
	signer, err := cesrgo.NewSigner(true, mopts.WithCode(mdex.Ed25519_Seed))
	if err != nil {
		t.Fatalf("NewSigner: %v", err)
	}

	siger, err := signer.SignIndexed([]byte{}, true, types.Index(0), nil)
	if err != nil {
		t.Fatalf("SignIndexed: %v", err)
	}

	qb64, err := siger.Qb64()
	if err != nil {
		t.Fatalf("Qb64: %v", err)
	}

	nVerfer := signer.GetVerfer()
	siger2, err := cesrgo.NewSiger(nVerfer, iopts.WithQb64(qb64))
	if err != nil {
		t.Fatalf("NewSiger: %v", err)
	}

	gVerfer := siger.GetVerfer()

	err = nVerfer.Verify(siger.GetRaw(), []byte{})
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}

	err = gVerfer.Verify(siger2.GetRaw(), []byte{})
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}

	err = nVerfer.Verify(siger.GetRaw(), []byte{1})
	if err == nil {
		t.Fatalf("Verify: %v", err)
	}

	err = gVerfer.Verify(siger2.GetRaw(), []byte{1})
	if err == nil {
		t.Fatalf("Verify: %v", err)
	}

	qb2, err := siger.Qb2()
	if err != nil {
		t.Fatalf("Qb2: %v", err)
	}

	siger3, err := cesrgo.NewSiger(nVerfer, iopts.WithQb2(qb2))
	if err != nil {
		t.Fatalf("NewSiger: %v", err)
	}

	err = nVerfer.Verify(siger3.GetRaw(), []byte{})
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}
}
