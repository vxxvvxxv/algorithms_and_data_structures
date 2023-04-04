package algorithms

import (
	"fmt"
	"testing"
)

var binaryArray = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

var counterBinarySearch int

func TestBinarySearch(t *testing.T) {
	tcc := []struct {
		nums []int
		find int
		want int
	}{
		{nums: binaryArray, find: 5, want: 5},
		{nums: binaryArray, find: 15, want: 15},
		{nums: binaryArray, find: 0, want: 0},
		{nums: binaryArray, find: 16, want: -1},
	}

	for _, tc := range tcc {
		t.Run(fmt.Sprintf("%v => %d", tc.nums, tc.find), func(t *testing.T) {
			got := BinarySearchRecursive(tc.nums, tc.find, 0, len(tc.nums)-1)

			fmt.Println("counterBinarySearch:", counterBinarySearch)

			if got != tc.want {
				t.Errorf("Error: got %d, want %d", got, tc.want)
			} else {
				t.Logf("OK: got %d, want %d", got, tc.want)
			}
		})

		t.Run(fmt.Sprintf("%v => %d", tc.nums, tc.find), func(t *testing.T) {
			got := BinarySearch(tc.nums, tc.find)
			if got != tc.want {
				t.Errorf("Error: got %d, want %d", got, tc.want)
			} else {
				t.Logf("OK: got %d, want %d", got, tc.want)
			}
		})
	}
}

func BinarySearchRecursive(nums []int, find int, start, end int) int {
	middle := (start + end) / 2

	if nums[middle] == find {
		return middle
	}

	if start == end {
		return -1
	}

	if find < nums[middle] {
		return BinarySearchRecursive(nums, find, start, middle-1)
	} else {
		return BinarySearchRecursive(nums, find, middle+1, end)
	}
}

func BinarySearch(nums []int, find int) int {
	fmt.Println("nums:", nums)
	fmt.Println("find:", find)

	counter := 0
	defer func() {
		fmt.Println("counter:", counter)
	}()

	start := 0
	end := len(nums)
	middle := 0
	found := false
	position := -1

	for !found || start <= end {
		counter++

		middle = (start + end) / 2

		if middle >= len(nums) {
			break
		}

		if nums[middle] == find {
			found = true
			position = middle
			return middle
		}

		if find < nums[middle] {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return position
}
