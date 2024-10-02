package dijkstra

import (
	"fmt"
	"graph"
	"graph/edge"
	"graph/vertex"
	"graph/weight"
	"testing"
	"time"
)

func TestFindShortestPathOptimized_Exists(t *testing.T) {
	g, start, end := UseTestGraph()

	path := FindShortestPathOptimized(g, start, end)

	if path == nil {
		t.Errorf("path is nil")
		return
	}

	for _, v := range path.Vertices {
		if v == end {
			fmt.Print(v.GetId())
			break
		}
		fmt.Print(v.GetId(), "->")
	}

	fmt.Println()
}

func TestFindShortestPathOptimized_DoesNotExist(t *testing.T) {
	a := vertex.NewVertex("A")
	b := vertex.NewVertex("B")
	c := vertex.NewVertex("C")

	g := graph.NewGraph(map[string]vertex.Vertex{
		a.GetId(): a,
		b.GetId(): b,
		c.GetId(): c,
	}, map[string]edge.Edge{
		"1": edge.NewEdge(a, b, weight.NewWeight(1, 3*time.Second, true)),
	})

	start := a
	end := c

	path := FindShortestPathOptimized(g, start, end)

	if path != nil {
		t.Errorf("path is not nil")
	}
}

func TestReverseIterationOptimized(t *testing.T) {
	a := vertex.NewVertex("A")
	b := vertex.NewVertex("B")

	path := reverseIterationOptimized(b, a)

	if path != nil {
		t.Errorf("path is nil")
	}
}

func BenchmarkFindShortestPathOptimized_PathExist(bench *testing.B) {
	g, start, end := UseTestGraph()

	// Benchmark the FindShortestPathOptimized
	for i := 0; i < bench.N; i++ {
		FindShortestPathOptimized(g, start, end)
	}
}
