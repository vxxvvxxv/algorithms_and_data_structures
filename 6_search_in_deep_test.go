package algorithms

import (
	"fmt"
	"testing"
)

var graphSearchInDeep = map[string][]string{
	"a": {"b", "c"},
	"b": {"f"},
	"c": {"d", "e"},
	"d": {"f"},
	"e": {"f"},
	"f": {"g"},
}

func TestSearchInDeep(t *testing.T) {
	tcc := []struct {
		have  map[string][]string
		start string
		end   string
		want  bool
	}{
		{have: graphSearchInDeep, start: "a", end: "g", want: true},
	}

	for _, tc := range tcc {
		t.Run(fmt.Sprintf("%v => %v", tc.have, tc.want), func(t *testing.T) {
			got := SearchInDeep(tc.have, tc.start, tc.end)

			if fmt.Sprintf("%+v", got) != fmt.Sprintf("%+v", tc.want) {
				t.Errorf("Error: got %+v, want %+v", got, tc.want)
			} else {
				t.Logf("OK: got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func SearchInDeep(graph map[string][]string, start, end string) bool {
	counter := 0
	defer func() {
		fmt.Println("counter:", counter)
	}()

	queue := []string{start}

	for len(queue) > 0 {
		counter++

		node := queue[0]  // get first element: start => a
		queue = queue[1:] // remove first element: start => []

		if graph[node] == nil {
			continue
		}

		// check if node is end
		if node == end {
			return true
		}

		for _, pathLink := range graph[node] {
			counter++

			if pathLink == end {
				return true
			}
		}

		// add node to queue
		// queue = append(queue, node) => [a]
		queue = append(queue, graph[node]...)
	}

	return false
}
