package main

import (
	"fmt"
	"goLangTraining/15_types_slice_struct_make/map"
	"goLangTraining/goroutines_medium/mining_example/mining_service"
	"time"
)

func main() {
	mining_service.Execute()

	doneChan := make(chan string)
	go func() {
		fmt.Println("Run for 5 seconds")
		time.Sleep(time.Second*5)
		doneChan <- "“I’m all done!”"
	}()
	nameMap := _map.MakeMap()
	fmt.Println("Map before Update",nameMap)
	_map.UpdateMap(nameMap)
	fmt.Println("Map after Update",nameMap)
	delete(nameMap,"Kalpesh")
	fmt.Println("After deleting one entry in map ",nameMap)
	<-doneChan
}
