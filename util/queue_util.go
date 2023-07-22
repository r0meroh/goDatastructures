package util

import (
	"GOPATH/stack_n_queue"
	"fmt"
)

func Run_queue() {
	// stack_n_queue.PrintOutStackExample()
	fmt.Println("\n===================================================")
	fmt.Println("Queue Example")
	fmt.Println("===================================================\n")

	fmt.Println("Creating an empty Queue. Adding to it and removing values")
	newQueue := stack_n_queue.Queue{}
	fmt.Println(newQueue)
	newQueue.Enqueue(1)
	newQueue.Enqueue(2)
	newQueue.Enqueue(3)
	fmt.Printf("Added to empty queue. Current queue is: %v. Removing values\n", newQueue)
	newQueue.Dequeue()
	fmt.Println(newQueue)
	fmt.Printf("Next value to be removed is: %v. \nFinal values in queue: ", newQueue.ReturnDequeueValue())
	fmt.Println(newQueue)
}
