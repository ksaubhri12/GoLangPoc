package main

import (
	"fmt"
	"goLangTraining/13_variadic/exercise"
	"strconv"
)

func main()  {
	//s := []string{"Moulik","Aditya"}
	Greeting("Hello", "Kalpesh","Moulik")
}

func Greeting(prefix string, who ...string)  {
	exercise.ComputeInput()
}

func printNumbers(input ... interface{})  {
	for _, number := range input{
		switch number.(type) {
		case int:
			fmt.Println(number)
			continue
		case string:
			val , _ := strconv.Atoi(number.(string))
			fmt.Println(val)
			continue
		default:
			fmt.Println(number)
			continue
		}
	}

}
