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

func TestFindShortestPath_Exists(t *testing.T) {
	g, start, end := UseTestGraph()

	path := FindShortestPath(g, start, end)

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

func TestFindShortestPath_DoesNotExist(t *testing.T) {
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

	path := FindShortestPath(g, start, end)

	if path != nil {
		t.Errorf("path is not nil")
	}
}

func TestReverseIteration(t *testing.T) {
	a := vertex.NewVertex("A")
	b := vertex.NewVertex("B")

	path := reverseIteration(b, a)

	if path != nil {
		t.Errorf("path is nil")
	}
}

func BenchmarkFindShortestPath_PathExist(bench *testing.B) {
	g, start, end := UseTestGraph()

	// Benchmark the FindShortestPath
	for i := 0; i < bench.N; i++ {
		FindShortestPath(g, start, end)
	}
}
