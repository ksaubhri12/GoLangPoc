package main

import ("fmt"
		"goLangTraining/10_variables/names"
)

var value int

func main()  {
	message := "Hello World!"
	message = message + "kalpesj"
	a,b,c   := 1,false,3
	g:= 'h'
	k := "O"
	var myName string
	myName = "Kalpesh"

	fmt.Println(message,a,b,c,g,k, myName)
	fmt.Println(names.MyName)
}
