package shapes_interfaces

import "math"

type Circle struct {
	radius float64
}

func NewCircle(measures Measures) *Circle  {
	return &Circle{radius:measures.Radius}
}

func (*Circle)GetShapeName() string  {
	return "Circle"
}

func (circle *Circle) Area() float64 {
	return math.Pi*circle.radius*circle.radius
}

func (circle *Circle) Perimeter() float64  {
	return 2*math.Pi*circle.radius
}
