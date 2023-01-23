package linked_list

type DoublyNode struct {
	prev *DoublyNode
	next *DoublyNode 
	data int
}

type DoublyLinkedList struct {
	head *DoublyNode
	tail *DoublyNode
	length int
}

// func (dl *DoublyLinkedList) InsertInFront (dn *DoublyNode) {
// 	oldHead := dl.head
// 	oldTail := dl.tail
// 	dl.head = dn.next
// 	if dl.length > 0{
// 		dl.tail = ol 
// 		dl.tail
// 	}
// }
