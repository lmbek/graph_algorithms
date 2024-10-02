package weight

import "time"

type Weight interface {
	GetDistance() float64
	SetDistance(float64)
	GetTime() time.Duration
	SetTime(time.Duration)
	GetIsOpen() bool
	SetIsOpen(bool)
	AddToWeight(other Weight)
	CompareWithWeight(other Weight) bool
}

type weight struct {
	distance float64
	time     time.Duration
	isOpen   bool
}

func NewWeight(distance float64, time time.Duration, isOpen bool) Weight {
	return &weight{
		distance: distance,
		time:     time,
		isOpen:   isOpen,
	}
}

func (w *weight) GetDistance() float64 {
	return w.distance
}

func (w *weight) SetDistance(distance float64) {
	w.distance = distance
}

func (w *weight) GetTime() time.Duration {
	return w.time
}

func (w *weight) SetTime(time time.Duration) {
	w.time = time
}

func (w *weight) GetIsOpen() bool {
	return w.isOpen
}

func (w *weight) SetIsOpen(isOpen bool) {
	w.isOpen = isOpen
}

func (w *weight) AddToWeight(other Weight) {
	w.SetDistance(w.GetDistance() + other.GetDistance())
	w.SetTime(w.GetTime() + other.GetTime())
	w.SetIsOpen(w.GetIsOpen() && other.GetIsOpen())
}

// CompareWithWeight compares the weight
// If w (distance) is smaller than other (distance), returns true.
// If w (distance) is bigger than other (distance), returns false.
// TODO: (weighed score feature: Look for ways to make 60/40 decision making compares (60%/40% on distance and time)
func (w *weight) CompareWithWeight(other Weight) bool {
	return w.GetDistance() < other.GetDistance()
}
