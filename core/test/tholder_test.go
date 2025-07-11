package test

import (
	"fmt"
	"math/big"
	"testing"

	cesr "github.com/jasoncolburne/cesrgo/core"
	"github.com/jasoncolburne/cesrgo/core/types"
)

// ported from KERIpy
var testCases = []struct {
	sith           any
	expectedSith   any
	weighted       bool
	thold          any
	limen          types.Qb64
	size           types.Size
	satisfactory   [][]int
	unsatisfactory [][]int
}{
	{
		sith:         "b",
		expectedSith: "b",
		weighted:     false,
		thold:        11,
		limen:        types.Qb64("MAAL"),
		size:         11,
		satisfactory: [][]int{
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		unsatisfactory: [][]int{
			{0, 1, 2},
		},
	},
	{
		sith:         11,
		expectedSith: "b",
		weighted:     false,
		thold:        11,
		limen:        types.Qb64("MAAL"),
		size:         11,
		satisfactory: [][]int{
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		unsatisfactory: [][]int{
			{0, 1, 2},
		},
	},
	{
		sith:         fmt.Sprintf("%x", 15),
		expectedSith: "f",
		weighted:     false,
		thold:        15,
		limen:        types.Qb64("MAAP"),
		size:         15,
		satisfactory: [][]int{
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
		},
		unsatisfactory: [][]int{
			{0, 1, 2},
		},
	},
	{
		sith:         2,
		expectedSith: "2",
		weighted:     false,
		thold:        2,
		limen:        types.Qb64("MAAC"),
		size:         2,
		satisfactory: [][]int{
			{0, 1},
			{0, 1, 2},
		},
		unsatisfactory: [][]int{
			{0},
		},
	},
	{
		sith:         1,
		expectedSith: "1",
		weighted:     false,
		thold:        1,
		limen:        types.Qb64("MAAB"),
		size:         1,
		satisfactory: [][]int{
			{0},
		},
		unsatisfactory: [][]int{
			{},
		},
	},
	{
		sith:         []any{"1/2", "1/2", "1/4", "1/4", "1/4"},
		expectedSith: []any{"1/2", "1/2", "1/4", "1/4", "1/4"},
		weighted:     true,
		thold:        []any{[]any{big.NewRat(1, 2), big.NewRat(1, 2), big.NewRat(1, 4), big.NewRat(1, 4), big.NewRat(1, 4)}},
		limen:        types.Qb64("4AAFA1s2c1s2c1s4c1s4c1s4"),
		size:         5,
		satisfactory: [][]int{
			{0, 1},
			{0, 2, 3},
			{0, 3, 4},
			{0, 2, 4},
			{1, 2, 3},
			{1, 3, 4},
			{1, 2, 4},
		},
		unsatisfactory: [][]int{
			{0, 2},
			{1, 3},
			{2, 3, 4},
		},
	},
	{
		sith:         []any{"1/2", "1/2", "1/4", "1/4", "1/4", "0"},
		expectedSith: []any{"1/2", "1/2", "1/4", "1/4", "1/4", "0"},
		weighted:     true,
		thold:        []any{[]any{big.NewRat(1, 2), big.NewRat(1, 2), big.NewRat(1, 4), big.NewRat(1, 4), big.NewRat(1, 4), big.NewRat(0, 1)}},
		limen:        types.Qb64("6AAGAAA1s2c1s2c1s4c1s4c1s4c0"),
		size:         6,
		satisfactory: [][]int{
			{0, 1},
			{0, 2, 3},
			{0, 3, 4},
			{0, 2, 4},
			{1, 2, 3},
			{1, 3, 4},
			{1, 2, 4},
			{0, 1, 2, 3, 4},
			{3, 2, 0},
			{0, 0, 1, 2, 1},
		},
		unsatisfactory: [][]int{
			{0, 2, 5},
			{1, 3, 5},
			{2, 3, 4, 5},
		},
	},
	{
		sith:         []any{[]any{"1/2", "1/2", "1/4", "1/4", "1/4"}},
		expectedSith: []any{"1/2", "1/2", "1/4", "1/4", "1/4"},
		weighted:     true,
		thold:        []any{[]any{big.NewRat(1, 2), big.NewRat(1, 2), big.NewRat(1, 4), big.NewRat(1, 4), big.NewRat(1, 4)}},
		limen:        types.Qb64("4AAFA1s2c1s2c1s4c1s4c1s4"),
		size:         5,
		satisfactory: [][]int{
			{1, 2, 3},
			{0, 1, 2},
			{1, 3, 4},
			{0, 1, 2, 3, 4},
			{3, 2, 0},
			{0, 0, 1, 2, 1, 4, 4},
		},
		unsatisfactory: [][]int{
			{0, 2},
			{2, 3, 4},
		},
	},
	{
		sith:         []any{[]any{"1/2", "1/2", "1/4", "1/4", "1/4"}, []any{"1/1", "1"}},
		expectedSith: []any{[]any{"1/2", "1/2", "1/4", "1/4", "1/4"}, []any{"1", "1"}},
		weighted:     true,
		thold:        []any{[]any{big.NewRat(1, 2), big.NewRat(1, 2), big.NewRat(1, 4), big.NewRat(1, 4), big.NewRat(1, 4)}, []any{big.NewRat(1, 1), big.NewRat(1, 1)}},
		limen:        types.Qb64("4AAGA1s2c1s2c1s4c1s4c1s4a1c1"),
		size:         7,
		satisfactory: [][]int{
			{1, 2, 3, 5},
			{0, 1, 6},
		},
		unsatisfactory: [][]int{
			{0, 1},
			{5, 6},
			{2, 3, 4},
			{},
		},
	},
	{
		sith:         "[[\"1/2\", \"1/2\", \"1/4\", \"1/4\", \"1/4\"], [\"1/1\", \"1\"]]",
		expectedSith: []any{[]any{"1/2", "1/2", "1/4", "1/4", "1/4"}, []any{"1", "1"}},
		weighted:     true,
		thold:        []any{[]any{big.NewRat(1, 2), big.NewRat(1, 2), big.NewRat(1, 4), big.NewRat(1, 4), big.NewRat(1, 4)}, []any{big.NewRat(1, 1), big.NewRat(1, 1)}},
		limen:        types.Qb64("4AAGA1s2c1s2c1s4c1s4c1s4a1c1"),
		size:         7,
		satisfactory: [][]int{
			{1, 2, 3, 5},
			{0, 1, 6},
		},
		unsatisfactory: [][]int{
			{0, 1},
			{5, 6},
			{2, 3, 4},
			{},
		},
	},
	{
		sith:         "[[\"1/2\", \"1/2\", \"1/4\", \"1/4\", \"1/4\"]]",
		expectedSith: []any{"1/2", "1/2", "1/4", "1/4", "1/4"},
		weighted:     true,
		thold:        []any{[]any{big.NewRat(1, 2), big.NewRat(1, 2), big.NewRat(1, 4), big.NewRat(1, 4), big.NewRat(1, 4)}},
		limen:        types.Qb64("4AAFA1s2c1s2c1s4c1s4c1s4"),
		size:         5,
		satisfactory: [][]int{
			{1, 2, 3},
			{0, 1, 2},
			{1, 3, 4},
			{0, 1, 2, 3, 4},
			{3, 2, 0},
			{0, 0, 1, 2, 1, 4, 4},
		},
		unsatisfactory: [][]int{
			{0, 2},
			{2, 3, 4},
		},
	},
	{
		sith:         "[[\"1/2\", \"1/2\", \"1/4\", \"1/4\", \"1/4\", \"0\"]]",
		expectedSith: []any{"1/2", "1/2", "1/4", "1/4", "1/4", "0"},
		weighted:     true,
		thold:        []any{[]any{big.NewRat(1, 2), big.NewRat(1, 2), big.NewRat(1, 4), big.NewRat(1, 4), big.NewRat(1, 4), big.NewRat(0, 1)}},
		limen:        types.Qb64("6AAGAAA1s2c1s2c1s4c1s4c1s4c0"),
		size:         6,
		satisfactory: [][]int{
			{0, 2, 4},
			{0, 1},
			{1, 3, 4},
			{0, 1, 2, 3, 4},
			{3, 2, 0},
			{0, 0, 1, 2, 1},
		},
		unsatisfactory: [][]int{
			{0, 2, 5},
			{2, 3, 4, 5},
		},
	},
	{
		sith:         "[{\"1/3\":[\"1/2\", \"1/2\", \"1/2\"]}, \"1/3\", \"1/2\", {\"1/2\": [\"1\", \"1\"]}]",
		expectedSith: []any{map[string]any{"1/3": []any{"1/2", "1/2", "1/2"}}, "1/3", "1/2", map[string]any{"1/2": []any{"1", "1"}}},
		weighted:     true,
		thold: []any{
			[]any{
				map[string]any{
					"1/3": []any{
						big.NewRat(1, 2),
						big.NewRat(1, 2),
						big.NewRat(1, 2),
					},
				},
				big.NewRat(1, 3),
				big.NewRat(1, 2),
				map[string]any{
					"1/2": []any{
						big.NewRat(1, 1),
						big.NewRat(1, 1),
					},
				},
			},
		},
		limen: types.Qb64("4AAIA1s3k1s2v1s2v1s2c1s3c1s2c1s2k1v1"),
		size:  7,
		satisfactory: [][]int{
			{0, 2, 3, 6},
			{3, 4, 5},
			{1, 2, 3, 4},
			{4, 6},
			{4, 2, 0, 3},
			{0, 0, 1, 2, 1, 5, 6, 3},
		},
		unsatisfactory: [][]int{
			{0, 2, 5},
			{2, 3, 4},
		},
	},
	{
		sith:         "[[{\"1/3\":[\"1/2\", \"1/2\", \"1/2\"]}, \"1/2\", {\"1/2\": [\"1\", \"1\"]}], [\"1/2\", {\"1/2\": [\"1\", \"1\"]}]]",
		expectedSith: []any{[]any{map[string]any{"1/3": []any{"1/2", "1/2", "1/2"}}, "1/2", map[string]any{"1/2": []any{"1", "1"}}}, []any{"1/2", map[string]any{"1/2": []any{"1", "1"}}}},
		weighted:     true,
		thold: []any{
			[]any{
				map[string]any{
					"1/3": []any{
						big.NewRat(1, 2),
						big.NewRat(1, 2),
						big.NewRat(1, 2),
					},
				},
				big.NewRat(1, 2),
				map[string]any{
					"1/2": []any{
						big.NewRat(1, 1),
						big.NewRat(1, 1),
					},
				},
			},
			[]any{
				big.NewRat(1, 2),
				map[string]any{
					"1/2": []any{
						big.NewRat(1, 1),
						big.NewRat(1, 1),
					},
				},
			},
		},
		limen: types.Qb64("4AAKA1s3k1s2v1s2v1s2c1s2c1s2k1v1a1s2c1s2k1v1"),
		size:  9,
		satisfactory: [][]int{
			{0, 2, 3, 5, 6, 7},
			{3, 4, 5, 6, 8},
			{1, 2, 3, 4, 6, 7},
			{4, 2, 0, 3, 8, 6},
			{0, 0, 1, 2, 1, 8, 3, 5, 6, 3},
		},
		unsatisfactory: [][]int{
			{0, 2, 5},
			{6, 7, 8},
		},
	},
}

func indices(input []int) []types.Index {
	indices := []types.Index{}
	for _, i := range input {
		indices = append(indices, types.Index(i))
	}
	return indices
}

func TestTholderSiths(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			tholder, err := cesr.NewTholder(nil, nil, testCase.sith)
			if err != nil {
				t.Fatalf("error creating tholder: %v", err)
			}

			if tholder.Weighted() != testCase.weighted {
				t.Fatalf("expected weighted: %v, got: %v", testCase.weighted, tholder.Weighted())
			}

			sith, err := tholder.Sith()
			if err != nil {
				t.Fatalf("error getting sith: %v", err)
			}
			if !deepEqual(sith, testCase.expectedSith) {
				t.Fatalf("expected sith: %v, got: %v", testCase.expectedSith, sith)
			}

			limen, err := tholder.Limen()
			if err != nil {
				t.Fatalf("error getting limen: %v", err)
			}

			if limen != testCase.limen {
				t.Fatalf("expected limen: %v, got: %v", testCase.limen, limen)
			}

			if tholder.Size() != testCase.size {
				t.Fatalf("expected size: %d, got: %d", testCase.size, tholder.Size())
			}

			if !deepEqual(tholder.Thold(), testCase.thold) {
				t.Fatalf("expected thold: %v, got: %v", testCase.thold, tholder.Thold())
			}

			for _, satisfactory := range testCase.satisfactory {
				indices := indices(satisfactory)
				if !tholder.Satisfy(indices) {
					t.Fatalf("expected satisfactory: %v, got: %v", satisfactory, tholder.Satisfy(indices))
				}
			}

			for _, unsatisfactory := range testCase.unsatisfactory {
				indices := indices(unsatisfactory)
				if tholder.Satisfy(indices) {
					t.Fatalf("expected unsatisfactory: %v, got: %v", unsatisfactory, tholder.Satisfy(indices))
				}
			}
		})
	}
}

