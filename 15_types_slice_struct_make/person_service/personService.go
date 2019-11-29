package person_service

import "goLangTraining/15_types_slice_struct_make/person_struct"

func PostPersonService() *person_struct.Person {
	person := CreatePerson("kalpesh","Saubhri","25","180")
	return person
}

func CreatePerson(name string, lastName string,age string, height string)  *person_struct.Person{
	person := person_struct.NewPerson(name,lastName,age,height)
	return person
}
