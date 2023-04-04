package algorithms

import (
	"fmt"
	"testing"
)

var counterQuickSort int

func TestQuickSort(t *testing.T) {
	tcc := []struct {
		nums []int
		want []int
	}{
		{nums: []int{}, want: []int{}},
		{nums: []int{6, 4, 3, 1, 7, 8, 4, 4, 3, 9}, want: []int{1, 3, 3, 4, 4, 4, 6, 7, 8, 9}},
	}

	for _, tc := range tcc {
		t.Run(fmt.Sprintf("%v => %v", tc.nums, tc.want), func(t *testing.T) {
			got := QuickSort(tc.nums)

			fmt.Println("counter:", counterQuickSort)

			if fmt.Sprintf("%+v", got) != fmt.Sprintf("%+v", tc.want) {
				t.Errorf("Error: got %+v, want %+v", got, tc.want)
			} else {
				t.Logf("OK: got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	pivotIdx := len(nums) / 2
	pivot := nums[pivotIdx]

	left, right := make([]int, 0, cap(nums)), make([]int, 0, cap(nums))

	for i := 0; i < len(nums); i++ {
		counterQuickSort++

		if i == pivotIdx {
			continue
		}

		if nums[i] < pivot {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}
