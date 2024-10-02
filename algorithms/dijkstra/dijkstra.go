package dijkstra

import (
	"graph"
	"graph/vertex"
)

// Run takes in a graph, with a specified start and end Vertex
// if it finds a path, it returns the Path, otherwise it returns nil
func Run(graph graph.Graph, start vertex.Vertex, end vertex.Vertex) *Path {
	return FindShortestPath(graph, start, end)
}
