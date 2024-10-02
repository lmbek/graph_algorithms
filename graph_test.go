package graph_test

import (
	"graph"
	"graph/edge"
	"graph/vertex"
	"graph/weight"
	"testing"
	"time"
)

func TestGraph_GetVertices(t *testing.T) {
	a := vertex.NewVertex("A")
	b := vertex.NewVertex("B")
	c := vertex.NewVertex("C")

	g := graph.NewGraph(map[string]vertex.Vertex{
		a.GetId(): a,
		b.GetId(): b,
		c.GetId(): c,
	}, nil)

	if g.GetVertices() == nil {
		t.Errorf("vertices is nil")
	}

	if g.GetVerticesCount() != 3 {
		t.Errorf("vertices is not 3")
	}

}

func TestGraph_GetEdges(t *testing.T) {
	a := vertex.NewVertex("A")
	b := vertex.NewVertex("B")
	c := vertex.NewVertex("C")

	g := graph.NewGraph(map[string]vertex.Vertex{
		a.GetId(): a,
		b.GetId(): b,
		c.GetId(): c,
	}, map[string]edge.Edge{
		"1": edge.NewEdge(a, b, weight.NewWeight(4, 3*time.Second, true)),
		"2": edge.NewEdge(a, c, weight.NewWeight(5, 42*time.Second, true)),
		"3": edge.NewEdge(b, a, weight.NewWeight(4, 42*time.Second, true)),
		"4": edge.NewEdge(b, c, weight.NewWeight(11, 42*time.Second, true)),
		"5": edge.NewEdge(c, a, weight.NewWeight(5, 30*time.Second, true)),
		"6": edge.NewEdge(c, b, weight.NewWeight(11, 13*time.Second, true)),
	})

	if g.GetEdges() == nil {
		t.Errorf("edges is nil")
	}

	if g.GetEdgesCount() != 6 {
		t.Errorf("edges is not 6")
	}
}

func TestNewGraph(t *testing.T) {
	a := vertex.NewVertex("A")
	b := vertex.NewVertex("B")
	c := vertex.NewVertex("C")
	d := vertex.NewVertex("D")
	e := vertex.NewVertex("E")
	f := vertex.NewVertex("F")

	g := graph.NewGraph(map[string]vertex.Vertex{
		a.GetId(): a,
		b.GetId(): b,
		c.GetId(): c,
		d.GetId(): d,
		e.GetId(): e,
		f.GetId(): f,
	}, map[string]edge.Edge{
		"1":  edge.NewEdge(a, b, weight.NewWeight(4, 3*time.Second, true)),
		"2":  edge.NewEdge(a, c, weight.NewWeight(5, 42*time.Second, true)),
		"3":  edge.NewEdge(b, a, weight.NewWeight(4, 42*time.Second, true)),
		"4":  edge.NewEdge(b, c, weight.NewWeight(11, 42*time.Second, true)),
		"5":  edge.NewEdge(b, d, weight.NewWeight(9, 42*time.Second, true)),
		"6":  edge.NewEdge(b, e, weight.NewWeight(7, 42*time.Second, true)),
		"7":  edge.NewEdge(c, a, weight.NewWeight(5, 30*time.Second, true)),
		"8":  edge.NewEdge(c, b, weight.NewWeight(11, 13*time.Second, true)),
		"9":  edge.NewEdge(c, e, weight.NewWeight(3, 120*time.Second, true)),
		"10": edge.NewEdge(d, b, weight.NewWeight(9, 11*time.Second, true)),
		"11": edge.NewEdge(d, e, weight.NewWeight(13, 120*time.Second, true)),
		"12": edge.NewEdge(d, f, weight.NewWeight(2, 3*time.Second, true)),
		"13": edge.NewEdge(e, b, weight.NewWeight(7, 3*time.Second, true)),
		"14": edge.NewEdge(e, c, weight.NewWeight(3, 209*time.Second, true)),
		"15": edge.NewEdge(e, d, weight.NewWeight(13, 50*time.Second, true)),
		"16": edge.NewEdge(e, f, weight.NewWeight(6, 5*time.Second, true)),
		"17": edge.NewEdge(f, d, weight.NewWeight(2, 49*time.Second, true)),
		"18": edge.NewEdge(f, e, weight.NewWeight(6, 67*time.Second, true)),
	})

	if g.GetVerticesCount() != 6 {
		t.Errorf("vertices is not 6")
	}

	if g.GetEdgesCount() != 18 {
		t.Errorf("edges is not 18")
	}
}
