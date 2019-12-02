package shapes_interfaces

type Square struct {
	length float64
}

func NewSquares(measures Measures)*Square  {
	return &Square{length:measures.Length}
}

func (*Square)GetShapeName()string  {
	return "Square"
}

func (square *Square)Area() float64  {
	return square.length*square.length
}

func (sqaure *Square)Perimeter() float64  {
	return 4*sqaure.length
}

