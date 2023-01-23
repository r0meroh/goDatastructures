package stack_n_queue

import "fmt"

// queue
type Queue struct {
	elements []int
}

// enqueue
func (myQueue *Queue) Enqueue(i int) {
	myQueue.elements = append(myQueue.elements, i)
}

// dequeue
func (myQueue *Queue) Dequeue() {
	myQueue.elements = myQueue.elements[1:]
}

// return dequeue value
func (myQueue *Queue) ReturnDequeueValue() int {
	temp := myQueue.elements[0]
	myQueue.Dequeue()
	return temp
}

// print example usage
func ExampleQueue() {
	// stack_n_queue.PrintOutStackExample()
	fmt.Println("\n===================================================")
	fmt.Println("Queue Example")
	fmt.Println("===================================================\n")

	fmt.Println("Creating an empty Queue. Adding to it and removing values")
	newQueue := Queue{}
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