package main

import "fmt"

const ArraySize = 7

type HashTable struct {
	values [ArraySize]*bucket
}

func (h *HashTable) Insert(k string) {
	index := hash(k)
	h.values[index].insert(k)
}

func (h *HashTable) Search(k string) bool {
	index := hash(k)
	return h.values[index].search(k)
}

func (h *HashTable) Delete(k string) bool {
	index := hash(k)
	return h.values[index].delete(k)
}

func hash(k string) int {
	sum := 0
	for _, val := range k {
		sum += int(val)
	}

	return sum % ArraySize
}

type node struct {
	key  string
	next *node
}

type bucket struct {
	head *node
}

func (b *bucket) insert(k string) {
	if !b.search(k) {
		current := b.head
		n := &node{
			key:  k,
			next: current,
		}
		b.head = n
	} else {
		fmt.Println("already exists")
	}

}

func (b *bucket) search(k string) bool {
	current := b.head
	// ------------------
	// This can be optimised
	// if current == nil {
	// 	return false
	// }
	// if current.key == k {
	// 	return true
	// }
	// for current.key != k {
	// 	if current.next == nil {
	// 		return false
	// 	}
	// 	current = current.next
	// }
	// -------------------
	// This can be optimised

	for current != nil {
		if current.key == k {
			return true
		}
		current = current.next
	}
	return false
}

func (b *bucket) delete(k string) bool {
	current := b.head
	if current.key == k {
		b.head = b.head.next
		return true
	}
	for current != nil {
		if current.next.key == k {
			current.next = current.next.next
			return false
		}
		current = current.next
	}
	return false
}

func Init() *HashTable {
	t := &HashTable{}
	for i := range t.values {
		t.values[i] = &bucket{}
	}
	return t
}

func main() {
	t := Init()
	t.Insert("RANDY")
	t.Insert("ANDY")
	t.Insert("MEE")
	t.Insert("MEE")
	t.Delete("ANDY")
	fmt.Println(t)
	fmt.Println(t.Search("ANDY"))
	fmt.Println(t.Search("RANDY"))
	fmt.Println(t.Search("MEE"))
}
