package linkedlist

type Node[T comparable] struct {
	Val T
	next *Node[T]
	prev *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func NewNode[T comparable](val T) *Node[T] {
	return &Node[T]{
		Val: val,
	}
}

func (list *LinkedList[T]) Insert(node *Node[T]) {
	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}

	list.tail.next = node
	node.prev = list.tail
	list.tail = node
}

func (list *LinkedList[T]) Search(val T) *Node[T] {
	temp := list.head

	for temp != nil && temp.Val != val {
		temp = temp.next
	}

	return temp
}

func (list *LinkedList[T]) Remove(val T) {
	temp := list.Search(val)

	temp.prev.next = temp.next
	temp.next.prev = temp.prev

	temp = nil 
}
