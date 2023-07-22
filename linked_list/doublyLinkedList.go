package linked_list

type DoublyNode struct {
	prev *DoublyNode
	next *DoublyNode
	data int
}

type DoublyLinkedList struct {
	head   *DoublyNode
	tail   *DoublyNode
	length int
}
