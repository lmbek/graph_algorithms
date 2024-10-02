package graph

import (
	"graph/edge"
	"graph/vertex"
)

type Graph interface {
	GetVerticesCount() int
	GetEdgesCount() int
	GetVertices() map[string]vertex.Vertex
	GetEdges() map[string]edge.Edge
}

type graph struct {
	vertices map[string]vertex.Vertex
	edges    map[string]edge.Edge
}

func NewGraph(vertices map[string]vertex.Vertex, edges map[string]edge.Edge) Graph {
	return &graph{vertices: vertices, edges: edges}
}

func (g *graph) GetVerticesCount() int {
	return len(g.vertices)
}

func (g *graph) GetEdgesCount() int {
	return len(g.edges)
}

func (g *graph) GetVertices() map[string]vertex.Vertex {
	return g.vertices
}

func (g *graph) GetEdges() map[string]edge.Edge {
	return g.edges
}
