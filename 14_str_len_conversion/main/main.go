package main

import (
	"fmt"
	"reflect"
)

func main()  {
	rune := 'd'
	x := 'd'
	fmt.Println(reflect.TypeOf(rune),rune,x)
}
