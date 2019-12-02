package main

import (
	"fmt"
	"goLangTraining/19_interfaces/enums"
	"goLangTraining/19_interfaces/factory"
	"goLangTraining/19_interfaces/shapes_interfaces"
)


func main()  {
	shapes := enums.Square
	measures := shapes_interfaces.Measures{}
	measures.Length = 4.0
	finalShape := factory.ReturnShapeFactory(shapes,measures)
	fmt.Println("Final shape is",finalShape.GetShapeName())
	fmt.Println("Area of shape is", finalShape.Area())
	fmt.Println("Perimeter of shape is",finalShape.Perimeter())
}
