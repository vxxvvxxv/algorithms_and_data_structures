package algorithms

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type treeStruct struct {
	Value    int          `json:"v"`
	Children []treeStruct `json:"c"`
}

//go:embed 9_tree.json
var treeBytes []byte
var treeList []treeStruct

var treeCounterByRecursive int
var treeCounterByStack int

func init() {
	if err := json.Unmarshal(treeBytes, &treeList); err != nil {
		panic(err)
	}
}

func TestTree(t *testing.T) {
	tcc := []struct {
		have []treeStruct
		want int
	}{
		{have: treeList, want: 69},
	}

	for _, tc := range tcc {
		t.Run(fmt.Sprintf("%v => %v", tc.have, tc.want), func(t *testing.T) {
			got := treeResultByRecursive(tc.have)

			fmt.Println("counter:", treeCounterByRecursive)

			assert.Equal(t, got, tc.want)
		})

		t.Run(fmt.Sprintf("%v => %v", tc.have, tc.want), func(t *testing.T) {
			got := treeResultByStack(tc.have)

			fmt.Println("counter:", treeCounterByStack)

			assert.Equal(t, got, tc.want)
		})
	}
}

func treeResultByRecursive(tree []treeStruct) int {
	sum := 0

	for _, node := range tree {
		treeCounterByRecursive++

		sum += node.Value

		if len(node.Children) == 0 {
			continue
		}

		sum += treeResultByRecursive(node.Children)
	}

	return sum
}

func treeResultByStack(tree []treeStruct) int {
	if len(tree) == 0 {
		return 0
	}

	sum := 0
	stack := make([]treeStruct, 0)

	for _, node := range tree {
		treeCounterByStack++

		stack = append(stack, node)
	}

	var node treeStruct

	for len(stack) > 0 {
		treeCounterByStack++

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		sum += node.Value

		if len(node.Children) == 0 {
			continue
		}

		for _, child := range node.Children {
			treeCounterByStack++

			stack = append(stack, child)
		}
	}

	return sum
}
