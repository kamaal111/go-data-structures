package main

import (
	"fmt"

	"github.com/kamaal111/go-data-structures/lists"
)

func main() {
	list := lists.LinkedList{}
	list.Append(22)
	list.Append(44)
	list.Display()
	fmt.Println(list.Head.Value)
	fmt.Println(list.Tail.Value)
	fmt.Println(list.Count())
}
