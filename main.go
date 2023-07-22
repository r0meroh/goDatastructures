package main

import (
	"GOPATH/myStructs"
	"fmt"
)

func main() {

	// stack_n_queue.PrintOutStackExample()
	// stack_n_queue.ExampleQueue()
	// linked_list.LinkedListExample()
	Hugo := myStructs.New_passenger("Hugo", "xbn1234", false)
	fmt.Println(myStructs.Print_name(Hugo))
	fmt.Println(myStructs.Print_ticket(Hugo))
	fmt.Println(myStructs.Is_boarded(Hugo))

}
