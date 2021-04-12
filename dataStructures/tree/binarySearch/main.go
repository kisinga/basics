package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (t *Node) Insert(key int) {
	if t.Key < key {
		// move right
		if t.Right == nil {
			t.Right = &Node{
				Key: key,
			}
		} else {
			t.Right.Insert(key)
		}
	} else {
		// move left
		if t.Left == nil {
			t.Left = &Node{
				Key: key,
			}
		} else {
			t.Left.Insert(key)
		}
	}
}

func (t *Node) Search(key int) bool {
	if t == nil {
		return false
	}
	if t.Key < key {
		// move right
		return t.Right.Search(key)
	} else if t.Key > key {
		// move left
		return t.Left.Search(key)
	}
	return true
}

func main() {
	t := &Node{Key: 100}
	t.Insert(100)
	t.Insert(200)
	t.Insert(167)
	t.Insert(2030)
	t.Insert(340)
	t.Insert(43)
	t.Insert(420)
	t.Insert(42)
	t.Insert(300)
	fmt.Println(t.Search(300))
	fmt.Println(t.Search(900))

}
