package main

import "fmt"

type Queue struct {
	data []int
}

// Enqueue adds a value to the queue
func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
}

// Dequeue removes the first element from the queue and returns it
func (q *Queue) Dequeue() int {
	temp := q.data[0]
	q.data = q.data[1:]
	return temp
}

func main() {
	q := Queue{}
	q.Enqueue(10)
	q.Enqueue(12)
	q.Enqueue(13)
	q.Enqueue(15)
	fmt.Println(q.data)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

}
