package exercise

import "fmt"

func ComputeInput()  {
	fmt.Println("Enter input")
	var input int
	number , _ := fmt.Scanln(&input)
	fmt.Println(100*number)
}
