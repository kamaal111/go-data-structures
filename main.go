package main

import (
	"fmt"
	"log"

	"github.com/kamaal111/go-data-structures/lists"
)

func main() {
	list := lists.LinkedList{}
	list.Append(22)
	list.Append(44)
	list.Append(11)
	list.Display()
	removedNode, err := list.Shift()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(removedNode)
	list.Display()
}
