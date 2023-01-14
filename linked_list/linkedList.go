package linked_list

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Length int
}

func (l LinkedList) PrintList() {
	currentNode := l.Head
	for l.Length != 0 {
		fmt.Printf("%d", currentNode)
		currentNode = currentNode.Next
		l.Length = l.Length - 1
	}
	fmt.Println("-----------")
}

func (l *LinkedList) InsertFront(n *Node) {
	oldHead := l.Head
	l.Head = n
	l.Head = oldHead
	l.Length = l.Length + 1
}