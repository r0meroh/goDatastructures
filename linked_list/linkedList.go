package linked_list

import (
	"errors"
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

func(l *LinkedList) InsertNodeWithValue (value int){
	oldHead := l.Head
	tempNode := &Node{Data: value}
	l.Head = tempNode
	l.Head.Next = oldHead
	l.Length ++
}

func (l *LinkedList) InsertFront(n *Node) {
	oldHead := l.Head
	l.Head = n
	l.Head.Next = oldHead
	l.Length ++
}

func (l *LinkedList) DeleteByValue(value int){
	l.CheckIfEmpty()
	if l.Head.Data == value{
		l.Head = l.Head.Next
		l.Length --
	}
	previousNodeToDelete := l.Head
	for previousNodeToDelete.Next.Data != value{
		previousNodeToDelete = previousNodeToDelete.Next
	}
	previousNodeToDelete.Next = previousNodeToDelete.Next.Next
	l.Length --
}

func (l *LinkedList) CheckIfEmpty()error{
	if l.Length == 0{
		
		return  errors.New("Found empty list")
	}
	return nil
}

func LinkedListExample (){
	
	myList := LinkedList{}

	node1 := &Node{Data: 1}
	node2 := &Node{Data: 2}
	node3 := &Node{Data: 3}

	myList.InsertFront(node1)
	myList.InsertFront(node2)
	myList.InsertFront(node3)
	myList.InsertNodeWithValue(55)
	myList.InsertNodeWithValue(66)
	myList.InsertNodeWithValue(77)
	myList.PrintList()
	fmt.Println("deleting 66" )
	myList.DeleteByValue(66)
	myList.PrintList()
}