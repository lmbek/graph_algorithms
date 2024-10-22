package dijkstra

import (
	"graph"
	"graph/edge"
	"graph/vertex"
	"graph/weight"
	"math"
	"time"
)

// TODO: OPTIMIZE WITH PRIORITY QUEUE

// FindShortestPathOptimized finds the shortest path using Dijkstra's algorithm with a time complexity of O(n^2) in dense graphs.
func FindShortestPathOptimized(graph graph.Graph, start vertex.Vertex, end vertex.Vertex) *Path {

	var queue = make(map[string]vertex.Vertex)

	for _, v := range graph.GetVertices() {

		v.SetWeight(weight.NewWeight(
			math.Inf(1),
			time.Duration(math.Inf(1)),
			false,
		))
		v.SetPrevious(nil)

		queue[v.GetId()] = v
	}

	// set weight on start vertex to 0
	start.SetWeight(weight.NewWeight(0, 0, true))

	for len(queue) != 0 {
		// get the vertex with the smallest distance
		u := smallestDistanceOptimized(queue)

		// if the smallest distance is infinite, no more reachable vertices
		if math.IsInf(u.GetWeight().GetDistance(), 1) {
			break
		}

		// check if end has been reached
		if u == end {
			return reverseIterationOptimized(u, start)
		}

		delete(queue, u.GetId())

		neighbors := findNeighborsOfUInQOptimized(graph, u, queue)

		edges := graph.GetEdges()

		for _, v := range neighbors {
			var selectedEdge edge.Edge
			for _, edge := range edges {
				if edge.GetFrom() == u && edge.GetTo() == v {
					selectedEdge = edge
					break
				}
			}

			var pathWeight weight.Weight = weight.NewWeight(0, 0, false)
			pathWeight.AddToWeight(u.GetWeight())
			if selectedEdge != nil {
				pathWeight.AddToWeight(selectedEdge.GetWeight())
			}

			if pathWeight.CompareWithWeight(v.GetWeight()) {
				v.SetWeight(pathWeight)
				v.SetPrevious(u)
			}
		}
	}

	return nil
}

// reverseIterationOptimized builds the path from the end vertex back to the start by following each vertex's "previous" link.
// It starts from the end and keeps going backward through the vertices, adding each one to the beginning of the list.
// The result is the shortest path from start to end, or nil if no valid path exists.
func reverseIterationOptimized(u vertex.Vertex, start vertex.Vertex) *Path {
	var sequence []vertex.Vertex

	if u.GetPrevious() != nil || u == start {
		for u != nil {
			// prepend to sequence
			sequence = append([]vertex.Vertex{u}, sequence...)
			u = u.GetPrevious()
		}
		return &Path{sequence}
	}

	return nil
}

// findNeighborsOfUInQOptimized - finds the neighboring vertices of the given vertex u in the graph g, that are still in the queue q.
func findNeighborsOfUInQOptimized(g graph.Graph, u vertex.Vertex, q map[string]vertex.Vertex) []vertex.Vertex {
	// list of edges connected to u
	var edges []edge.Edge

	for _, e := range g.GetEdges() {
		// check that edge goes from u
		if e.GetFrom().GetId() == u.GetId() {
			edges = append(edges, e)
		}
	}

	// list of neighbors still in queue
	var neighbors []vertex.Vertex

	for _, e := range edges {
		// check that neighbor vertex is in queue
		_, exists := q[e.GetTo().GetId()]
		if exists {
			neighbors = append(neighbors, e.GetTo())
		}
	}

	return neighbors
}

// smallestDistanceOptimized - gets the vertex with the smallest distance from map
func smallestDistanceOptimized(queue map[string]vertex.Vertex) vertex.Vertex {
	var smallest = vertex.NewVertex("temp")

	smallest.SetWeight(weight.NewWeight(math.Inf(1), time.Duration(math.Inf(1)), true))

	for _, v := range queue {
		if v.GetWeight().CompareWithWeight(smallest.GetWeight()) {
			smallest = v
		}
	}

	return smallest
}
