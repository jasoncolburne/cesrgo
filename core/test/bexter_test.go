package test

import (
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	"github.com/jasoncolburne/cesrgo/core/matter/options"
)

func TestBexterRoundTrip(t *testing.T) {
	testCases := []string{
		"1234ABCD",
		"1234ABCDE",
		"1234ABCDEF",
		"1234ABCDEFG",
		"1234ABCDEFGH",
		"1234ABCDEFGHI",
		"1234ABCDEFGHIJ",
	}

	for _, testCase := range testCases {
		bext := testCase
		t.Run(bext, func(t *testing.T) {
			bexter, err := cesr.NewBexter(&bext)
			if err != nil {
				t.Fatalf("failed to create bexter: %v", err)
			}

			qb2, err := bexter.Qb2()
			if err != nil {
				t.Fatalf("failed to get qb2: %v", err)
			}

			qb2Bexter, err := cesr.NewBexter(nil, options.WithQb2(qb2))
			if err != nil {
				t.Fatalf("failed to create qb2 bexter: %v", err)
			}

			qb64, err := qb2Bexter.Qb64()
			if err != nil {
				t.Fatalf("failed to get qb64: %v", err)
			}

			qb64Bexter, err := cesr.NewBexter(nil, options.WithQb64(qb64))
			if err != nil {
				t.Fatalf("failed to create qb64 bexter: %v", err)
			}

			qb64b, err := qb64Bexter.Qb64b()
			if err != nil {
				t.Fatalf("failed to get qb64b: %v", err)
			}

			qb64bBexter, err := cesr.NewBexter(nil, options.WithQb64b(qb64b))
			if err != nil {
				t.Fatalf("failed to create qb64b bexter: %v", err)
			}

			qb64bBext, err := qb64bBexter.Bext()
			if err != nil {
				t.Fatalf("failed to get qb64b bext: %v", err)
			}

			if bext != qb64bBext {
				t.Fatalf("qb64b bext does not match original bext: %s != %s", bext, qb64bBext)
			}
		})
	}
}
