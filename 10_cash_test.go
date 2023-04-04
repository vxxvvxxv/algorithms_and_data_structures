package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var counterCash = 0

func TestCash(t *testing.T) {
	tcc := []struct {
		nums        int
		want        int
		wantCounter int
	}{
		{nums: 5, want: 120, wantCounter: 4},
		{nums: 5, want: 120, wantCounter: 0},
		{nums: 5, want: 120, wantCounter: 0},
		{nums: 7, want: 5040, wantCounter: 6},
		{nums: 8, want: 40320, wantCounter: 7},
		{nums: 7, want: 5040, wantCounter: 0},
		{nums: 8, want: 40320, wantCounter: 0},
	}

	fnWithCash := factorialWithCash(factorial)

	for _, tc := range tcc {
		t.Run(fmt.Sprintf("%v => %v", tc.nums, tc.want), func(t *testing.T) {
			counterCash = 0

			got := fnWithCash(tc.nums)

			fmt.Println("counter:", counterCash)

			assert.Equal(t, got, tc.want)
			assert.Equal(t, counterCash, tc.wantCounter)
		})
	}
}

func factorialWithCash(fn func(int) int) func(int) int {
	cash := make(map[int]int)
	return func(n int) int {
		if val, ok := cash[n]; ok {
			return val
		}

		result := fn(n)
		cash[n] = result

		return result
	}
}

func factorial(n int) int {
	result := 1

	for n != 1 {
		counterCash++

		result *= n
		n -= 1
	}

	return result
}
