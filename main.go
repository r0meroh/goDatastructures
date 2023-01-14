package main

import (
	"GOPATH/linked_list"
	// "GOPATH/stack_n_queue"
	 "fmt"
)


func main() {
	// fmt.Println("\n===================================================")
	// fmt.Println("Stack Example")
	// fmt.Println("===================================================\n")
	// stack_n_queue.PrintOutStackExample()
	
	// fmt.Println("\n===================================================")
	// fmt.Println("Queue Example")
	// fmt.Println("===================================================\n")

	// stack_n_queue.ExampleQueue()

	myList := linked_list.LinkedList{}

	node1 := &linked_list.Node{}
	node2 := &linked_list.Node{}
	node3 := &linked_list.Node{}

	myList.InsertFront(node1)
	myList.InsertFront(node2)
	myList.InsertFront(node3)
	fmt.Println(myList)

	myList.PrintList()


}