package util

import (
	"GOPATH/stack_n_queue"
	"fmt"
)

func Run_stack() {
	fmt.Println("\n===================================================")
	fmt.Println("Stack Example")
	fmt.Println("===================================================\n")
	fmt.Println("Creating an empty stack, pushing to it and popping elements")
	newStack := stack_n_queue.Stack{}
	fmt.Println(newStack)
	newStack.Push(5)
	newStack.Push(10)
	fmt.Println("elements pushed")
	fmt.Println(newStack)
	newStack.Pop()
	fmt.Println("popped off an element")
	fmt.Println(newStack)
	newStack.Push(6)
	fmt.Println(newStack)
	popped := newStack.ReturnPop()
	fmt.Println(popped, " was popped off stack. Stack is currently ", newStack)
}
