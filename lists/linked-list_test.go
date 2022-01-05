package lists_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kamaal111/go-data-structures/lists"
)

var _ = Describe("Linked List", func() {
	It("inserts 10 nodes", func() {
		list := lists.LinkedList{}
		for i := 0; i < 5; i += 1 {
			list.Push(i)
			Expect(list.Head.Value.(int)).To(Equal(0))
			Expect(list.Tail.Value.(int)).To(Equal(i))
		}

		for i := 5; i < 10; i += 1 {
			list.Unshift(i)
			Expect(list.Head.Value.(int)).To(Equal(i))
			Expect(list.Tail.Value.(int)).To(Equal(4))
		}

		Expect(list.Head.Value.(int)).To(Equal(9))
		Expect(list.Count()).To(Equal(10))
	})

	It("removes the first node in list", func() {
		list := lists.LinkedList{}

		for i := 0; i < 5; i += 1 {
			list.Push(i)
		}

		Expect(list.Count()).To(Equal(5))
		Expect(list.Tail.Value).To(Equal(4))

		node, err := list.Shift()
		Expect(err).To(BeNil())
		Expect(node.Value).To(Equal(0))

		Expect(list.Count()).To(Equal(4))
		Expect(list.Head.Value).To(Equal(1))
	})

	It("removes the last node in list", func() {
		list := lists.LinkedList{}

		for i := 0; i < 5; i += 1 {
			list.Push(i)
		}

		Expect(list.Count()).To(Equal(5))
		Expect(list.Tail.Value).To(Equal(4))

		node, err := list.Pop()
		Expect(err).To(BeNil())
		Expect(node.Value).To(Equal(4))

		Expect(list.Count()).To(Equal(4))
		Expect(list.Tail.Value).To(Equal(3))
	})
})
