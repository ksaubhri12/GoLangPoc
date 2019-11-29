package main

import (
	"fmt"
	"time"
)

func learning()  {
	fmt.Println("Learninng Go routine")
}

func main()  {
	go learning()
	time.Sleep(1*time.Second)
	fmt.Println("Main function")
}
