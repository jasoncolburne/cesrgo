package test

import (
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
)

func TestLabelerRoundTrip(t *testing.T) {
	testCases := []string{
		"testlabel",
		"testlabel1",
		"testlabel12",
		"testlabel123",
		"testlabel1234",
		"testlabel12345",
	}

	for _, testCase := range testCases {
		label := testCase
		t.Run(label, func(t *testing.T) {
			labeler, err := cesr.NewLabeler(&label)
			if err != nil {
				t.Fatalf("failed to create labeler: %v", err)
			}

			qb2, err := labeler.Qb2()
			if err != nil {
				t.Fatalf("failed to get label: %v", err)
			}

			qb2Labeler, err := cesr.NewLabeler(nil, options.WithQb2(qb2))
			if err != nil {
				t.Fatalf("failed to create labeler: %v", err)
			}

			qb64, err := qb2Labeler.Qb64()
			if err != nil {
				t.Fatalf("failed to get qb64: %v", err)
			}

			qb64Labeler, err := cesr.NewLabeler(nil, options.WithQb64(qb64))
			if err != nil {
				t.Fatalf("failed to create labeler: %v", err)
			}

			qb64b, err := qb64Labeler.Qb64b()
			if err != nil {
				t.Fatalf("failed to get qb64b: %v", err)
			}

			qb64bLabeler, err := cesr.NewLabeler(nil, options.WithQb64b(qb64b))
			if err != nil {
				t.Fatalf("failed to create labeler: %v", err)
			}

			qb64bLabel, err := qb64bLabeler.Label()
			if err != nil {
				t.Fatalf("failed to get label: %v", err)
			}

			if qb64bLabel != label {
				t.Fatalf("label mismatch: %s != %s", qb64bLabel, label)
			}
		})
	}
}
