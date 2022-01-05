package lists

import "fmt"

type LinkedList struct {
	count int
	Head  *ListNode
	Tail  *ListNode
}

type ListNode struct {
	Value    interface{}
	Previous *ListNode
	Next     *ListNode
}

func (list *LinkedList) Append(value interface{}) {
	node := &ListNode{
		Value:    value,
		Previous: list.Tail,
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
	node := &ListNode{
		Value: value,
	}

	if list.Head != nil {
		list.Head.Previous = node
		node.Next = list.Head
		list.Head = node
	} else {
		list.Head = node
	}

	list.count += 1
}

func (list *LinkedList) Shift() (*ListNode, error) {
	if list.Head == nil {
		return nil, fmt.Errorf("list is empty")
	}

	head := list.Head
	list.Head = head.Next
	head.Next = nil

	list.count -= 1
	return head, nil
}

func (list *LinkedList) Count() int {
	return list.count
}

func (l *LinkedList) Display() {
	list := l.Head
	fmt.Printf("Head -> ")
	for list != nil {
		fmt.Printf("%+v -> ", list.Value)
		list = list.Next
	}
	fmt.Printf("Tail\n")
}
