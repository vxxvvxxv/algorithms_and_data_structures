package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// KeyType may be any type that is comparable (https://go.dev/ref/spec#Comparison_operators),
// and ValueType may be any type at all, including another map!

// This it's not comparable:
// map[func()]struct{}{}
// map[struct{s: []int}]struct{}{}

type mapStructKey struct{}
type mapStructKey2 struct {
	s int
}

func TestMapSet(t *testing.T) {
	tcc := []struct {
		have interface{}
		want interface{}
	}{
		{have: map[interface{}]struct{}{}, want: map[interface{}]struct{}{}},
		{have: map[mapStructKey]struct{}{}, want: map[mapStructKey]struct{}{}},
		{have: map[*mapStructKey]struct{}{}, want: map[*mapStructKey]struct{}{}},
		{have: map[mapStructKey2]struct{}{}, want: map[mapStructKey2]struct{}{}},
	}

	for i, tc := range tcc {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			assert.Equal(t, tc.want, tc.have)
		})
	}
}
