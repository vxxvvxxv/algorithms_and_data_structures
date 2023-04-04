package algorithms

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tcc := []struct {
		nums []int
		want []int
	}{
		{nums: []int{}, want: []int{}},
		{nums: []int{6, 4, 3, 1, 7, 8, 4, 4, 3, 9}, want: []int{1, 3, 3, 4, 4, 4, 6, 7, 8, 9}},
	}

	for _, tc := range tcc {
		t.Run(fmt.Sprintf("%v => %v", tc.nums, tc.want), func(t *testing.T) {
			got := BubbleSortON2(tc.nums)
			if fmt.Sprintf("%+v", got) != fmt.Sprintf("%+v", tc.want) {
				t.Errorf("Error: got %+v, want %+v", got, tc.want)
			} else {
				t.Logf("OK: got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func BubbleSortON2(nums []int) []int {
	fmt.Println("nums:", nums)

	counter := 0
	defer func() {
		fmt.Println("counter:", counter)
	}()

	if len(nums) < 2 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		counter++

		for j := 0; j < len(nums)-1; j++ {
			counter++

			if nums[j+1] < nums[j] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}

	return nums
}
