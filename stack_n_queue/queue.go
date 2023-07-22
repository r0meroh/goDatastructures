package stack_n_queue

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
