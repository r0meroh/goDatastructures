package stack_n_queue

// stack

type Stack struct {
	element []int
}

// push
func (myStack *Stack) Push(i int) {
	myStack.element = append(myStack.element, i)
}

// pop

func (myStack *Stack) Pop() {
	myStack.element = myStack.element[:len(myStack.element)-1]
}

// Pop with removed item being returned

func (myStack *Stack) ReturnPop() int {
	temp := myStack.element[len(myStack.element)-1]
	myStack.Pop()
	return temp
}
