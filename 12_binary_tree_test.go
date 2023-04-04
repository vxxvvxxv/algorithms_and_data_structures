package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type binaryTree struct {
	root *BinaryTreeNode
	size int
}

func newBinaryTree() *binaryTree {
	return &binaryTree{}
}

func (tree *binaryTree) Count() int {
	return tree.size
}

func (tree *binaryTree) Add(val int) {
	if tree.root == nil {
		tree.root = &BinaryTreeNode{Val: val}
		tree.size++
		return
	}

	tree.root.Add(val)
	tree.size++
}

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func (node *BinaryTreeNode) IsEmpty() bool {
	return node == nil
}

func (node *BinaryTreeNode) Count() int {
	if node.IsEmpty() {
		return 0
	}

	return 1 + node.Left.Count() + node.Right.Count()
}

func (node *BinaryTreeNode) Add(val int) {
	if node.IsEmpty() {
		return
	}

	if val < node.Val {
		if node.Left == nil {
			node.Left = &BinaryTreeNode{Val: val}
			return
		} else {
			node.Left.Add(val)
		}
	} else if val > node.Val {
		if node.Right == nil {
			node.Right = &BinaryTreeNode{Val: val}
			return
		} else {
			node.Right.Add(val)
		}
	}
}

func TestBinaryTree(t *testing.T) {
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
			gotBinaryTree := newBinaryTree()

			for _, num := range tc.nums {
				gotBinaryTree.Add(num)
			}

			assert.Equal(t, len(tc.nums), gotBinaryTree.Count())

			if len(tc.nums) != 0 {
				assert.Equal(t, len(tc.nums), gotBinaryTree.root.Count())
			}
		})
	}
}
