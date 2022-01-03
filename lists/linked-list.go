package lists

import "fmt"

type LinkedList struct {
	count int
	Head  *listNode
	Tail  *listNode
}

type listNode struct {
	Value    interface{}
	Next     *listNode
	Previous *listNode
}

func (list *LinkedList) Append(value interface{}) {
	node := &listNode{
		Value: value,
	}
	if list.Tail != nil {
		list.Tail.Next = list.Tail
		list.Tail.Next = node
	} else {
		list.Head = node
	}
	list.Tail = node
	list.count += 1
}

func (list *LinkedList) Unshift(value interface{}) {
	node := &listNode{
		Next:  list.Head,
		Value: value,
	}
	if list.Head != nil {
		list.Head.Previous = node
	}
	list.Head = node

	head := list.Head
	for head.Next != nil {
		head = head.Next
	}
	list.Tail = head
	list.count += 1
}

func (list *LinkedList) Count() int {
	return list.count
}

func (l *LinkedList) Display() {
	list := l.Head
	for list != nil {
		fmt.Printf("%+v -> ", list.Value)
		list = list.Next
	}
	fmt.Printf("Head\n")
}
