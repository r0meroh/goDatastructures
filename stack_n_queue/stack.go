package stack_n_queue

import "fmt"
// stack

type Stack struct {
	element []int
}

// push
func (myStack *Stack) Push (i int){
	myStack.element = append(myStack.element, i)
}

// pop

func (myStack *Stack) Pop () {
	myStack.element = myStack.element[:len(myStack.element)-1]
}

// Pop with removed item being returned

func (myStack *Stack) ReturnPop () int {
	temp := myStack.element[len(myStack.element)-1]
	myStack.Pop()
	return temp
}

func PrintOutStackExample (){
	fmt.Println("\n===================================================")
	fmt.Println("Stack Example")
	fmt.Println("===================================================\n")
	fmt.Println("Creating an empty stack, pushing to it and popping elements")
	newStack := Stack{}
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

