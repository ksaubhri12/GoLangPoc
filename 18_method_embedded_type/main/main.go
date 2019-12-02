package main

import (
	"fmt"
	"goLangTraining/18_method_embedded_type/embeddded_type"
	"reflect"
)

func main() {
	carVehicle := embeddded_type.Vehicle{
		Seats:  "4",
		Color:  "Green",
		Engine: "",
	}

	car := embeddded_type.Car{
		Vehicle: &carVehicle,
		Wheels:  "4",
		Doors:   "4",
	}
	carVehicle.SetColor("Red")
	carVehicle.SetEnginePower("456 HP")

	fmt.Println(car)
	fmt.Println(carVehicle)
	fmt.Println(*car.Vehicle , reflect.TypeOf(*car.Vehicle))

}
