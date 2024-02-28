package percolation

import (
	"math/rand"

	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/models"
	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/quickfind"
	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/quickunion"
)

func New(n int) (qu models.Percolation, qf models.Percolation) {

	qf = quickfind.New(n)
	qu = quickunion.New(n)

	// connect the first row
	for i := range n {
		qu.Connect(models.Node{Row: 0, Col: i}, models.VirtualRootTop)
		qu.Connect(models.Node{Row: 0, Col: i}, models.VirtualRootTop)
	}

	// connect the last row to the virtual bottom root
	for i := range n {
		qu.Connect(models.Node{Row: n - 1, Col: i}, models.VirtualRootBottom)
		qf.Connect(models.Node{Row: n - 1, Col: i}, models.VirtualRootBottom)
	}

	return qu, qf

}

func main() {

	gridSize := 10
	qu, qf := New(gridSize)

	qu.Open(models.Node{Row: 0, Col: 0})
	qf.Open(models.Node{Row: 0, Col: 0})

	count := 0
	// open random sites until the system percolates
	for !qu.Percolates() {
		// create a random row and column restricted to the size of the grid
		qu.Open(models.Node{Row: rand.Intn(gridSize), Col: rand.Intn(gridSize)})
		count++
	}

	for !qf.Percolates() {
		// create a random row and column restricted to the size of the grid
		qf.Open(models.Node{Row: rand.Intn(gridSize), Col: rand.Intn(gridSize)})
		count++
	}

	// print the number of open sites

}
