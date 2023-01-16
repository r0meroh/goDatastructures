package main

import (
	"GOPATH/linked_list"
	// "GOPATH/stack_n_queue"
	
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

	node1 := &linked_list.Node{Data: 1}
	node2 := &linked_list.Node{Data: 2}
	node3 := &linked_list.Node{Data: 3}

	myList.InsertFront(node1)
	myList.InsertFront(node2)
	myList.InsertFront(node3)
	myList.InsertNodeWithValue(55)
	myList.InsertNodeWithValue(66)
	myList.InsertNodeWithValue(77)
	myList.PrintList()



}