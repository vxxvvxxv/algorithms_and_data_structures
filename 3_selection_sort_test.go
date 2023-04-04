package algorithms

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tcc := []struct {
		nums []int
		want []int
	}{
		{nums: []int{}, want: []int{}},
		{nums: []int{6, 4, 3, 1, 7, 8, 4, 4, 3, 9}, want: []int{1, 3, 3, 4, 4, 4, 6, 7, 8, 9}},
	}

	for _, tc := range tcc {
		t.Run(fmt.Sprintf("%v => %v", tc.nums, tc.want), func(t *testing.T) {
			got := SelectionSortON2(tc.nums)
			if fmt.Sprintf("%+v", got) != fmt.Sprintf("%+v", tc.want) {
				t.Errorf("Error: got %+v, want %+v", got, tc.want)
			} else {
				t.Logf("OK: got %+v, want %+v", got, tc.want)
			}
		})
	}

	for _, tc := range tcc {
		t.Run(fmt.Sprintf("%v => %v", tc.nums, tc.want), func(t *testing.T) {
			got := SelectionSortON2Back(tc.nums)
			if fmt.Sprintf("%+v", got) != fmt.Sprintf("%+v", tc.want) {
				t.Errorf("Error: got %+v, want %+v", got, tc.want)
			} else {
				t.Logf("OK: got %+v, want %+v", got, tc.want)
			}
		})
	}

	for _, tc := range tcc {
		t.Run(fmt.Sprintf("O%v => %v", tc.nums, tc.want), func(t *testing.T) {
			got := SelectionSortOn(tc.nums)
			if fmt.Sprintf("%+v", got) != fmt.Sprintf("%+v", tc.want) {
				t.Errorf("Error: got %+v, want %+v", got, tc.want)
			} else {
				t.Logf("OK: got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func SelectionSortON2(nums []int) []int {
	fmt.Println("nums:", nums)

	counter := 0
	defer func() {
		fmt.Println("counter:", counter)
	}()

	if len(nums) == 0 {
		return nums
	}

	min := 0
	for i := 0; i < len(nums); i++ {
		counter++
		min = nums[i]

		for j := i + 1; j < len(nums); j++ {
			counter++

			if min > nums[j] {
				min = nums[j]

				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}

func SelectionSortON2Back(nums []int) []int {
	fmt.Println("nums:", nums)

	counter := 0
	defer func() {
		fmt.Println("counter:", counter)
	}()

	if len(nums) < 2 {
		return nums
	}

	for i := len(nums) - 2; i >= 0; i-- {
		counter++

		for j := len(nums) - 1; j >= 0; j-- {
			counter++

			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}

func SelectionSortOn(nums []int) []int {
	fmt.Println("nums:", nums)

	counter := 0
	defer func() {
		fmt.Println("counter:", counter)
	}()

	if len(nums) < 2 {
		return nums
	}

	min := 0
	currentIdx := 0

	for i := 1; i < len(nums); i++ {
		counter++

		min = nums[currentIdx]

		if min > nums[i] {
			min = nums[i]

			nums[i], nums[currentIdx] = nums[currentIdx], nums[i]
		}

		currentIdx++
	}

	return nums
}
