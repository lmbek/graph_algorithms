package graph

import (
	"graph_algorithms/graph/edge"
	"graph_algorithms/graph/vertex"
)

type Graph interface {
	AddVertex(vertex vertex.Vertex) vertex.Vertex
	AddEdge(id string, edge edge.Edge) edge.Edge
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
	//g.vertices = make(map[string]vertex.Vertex)
	//g.edges = make(map[string]edge.Edge)
	g := graph{vertices: vertices, edges: edges}
	return &g
}

func (g *graph) AddVertex(vertex vertex.Vertex) vertex.Vertex {
	g.vertices[vertex.GetId()] = vertex
	return vertex
}

func (g *graph) AddEdge(id string, edge edge.Edge) edge.Edge {
	g.edges[id] = edge
	return edge
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
