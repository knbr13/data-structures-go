package doublylinkedlist

type DoublyLinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length uint
}

func New[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (l *DoublyLinkedList[T]) Len() uint { return l.length }

func (l *DoublyLinkedList[T]) PushBack(v T) {
	n := NewNode(v)
	if l.length == 0 {
		l.head, l.tail = n, n
	} else {
		n.Prev = l.tail
		l.tail.Next = n
		l.tail = n
	}
	l.length++
}

func (l *DoublyLinkedList[T]) PushFront(v T) {
	n := NewNode(v)
	if l.length == 0 {
		l.head, l.tail = n, n
	} else {
		n.Next = l.head
		l.head.Prev = n
		l.head = n
	}
	l.length++
}

func (l *DoublyLinkedList[T]) PopFront() (T, bool) {
	if l.length == 0 {
		var v T
		return v, false
	}
	v := l.head.Value
	l.head = l.head.Next
	if l.head != nil {
		l.head.Prev = nil
	} else {
		l.tail = nil
	}
	l.length--
	return v, true
}

func (l *DoublyLinkedList[T]) PopBack() (T, bool) {
	if l.length == 0 {
		var v T
		return v, false
	}
	v := l.tail.Value
	l.tail = l.tail.Prev
	if l.tail != nil {
		l.tail.Next = nil
	} else {
		l.head = nil
	}
	l.length--
	return v, true
}
