package vertex

import (
	"graph_algorithms/graph/weight"
	"testing"
	"time"
)

func TestNewVertex(t *testing.T) {
	v := NewVertex("A")

	if v == nil {
		t.Errorf("vertex was nil")
		return
	}

	if v.GetId() != "A" {
		t.Errorf("the vertex id was not A.")
		return
	}
}

func TestVertex_GetId(t *testing.T) {
	a := NewVertex("A")
	id := a.GetId()
	if id != "A" {
		t.Errorf("the vertex id was not A.")
	}

	b := NewVertex("B")
	id = b.GetId()
	if id != "B" {
		t.Errorf("the vertex id was not B.")
	}
}

func TestVertex_GetSetWeight(t *testing.T) {
	v := NewVertex("A")

	if v.GetWeight() == nil {
		t.Errorf("the vertex weight is nil")
		return
	}

	if v.GetWeight().GetDistance() != 0 {
		t.Errorf("the weight distance is not 0")
		return
	}

	var distance float64 = 2
	var timeDuration time.Duration = 2 * time.Second
	var isOpen bool = true

	v.SetWeight(weight.NewWeight(distance, timeDuration, isOpen))

	if v.GetWeight().GetDistance() != distance || v.GetWeight().GetTime() != timeDuration || v.GetWeight().GetIsOpen() != isOpen {
		t.Errorf("the vertex weight is not the same as expected values")
	}
}

func TestVertex_GetSetPrevious(t *testing.T) {
	v := NewVertex("A")

	if v.GetPrevious() != nil {
		t.Errorf("the previous vertex is already set")
		return
	}

	prev := NewVertex("B")
	v.SetPrevious(prev)

	if v.GetPrevious() != prev {
		t.Errorf("previous vertex is not the expected vertex")
	}
}
