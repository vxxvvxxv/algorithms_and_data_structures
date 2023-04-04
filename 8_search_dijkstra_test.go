package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var graphSearchInDeepDijkstra = map[string]map[string]int{
	"a": {"b": 2, "c": 1},
	"b": {"f": 7},
	"c": {"d": 5, "e": 2},
	"d": {"f": 2},
	"e": {"f": 1},
	"f": {"g": 1},
	"g": {},
}
var expectedSearchInDeepDijkstra = map[string]int{
	"b": 2,
	"c": 1,
	"d": 6,
	"e": 3,
	"f": 4,
	"g": 5,
}

func TestSearchInDeepDijkstra(t *testing.T) {
	tcc := []struct {
		have  map[string]map[string]int
		start string
		end   string
		want  map[string]int
	}{
		{have: graphSearchInDeepDijkstra, start: "a", end: "g", want: expectedSearchInDeepDijkstra},
	}

	for _, tc := range tcc {
		t.Run(fmt.Sprintf("%v => %v", tc.have, tc.want), func(t *testing.T) {
			got := SearchInDeepDijkstra(tc.have, tc.start, tc.end)

			assert.Equal(t, got, tc.want)
		})
	}
}

func SearchInDeepDijkstra(graph map[string]map[string]int, start, end string) map[string]int {
	counter := 0
	defer func() {
		fmt.Println("counter:", counter)
	}()

	costs := make(map[string]int)
	processed := make(map[string]struct{})
	neighbors := make(map[string]int)

	for node := range graph {
		if node == start {
			continue
		}

		if nodeCoast, ok := graph[start][node]; !ok {
			costs[node] = 999999
		} else {
			costs[node] = nodeCoast
		}

	}

	fmt.Println("costs:", costs)

	node := findLowerCoast(costs, processed)

	for node != "" {
		coast := costs[node]
		neighbors = graph[node]

		for neighbor := range neighbors {
			if _, ok := graph[node][neighbor]; !ok {
				continue
			}

			newCost := coast + graph[node][neighbor]

			if newCost < costs[neighbor] {
				costs[neighbor] = newCost
			}
		}

		processed[node] = struct{}{}
		node = findLowerCoast(costs, processed)
	}

	return costs
}

func findLowerCoast(costs map[string]int, processed map[string]struct{}) (lowerCoastNode string) {
	lowerCoast := 999999

	for node, coast := range costs {
		if _, ok := processed[node]; ok {
			continue
		}

		if coast < lowerCoast {
			lowerCoast = coast
			lowerCoastNode = node
		}
	}

	return lowerCoastNode
}
