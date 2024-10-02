package edge

import (
	"graph_algorithms/graph/vertex"
	"graph_algorithms/graph/weight"
)

type Edge interface {
	GetTo() vertex.Vertex
	GetFrom() vertex.Vertex
	GetWeight() weight.Weight
}

type edge struct {
	to     vertex.Vertex
	from   vertex.Vertex
	weight weight.Weight
}

func NewEdge(from vertex.Vertex, to vertex.Vertex, weight weight.Weight) Edge {
	return &edge{from: from, to: to, weight: weight}
}

func (edge *edge) GetTo() vertex.Vertex {
	return edge.to
}

func (edge *edge) GetFrom() vertex.Vertex {
	return edge.from
}

func (edge *edge) GetWeight() weight.Weight {
	return edge.weight
}
