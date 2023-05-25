package linkedlist

import (
	"fmt"
)

/*
usage:
	list := NewList()
	list.AddNode(NewNode(value))
	list.Erase(nd)
*/

type Node struct {
	val  int
	next *Node
	prev *Node
}

func NewNode(val int) *Node {
	return &Node{val: val, next: nil, prev: nil}
}

type List struct {
	head *Node
	tail *Node
}

func NewList() *List {
	return &List{head: nil, tail: nil}
}

func (l *List) Empty() bool {
	return l.head == nil && l.tail == nil
}

func (l *List) AddNode(u *Node) {
	if l.Empty() {
		l.head = u
		l.tail = u
	} else {
		l.tail.next = u
		l.tail.next.prev = l.tail
		l.tail = l.tail.next
	}
}

func (l *List) Erase(u *Node) {
	if u == nil {
		return
	}

	if u.next != nil {
		u.next.prev = u.prev
	} else {
		l.tail = u.prev
	}
	if u.prev != nil {
		u.prev.next = u.next
	} else {
		l.head = u.next
	}
}

func PrintList(l *List) {
	p := l.head
	for p != nil {
		fmt.Printf("%d ", p.val)
		p = p.next
	}
	fmt.Printf("\n")
}
