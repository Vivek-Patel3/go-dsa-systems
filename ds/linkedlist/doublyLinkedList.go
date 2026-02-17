package linkedlist

type node[T any] struct {
	val T
	next *node[T]
	prev *node[T]
}

type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func NewNode[T any](val T) *node[T] {
	return &node[T]{
		val: val,
	}
}

func (list *LinkedList[T]) Insert(node *node[T]) {
	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}

	list.tail.next = node
	node.prev = list.tail
	list.tail = node
}

func (list *LinkedList[T]) Remove(node *node[T]) {
	if node == nil {
		return
	}
	
	temp := list.head

	for temp != node {
		temp = temp.next
	}

	temp.prev.next = temp.next
	temp.next.prev = temp.prev

	temp = nil 
}
