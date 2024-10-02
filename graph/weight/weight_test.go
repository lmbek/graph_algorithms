package weight

import (
	"testing"
	"time"
)

func TestNewWeight_GetSetDistance(t *testing.T) {
	w := NewWeight(0, 0, false)

	if w.GetDistance() != 0 {
		t.Errorf("distance is not 0")
		return
	}

	w.SetDistance(15)

	if w.GetDistance() != 15 {
		t.Errorf("distance is not 15")
		return
	}
}

func TestWeight_GetSetTime(t *testing.T) {
	w := NewWeight(0, 0, false)

	w.SetTime(5 * time.Second)

	if w.GetTime() != 5*time.Second {
		t.Errorf("the time was different than expected.")
		return
	}
}

func TestWeight_GetSetIsOpen(t *testing.T) {
	w := NewWeight(0, 0, false)

	if w.GetIsOpen() == true {
		t.Errorf("weight is true")
	}

	w.SetIsOpen(true)

	if w.GetIsOpen() == false {
		t.Errorf("weight is false")
	}
}

func TestWeight_AddToWeight(t *testing.T) {
	w1 := NewWeight(1, 5*time.Second, false)
	w2 := NewWeight(1, 5*time.Second, false)

	w1.AddToWeight(w2)

	expectedWeight := NewWeight(2, 10*time.Second, false)

	if w1.GetDistance() != expectedWeight.GetDistance() || w1.GetTime() != expectedWeight.GetTime() || w1.GetIsOpen() != expectedWeight.GetIsOpen() {
		t.Errorf("the weight was not same as expected.")
	}
}

func TestWeight_CompareWithWeight(t *testing.T) {
	w1 := NewWeight(1, 5*time.Second, false)
	w2 := NewWeight(2, 10*time.Second, true)

	// test if lower
	if w2.CompareWithWeight(w1) == true {
		t.Errorf("weight 2 is higher than weight 1")
	}

	// test if higher
	if w1.CompareWithWeight(w2) == false {
		t.Errorf("weight 2 is lower than weight 1")
	}
}
