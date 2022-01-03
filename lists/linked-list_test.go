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
			list.Append(i)
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
			list.Append(i)
		}

		Expect(list.Count()).To(Equal(5))
		Expect(list.Tail.Value.(int)).To(Equal(4))

		list.Shift()

		Expect(list.Count()).To(Equal(4))
		Expect(list.Head.Value.(int)).To(Equal(1))
	})
})
