package linkedlist

type Node struct {
	val  int
	next *Node
}

type SinglyLinkedList struct {
	Head *Node
	Tail *Node
}

func (s *SinglyLinkedList) InsertAtHead(val int) {
	newNode := &Node{val: val, next: nil}
	if s.Head == nil {
		s.Head = newNode
	} else {
		newNode.next = s.Head
		s.Head = newNode
	}
}
