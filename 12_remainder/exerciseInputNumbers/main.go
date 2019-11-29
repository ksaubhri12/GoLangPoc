package main

import "fmt"

func main()  {
	var firstNumber int
	var secondNumber int

	fmt.Println("Enter first number")
	fmt.Scan(&firstNumber)
	fmt.Println("Enter Second number")
	fmt.Scan(&secondNumber)

	fmt.Println(firstNumber%secondNumber)
}
