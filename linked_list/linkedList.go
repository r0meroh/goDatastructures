package linked_list

import (
	"fmt"
	"strings"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Length int
}

func (l LinkedList) PrintList() {
	flowerBox := strings.Repeat("=", l.Length * 3)
	currentNode := l.Head
	fmt.Println("Printing linked list:")
	fmt.Println(flowerBox)
	
	for l.Length != 0 {
		if l.Length > 1{
		fmt.Printf("%d->", currentNode.Data)
		currentNode = currentNode.Next
		l.Length --
		}else{
			fmt.Printf("%d\n", currentNode.Data)
			l.Length --
		}
	}
	fmt.Println(flowerBox)
}

func (l *LinkedList) InsertFront(n *Node) {
	oldHead := l.Head
	l.Head = n
	l.Head.Next = oldHead
	l.Length ++
}