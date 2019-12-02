package factory

import (
	"goLangTraining/19_interfaces/enums"
	"goLangTraining/19_interfaces/shapes_interfaces"
)

func ReturnShapeFactory(shape enums.Shapes , measures shapes_interfaces.Measures) shapes_interfaces.Shape  {
	switch shape {
	case enums.Square:
		return shapes_interfaces.NewSquares(measures)
	case enums.Circle:
		return shapes_interfaces.NewCircle(measures)
	case enums.RightAngleTriangle:
		return shapes_interfaces.NewRightAngleTriangle(measures)
	}
	return  nil
}
