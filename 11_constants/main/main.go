package main

import "fmt"

func main() {
	//a := 10
	//
	//fmt.Println(constName.B, constName.G)
	//fmt.Println(constName.KB,constName.MB)
	//fmt.Println(constName.AGE)
	//fmt.Println("a memory address is: ",&a)
	//
	//var meter float64
	//fmt.Println("Input Number")
	//fmt.Scan(&meter)
	//yards := meter*1.6
	//fmt.Println("Yards is:",yards)
	x := 5
	two(&x)
	fmt.Println(x)


}

func two(x *int)  {
	*x = 2
}
