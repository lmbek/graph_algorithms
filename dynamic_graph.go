package graph

import (
	"graph/edge"
	"graph/vertex"
)

type DynamicGraph interface {
	Graph
	AddVertex(vertex vertex.Vertex) vertex.Vertex
	AddEdge(id string, edge edge.Edge) edge.Edge
}

func NewDynamicGraph(vertices map[string]vertex.Vertex, edges map[string]edge.Edge) DynamicGraph {
	return &graph{vertices: vertices, edges: edges}
}

func (g *graph) AddVertex(vertex vertex.Vertex) vertex.Vertex {
	g.vertices[vertex.GetId()] = vertex
	return vertex
}

func (g *graph) AddEdge(id string, edge edge.Edge) edge.Edge {
	g.edges[id] = edge
	return edge
}
