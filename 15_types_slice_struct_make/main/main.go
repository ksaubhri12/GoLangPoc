package main

import (
	"fmt"
	"goLangTraining/15_types_slice_struct_make/person_service"
	"reflect"
)

func main()  {

	person := person_service.PostPersonService()
	fmt.Println(person)
	person.SetName("Prashasti")
	fmt.Println(person)

	groundStruct := struct {
		area string;
		colour string
	}{}
	fmt.Println(reflect.TypeOf(groundStruct))
	groundStruct.area = "12345"
	groundStruct.colour = "green"
	fmt.Println(groundStruct)
}
