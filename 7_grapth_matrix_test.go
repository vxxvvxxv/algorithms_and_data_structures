package algorithms

import (
	"fmt"
	"testing"
)

// "a": {"b", "c"},
//
//	"b": {"f"},
//	"c": {"d", "e"},
//	"d": {"f"},
//	"e": {"f"},
//	"f": {"g"},
var graphSearchInDeepMatrixTable = [][]int{
	// 0, 1, 2, 3, 4, 5, 6
	{0, 1, 1, 0, 0, 0, 0}, // 0
	{0, 0, 0, 0, 1, 0, 0}, // 1
	{0, 0, 0, 1, 0, 1, 0}, // 2
	{0, 0, 0, 0, 1, 0, 0}, // 3
	{0, 0, 0, 0, 0, 0, 1}, // 4
	{0, 0, 0, 0, 1, 0, 0}, // 5
	{0, 0, 0, 0, 0, 0, 0}, // 6
}

func TestSearchInDeepMatrix(t *testing.T) {
	tcc := []struct {
		have  [][]int
		start int
		end   int
		want  bool
	}{
		{have: graphSearchInDeepMatrixTable, start: 0, end: 6, want: true},
	}

	for _, tc := range tcc {
		t.Run(fmt.Sprintf("%v => %v", tc.have, tc.want), func(t *testing.T) {
			got := SearchInDeepMatrix(tc.have, tc.start, tc.end)

			if fmt.Sprintf("%+v", got) != fmt.Sprintf("%+v", tc.want) {
				t.Errorf("Error: got %+v, want %+v", got, tc.want)
			} else {
				t.Logf("OK: got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func SearchInDeepMatrix(matrix [][]int, start, end int) bool {
	counter := 0
	defer func() {
		fmt.Println("counter:", counter)
	}()

	// TODO: convert graph to matrix

	return true
}
