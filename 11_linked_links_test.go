package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type List struct {
	count     int
	startNode *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *List) Count() int {
	return list.count
}

func (list *List) CalculateCount() int {
	counter := 0

	if list.startNode == nil {
		return counter
	}

	node := list.startNode
	counter++

	for node.Next != nil {
		counter++
		node = node.Next
	}

	return counter
}

func (list *List) ToSlice() []int {
	result := make([]int, 0, list.Count())

	if list.startNode == nil {
		return result
	}

	node := list.startNode

	for i := 0; i < list.Count(); i++ {
		result = append(result, node.Val)

		if node.Next != nil {
			node = node.Next
		}
	}

	return result
}

func (list *List) AddNode(node *ListNode) {
	if list.startNode == nil {
		list.startNode = node
	} else {
		lastNode := list.startNode
		for lastNode.Next != nil {
			lastNode = lastNode.Next
		}

		lastNode.Next = node
	}

	list.count++
}

func (list *List) FillFromSlice(nums []int) {
	for _, num := range nums {
		list.AddNode(&ListNode{Val: num})
	}
}

func createList() *List {
	return &List{}
}

func TestListNode(t *testing.T) {
	tcc := []struct {
		nums []int
	}{
		{nums: []int{}},
		{nums: []int{1}},
		{nums: []int{1, 2, 3, 4}},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tc := range tcc {
		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
			gotList := createList()
			fmt.Println("gotList.count:", gotList.Count())
			gotList.FillFromSlice(tc.nums)

			assert.Equal(t, len(tc.nums), gotList.Count())
			assert.Equal(t, len(tc.nums), gotList.CalculateCount())

			assert.Equal(t, tc.nums, gotList.ToSlice())
		})
	}
}
