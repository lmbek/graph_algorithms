package dijkstra

import (
	"graph/vertex"
	"testing"
)

func TestReverseIteration(t *testing.T) {
	a := vertex.NewVertex("A")
	b := vertex.NewVertex("B")

	path := reverseIteration(b, a)

	if path != nil {
		t.Errorf("path is nil")
	}
}
