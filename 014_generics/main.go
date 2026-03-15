package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
// This is a generic function that works for any type T that is comparable.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the 'comparable' constraint,
		// so we can use ==.
		if v == x {
			return i
		}
	}
	return -1
}

// List represents a singly-linked list that can hold values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Push(v T) {
	if l.next == nil {
		l.next = &List[T]{val: v}
	} else {
		l.next.Push(v)
	}
}

func (l *List[T]) Print() {
	if l == nil {
		return
	}
	fmt.Printf("%v ", l.val)
	if l.next != nil {
		l.next.Print()
	}
}

func main() {
	// 1. Generic Functions
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println("Index of 15 in si:", Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println("Index of \"hello\" in ss:", Index(ss, "hello"))

	// 2. Generic Types
	// Creating a linked list of ints
	fmt.Print("\nLinked List of Ints: ")
	head := &List[int]{val: 1}
	head.Push(2)
	head.Push(3)
	head.Print()
	fmt.Println()

	// Creating a linked list of strings
	fmt.Print("Linked List of Strings: ")
	sHead := &List[string]{val: "Hello"}
	sHead.Push("Generics")
	sHead.Push("World")
	sHead.Print()
	fmt.Println()
}
