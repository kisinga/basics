package main

import (
	"math/rand"

	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/models"
	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/quickfind"
	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/quickunion"
)

func New(n int) (qu models.Percolation, qf models.Percolation) {

	qf = quickfind.New(n)
	qu = quickunion.New(n)

	return qu, qf

}

func main() {

	gridSize := 5
	_, qf := New(gridSize)

	count := 0
	// open random sites until the system percolates
	// for !qu.Percolates() {
	// 	// create a random row and column restricted to the size of the grid
	// 	qu.Open(models.Node{Row: rand.Intn(gridSize), Col: rand.Intn(gridSize)})
	// 	count++
	// }

	for !qf.Percolates() {
		// create a random row and column restricted to the size of the grid
		qf.Open(rand.Intn(gridSize), rand.Intn(gridSize))
		count++
	}

	println(count)

	// print the number of open sites

}
