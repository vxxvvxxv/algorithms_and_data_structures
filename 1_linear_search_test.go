package algorithms

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	tcc := []struct {
		nums []int
		find int
		want int
	}{
		{nums: []int{1, 5, 7, 2, 6, 8, 11, 75}, find: 5, want: 1},
		{nums: []int{1, 5, 7, 2, 6, 8, 11, 75}, find: 11, want: 6},
		{nums: []int{1, 5, 7, 2, 6, 8, 11, 75}, find: 75, want: 7},
		{nums: []int{0, 0, 0}, find: 1, want: -1},
		{nums: []int{}, find: 1, want: -1},
	}

	for _, tc := range tcc {
		t.Run(fmt.Sprintf("%v => %d", tc.nums, tc.find), func(t *testing.T) {
			got := LinearSearch(tc.nums, tc.find)
			if got != tc.want {
				t.Errorf("Error: got %d, want %d", got, tc.want)
			} else {
				t.Logf("OK: got %d, want %d", got, tc.want)
			}
		})
	}
}

func LinearSearch(nums []int, find int) int {
	fmt.Println("nums:", nums)
	fmt.Println("find:", find)

	counter := 0
	defer func() {
		fmt.Println("counter:", counter)
	}()

	for i := range nums {
		counter++

		if nums[i] == find {
			return i
		}
	}

	return -1
}
