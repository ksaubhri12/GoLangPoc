package shapes_interfaces

import "math"

type RightAngleTriangle struct {
	base, perpendicular float64
}

func NewRightAngleTriangle(measures Measures) *RightAngleTriangle {
	return &RightAngleTriangle{base: measures.Base, perpendicular: measures.Perpendicular}
}
func (*RightAngleTriangle)GetShapeName()string  {
	return "Right Angle Triangle"
}

func (rightAngleTriangle *RightAngleTriangle) Area() float64 {
	return 0.5 * rightAngleTriangle.base * rightAngleTriangle.perpendicular
}

func (rightAngleTriangle *RightAngleTriangle) Perimeter() float64  {
	hypotenues := math.Sqrt(rightAngleTriangle.perpendicular*rightAngleTriangle.perpendicular+
		rightAngleTriangle.base*rightAngleTriangle.base)
	return rightAngleTriangle.base + rightAngleTriangle.perpendicular + hypotenues
}


