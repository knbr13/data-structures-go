package queue

import (
	linkedlist "github.com/abdullah-alaadine/data-structures-go/linked_list"
)

type Queue struct {
	list linkedlist.Linkedlist
}

func NewQueue() *Queue {
	return &Queue{
		list: *linkedlist.NewLinkedList(),
	}
}