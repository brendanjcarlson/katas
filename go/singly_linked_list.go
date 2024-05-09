package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (list *LinkedList) Display() {
	if list.head == nil {
		fmt.Println("list is empty")
		return
	}

	fmt.Printf("DISPLAY: \n\t%d", list.head.data)

	curr := list.head.next
	for curr != nil {
		fmt.Printf(" -> %d", curr.data)
		curr = curr.next
	}
	fmt.Printf("\n\tlength: %d\n", list.length)
	fmt.Println()
}

func (list *LinkedList) Append(data int) {
	node := &Node{data: data}

	if list.head == nil {
		list.head = node
	} else {
		curr := list.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = node
	}

	list.length++
}

func (list *LinkedList) Prepend(data int) {
	node := &Node{data: data}

	if list.head == nil {
		list.head = node
	} else {
		node.next = list.head
		list.head = node
	}

	list.length++
}

func (list *LinkedList) InsertAfter(target, data int) {
	node := &Node{data: data}

	curr := list.head
	for curr != nil {
		if curr.data == target {
			node.next = curr.next
			curr.next = node
			list.length++
			return
		}
		curr = curr.next
	}

	fmt.Printf("target not found\n\n")
}

func (list *LinkedList) InsertBefore(target, data int) {
	if list.head == nil {
		fmt.Println("list is empty")
	}

	if list.head.data == target {
		list.Prepend(data)
		return
	}

	node := &Node{data: data}

	curr := list.head
	for curr.next != nil {
		if curr.next.data == target {
			node.next = curr.next
			curr.next = node
			list.length++
			return
		}
		curr = curr.next
	}

	fmt.Printf("target not found\n\n")
}

func (list *LinkedList) RemoveFront() {
	if list.head == nil {
		fmt.Println("list is empty")
		return
	}
	list.head = list.head.next
	list.length--
}

func (list *LinkedList) RemoveEnd() {
	if list.head == nil {
		fmt.Println("list is empty")
		return
	}

	if list.head.next == nil {
		list.head = nil
	} else {

		curr := list.head
		for curr.next.next != nil {
			curr = curr.next
		}

		curr.next = nil
	}
	list.length--
}

func (list *LinkedList) ValueAt(position int) {
	if position > list.length-1 || position < 0 {
		fmt.Printf("VALUEAT: \n\tposition out of bounds\n\n")
		return
	}

	curr := list.head
	for i := 0; i < position; i++ {
		curr = curr.next
	}
	fmt.Printf("VALUEAT: \n\tposition %d is %d\n\n", position, curr.data)
}

func (list *LinkedList) Walk() {
	curr := list.head
	for curr != nil {
		fmt.Printf("%d\n", curr.data)
		curr = curr.next
	}
}

func main() {
	fmt.Printf("Linked List Kata\n\n")

	list := new(LinkedList)

	list.Append(4)
	list.Prepend(3)
	list.InsertBefore(3, 1)
	list.InsertAfter(1, 2)
	list.Prepend(0)

	list.Display()

	list.ValueAt(3)
}
