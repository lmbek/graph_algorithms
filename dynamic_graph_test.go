package graph

import (
	"graph/edge"
	"graph/vertex"
	"graph/weight"
	"testing"
	"time"
)

func TestNewDynamicGraph_AddVertex(t *testing.T) {
	g := NewDynamicGraph(map[string]vertex.Vertex{}, map[string]edge.Edge{})

	a := g.AddVertex(vertex.NewVertex("A"))
	b := g.AddVertex(vertex.NewVertex("B"))

	if a == nil {
		t.Errorf("vertex a is nil")
	}

	if b == nil {
		t.Errorf("vertex b is nil")
	}
}

func TestNewDynamicGraph_AddEdge(t *testing.T) {
	g := NewDynamicGraph(map[string]vertex.Vertex{}, map[string]edge.Edge{})
	a := g.AddVertex(vertex.NewVertex("A"))
	b := g.AddVertex(vertex.NewVertex("B"))

	e := g.AddEdge("1", edge.NewEdge(a, b, weight.NewWeight(5, 3*time.Second, true)))
	f := g.AddEdge("2", edge.NewEdge(b, a, weight.NewWeight(5, 3*time.Second, true)))

	if e == nil {
		t.Errorf("edge e is nil")
	}

	if f == nil {
		t.Errorf("edge f is nil")
	}
}
