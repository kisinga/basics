package main

import (
	"fmt"

	"github.com/kisinga/basics/sorting/O.n.n/bubble"
	"github.com/kisinga/basics/sorting/O.n.n/insertion"
	"github.com/kisinga/basics/sorting/O.n.n/seletion"
	"github.com/kisinga/basics/sorting/O.nlogn/quick"
)

var array = []int{9, 1, 2, 3, 9, 8, 7}

func main() {
	fmt.Println("Input:\n", array)
	// create a copy of the array
	b := array[:]
	bubble.Sort(b)
	fmt.Println(b)

	c := array[:]
	seletion.Sort(c)
	fmt.Println(c)

	d := array[:]
	insertion.Sort(d)
	fmt.Println(d)

	e := array[:]
	quick.Sort(e)
	fmt.Println(d)

}
