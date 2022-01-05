package lists

import "fmt"

type LinkedList struct {
	count int
	Head  *listNode
	Tail  *listNode
}

type listNode struct {
	Value    interface{}
	Previous *listNode
	Next     *listNode
}

func (list *LinkedList) Push(value interface{}) {
	node := &listNode{
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
	node := &listNode{
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

func (list *LinkedList) Shift() (*listNode, error) {
	if list.Head == nil {
		return nil, fmt.Errorf("list is empty")
	}

	head := list.Head
	list.Head = head.Next
	head.Next = nil

	list.count -= 1
	return head, nil
}

func (list *LinkedList) Pop() (*listNode, error) {
	if list.Head == nil {
		return nil, fmt.Errorf("list is empty")
	}

	tail := list.Tail
	list.Tail = tail.Previous
	tail.Previous = nil
	list.Tail.Next = nil

	list.count -= 1
	return tail, nil
}

func (list *LinkedList) ToSlice() []interface{} {
	count := list.count
	if count == 0 {
		return []interface{}{}
	}

	array := make([]interface{}, count)

	node := list.Head
	iteration := 0
	for node != nil {
		array[iteration] = node.Value
		node = node.Next
		iteration += 1
	}
	return array
}

func (list *LinkedList) Count() int {
	return list.count
}

func (list *LinkedList) Display() {
	node := list.Head
	fmt.Printf("Head -> ")
	for node != nil {
		fmt.Printf("%+v -> ", node.Value)
		node = node.Next
	}
	fmt.Printf("Tail\n")
}
