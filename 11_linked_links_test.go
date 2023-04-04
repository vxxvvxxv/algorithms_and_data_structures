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

func createList(nums []int) *List {
	list := &List{}

	var currentNode *ListNode
	var prevNode *ListNode

	for i, num := range nums {
		currentNode = &ListNode{
			Val: num,
		}

		// Fill the start point
		if i == 0 {
			list.startNode = currentNode
		} else {
			prevNode.Next = currentNode
		}

		prevNode = currentNode

		list.count++
	}

	return list
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
			gotList := createList(tc.nums)
			fmt.Println("gotList.count:", gotList.Count())

			assert.Equal(t, len(tc.nums), gotList.Count())
			assert.Equal(t, len(tc.nums), gotList.CalculateCount())
		})
	}
}
