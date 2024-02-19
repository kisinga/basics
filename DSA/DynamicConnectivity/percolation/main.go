package percolation

import (
	"math/rand"

	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/models"
)

func New(n int) *models.Data {

	data := models.Data{}

	// connect the first row
	for i := range n {
		data.Connect(models.Node{Row: 0, Col: i}, models.VirtualRootTop)
	}

	// connect the last row to the virtual bottom root
	for i := range n {
		data.Connect(models.Node{Row: n - 1, Col: i}, models.VirtualRootBottom)
	}

	// initialize the rest of the grid
	for i := 1; i < n; i++ {
		for j := 0; j <= n; j++ {
			data.AddNode(i, j)
		}
	}

	return &data

}

func main() {

	gridSize := 10
	var data models.Percolation = New(gridSize)
	data.Open(1, 1)

	count := 0
	// open random sites until the system percolates
	for !data.Percolates() {
		// create a random row and column restricted to the size of the grid
		data.Open(rand.Intn(gridSize), rand.Intn(gridSize))
		count++
	}

	// print the number of open sites

}
