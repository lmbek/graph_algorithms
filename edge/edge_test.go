package edge

import (
	"graph/vertex"
	"graph/weight"
	"testing"
	"time"
)

func TestNewEdge(t *testing.T) {
	e := NewEdge(
		vertex.NewVertex("A"),
		vertex.NewVertex("B"),
		weight.NewWeight(5, 3*time.Second+357*time.Millisecond, false),
	)

	if e == nil {
		t.Errorf("edge e was nil")
		return
	}

	if e.GetFrom().GetId() != "A" || e.GetTo().GetId() != "B" {
		t.Errorf("the edges to and from vertices are not what was expected.")
		return
	}

	if e.GetWeight().GetDistance() != 5 || e.GetWeight().GetTime() != 3*time.Second+357*time.Millisecond || e.GetWeight().GetIsOpen() != false {
		t.Errorf("the edges weight is not was was expected.")
		return
	}
}