func TestTholderLimens(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			tholder, err := cesr.NewTholder(nil, &testCase.limen, nil)
			if err != nil {
				t.Fatalf("error creating tholder: %v", err)
			}

			if tholder.Weighted() != testCase.weighted {
				t.Fatalf("expected weighted: %v, got: %v", testCase.weighted, tholder.Weighted())
			}

			sith, err := tholder.Sith()
			if err != nil {
				t.Fatalf("error getting sith: %v", err)
			}
			if !deepEqual(sith, testCase.expectedSith) {
				t.Fatalf("expected sith: %v, got: %v", testCase.expectedSith, sith)
			}

			limen, err := tholder.Limen()
			if err != nil {
				t.Fatalf("error getting limen: %v", err)
			}

			if limen != testCase.limen {
				t.Fatalf("expected limen: %v, got: %v", testCase.limen, limen)
			}

			if tholder.Size() != testCase.size {
				t.Fatalf("expected size: %d, got: %d", testCase.size, tholder.Size())
			}

			if !deepEqual(tholder.Thold(), testCase.thold) {
				t.Fatalf("expected thold: %v, got: %v", testCase.thold, tholder.Thold())
			}

			for _, satisfactory := range testCase.satisfactory {
				indices := indices(satisfactory)
				if !tholder.Satisfy(indices) {
					t.Fatalf("expected satisfactory: %v, got: %v", satisfactory, tholder.Satisfy(indices))
				}
			}

			for _, unsatisfactory := range testCase.unsatisfactory {
				indices := indices(unsatisfactory)
				if tholder.Satisfy(indices) {
					t.Fatalf("expected unsatisfactory: %v, got: %v", unsatisfactory, tholder.Satisfy(indices))
				}
			}
		})
	}
}

