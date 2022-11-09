package main

import "fmt"

type node struct {
	data     int
	next     *node
	previous *node
}

type LinkedList struct {
	Head   *node
	Tail   *node
	Length int
}

func (l *LinkedList) prepend(n *node) {
	second := l.Head
	l.Head = n
	l.Head.next = second
	l.Length++
}

func (l *LinkedList) DeleteWithValue(val int) {
	if l.Length == 0 {
		return
	}
	if val == l.Head.data {
		l.Head = l.Head.next
		l.Length--
		return
	}
	previousValue := l.Head
	for previousValue.next.data != val {
		if previousValue.next.next == nil {
			return
		}
		previousValue = previousValue.next
	}
	previousValue.next = previousValue.next.next
	l.Length--
}

func (l LinkedList) Data() []int {
	data := []int{}
	current := l.Head
	for i := 0; i < l.Length; i++ {
		data = append(data, current.data)
		current = current.next
	}
	return data
}

func main() {
	myList := LinkedList{}
	nodes := []*node{{data: 12}, {data: 15}, {data: 90}, {data: 4}, {data: 23}, {data: 42}}
	for _, val := range nodes {
		myList.prepend(val)
	}
	fmt.Println(myList.Data())

	myList.DeleteWithValue(0)

	fmt.Println(myList.Data())

}
