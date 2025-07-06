package test

import (
	"testing"

	"github.com/jasoncolburne/cesrgo"
	cesr "github.com/jasoncolburne/cesrgo/core"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
)

func TestTraitorTraits(t *testing.T) {
	for _, trait := range cesrgo.TRAITS {
		label := string(trait)
		t.Run(label, func(t *testing.T) {
			traitor, err := cesr.NewTraitor(&trait)
			if err != nil {
				t.Fatalf("failed to create traitor for %s: %v", trait, err)
			}

			derivedTrait, err := traitor.Trait()
			if err != nil {
				t.Fatalf("failed to get trait: %v", err)
			}

			if derivedTrait != trait {
				t.Fatalf("derivedTrait != %s", trait)
			}
		})
	}
}

func TestTraitorRoundTrip(t *testing.T) {
	trait := cesrgo.Trait_EstOnly
	traitTraitor, err := cesr.NewTraitor(&trait)
	if err != nil {
		t.Fatalf("failed to create trait traitor: %v", err)
	}

	qb2, err := traitTraitor.Qb2()
	if err != nil {
		t.Fatalf("failed to get qb2: %v", err)
	}

	qb2Traitor, err := cesr.NewTraitor(nil, options.WithQb2(qb2))
	if err != nil {
		t.Fatalf("failed to create trait traitor from qb2: %v", err)
	}

	qb64, err := qb2Traitor.Qb64()
	if err != nil {
		t.Fatalf("failed to get qb64: %v", err)
	}

	qb64Traitor, err := cesr.NewTraitor(nil, options.WithQb64(qb64))
	if err != nil {
		t.Fatalf("failed to create trait traitor from qb64: %v", err)
	}

	qb64b, err := qb64Traitor.Qb64b()
	if err != nil {
		t.Fatalf("failed to get qb64b: %v", err)
	}

	qb64bTraitor, err := cesr.NewTraitor(nil, options.WithQb64b(qb64b))
	if err != nil {
		t.Fatalf("failed to create trait traitor from qb64b: %v", err)
	}

	derivedTrait, err := qb64bTraitor.Trait()
	if err != nil {
		t.Fatalf("failed to get trait: %v", err)
	}

	if derivedTrait != trait {
		t.Fatalf("derivedTrait != %s", trait)
	}
}
