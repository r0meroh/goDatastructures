package main

import (
	"GOPATH/stack_n_queue"
	"fmt"
)


func main() {

	newStack := stack_n_queue.Stack{}
	fmt.Println(newStack)
	newStack.Push(5)
	newStack.Push(10)
	fmt.Println(newStack)
	newStack.Pop()
	fmt.Println(newStack)

}