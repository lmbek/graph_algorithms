package vertex

import (
	"graph_algorithms/graph/weight"
)

type Vertex interface {
	GetId() string
	GetWeight() weight.Weight
	SetWeight(weight weight.Weight)
	GetPrevious() Vertex
	SetPrevious(vertex Vertex)
}

type vertex struct {
	id       string
	weight   weight.Weight
	previous Vertex
}

func NewVertex(id string) Vertex {
	w := weight.NewWeight(0, 0, true)

	return &vertex{
		id:       id,
		weight:   w,
		previous: nil,
	}
}

func (v *vertex) GetId() string {
	return v.id
}

func (v *vertex) GetWeight() weight.Weight {
	return v.weight
}
func (v *vertex) SetWeight(weight weight.Weight) {
	v.weight = weight
}

func (v *vertex) GetPrevious() Vertex {
	return v.previous
}
func (v *vertex) SetPrevious(previous Vertex) {
	v.previous = previous
}