func TestTholderTholds(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			tholder, err := cesr.NewTholder(testCase.thold, nil, nil)
			if err != nil {
				t.Fatalf("error creating tholder: %v", err)
			}

			if tholder.Weighted() != testCase.weighted {
				t.Fatalf("expected weighted: %v, got: %v", testCase.weighted, tholder.Weighted())
			}

			sith, err := tholder.Sith()
			if err != nil {
				t.Fatalf("error getting sith: %v", err)
			}
			if !deepEqual(sith, testCase.expectedSith) {
				t.Fatalf("expected sith: %v, got: %v", testCase.expectedSith, sith)
			}

			limen, err := tholder.Limen()
			if err != nil {
				t.Fatalf("error getting limen: %v", err)
			}

			if limen != testCase.limen {
				t.Fatalf("expected limen: %v, got: %v", testCase.limen, limen)
			}

			if tholder.Size() != testCase.size {
				t.Fatalf("expected size: %d, got: %d", testCase.size, tholder.Size())
			}

			if !deepEqual(tholder.Thold(), testCase.thold) {
				t.Fatalf("expected thold: %v, got: %v", testCase.thold, tholder.Thold())
			}

			for _, satisfactory := range testCase.satisfactory {
				indices := indices(satisfactory)
				if !tholder.Satisfy(indices) {
					t.Fatalf("expected satisfactory: %v, got: %v", satisfactory, tholder.Satisfy(indices))
				}
			}

			for _, unsatisfactory := range testCase.unsatisfactory {
				indices := indices(unsatisfactory)
				if tholder.Satisfy(indices) {
					t.Fatalf("expected unsatisfactory: %v, got: %v", unsatisfactory, tholder.Satisfy(indices))
				}
			}
		})
	}
}

func deepEqual(a, b any) bool {
	switch aType := a.(type) {
	case map[string]any:
		bType, ok := b.(map[string]any)
		if !ok {
			return false
		}
		if len(aType) != len(bType) {
			return false
		}
		for k, v := range aType {
			if !deepEqual(v, bType[k]) {
				return false
			}
		}
		return true
	case []any:
		bType, ok := b.([]any)
		if !ok {
			return false
		}
		if len(aType) != len(bType) {
			return false
		}
		for i := range aType {
			if !deepEqual(aType[i], bType[i]) {
				return false
			}
		}
		return true
	case *big.Rat:
		bType, ok := b.(*big.Rat)
		if !ok {
			return false
		}
		return aType.Cmp(bType) == 0
	default:
		return a == b
	}
}
