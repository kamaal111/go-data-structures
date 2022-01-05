package main

import (
	"fmt"
	"log"

	"github.com/kamaal111/go-data-structures/lists"
)

func main() {
	list := lists.LinkedList{}
	list.Push(22)
	list.Push(44)
	list.Push(11)
	list.Display()
	fmt.Println(list.Count())
	removedNode, err := list.Pop()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(removedNode)
	list.Display()
	fmt.Println(list.Count())
}
