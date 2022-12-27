package xcontainer

import (
	"math/rand"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestConcurrentSet(t *testing.T) {
	t.Parallel()

	t.Run("New", func(t *testing.T) {
		t.Parallel()

		type testCase[O constraints.Ordered] struct {
			name         string
			input        []O
			wantedLength int
		}

		testCases := []testCase[int]{
			{
				name:         "AllUnique",
				input:        []int{1, 2, 3},
				wantedLength: 3,
			},
			{
				name:         "HasDuplicates",
				input:        []int{1, 2, 3, 1},
				wantedLength: 3,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				set := NewConcurrentSet(tc.input...)
				if set.Len() != tc.wantedLength {
					t.Errorf("got %d, wanted %d", set.Len(), tc.wantedLength)
				}
			})
		}
	})

	t.Run("Contains", func(t *testing.T) {
		t.Parallel()

		type testCase[O constraints.Ordered] struct {
			name          string
			containsInput O
			containsWant  bool
		}

		cs := NewConcurrentSet([]int{1, 2, 3, 4, 5}...)

		intCases := []testCase[int]{
			{
				name:          "SuccessfullyFound",
				containsInput: 5,
				containsWant:  true,
			},
		}

		for _, tc := range intCases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				if got := cs.Contains(tc.containsInput); got != tc.containsWant {
					t.Errorf("got %v, want %v", got, tc.containsWant)
				}
			})
		}
	})
}

func BenchmarkConcurrentSetContains(b *testing.B) {
	containsCs := NewConcurrentSet([]int{1, 2, 3, 4, 5}...)

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			containsCs.Contains(rand.Intn(1000))
		}
	})
}

func BenchmarkConcurrentSetAdd(b *testing.B) {
	addCs := NewConcurrentSet([]int{1, 2, 3, 4, 5}...)

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			addCs.Add(rand.Intn(1000))
		}
	})
}
