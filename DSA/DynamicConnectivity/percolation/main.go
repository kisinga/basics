package main

import (
	"math/rand"
	"sync"

	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/models"
	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/quickfind"
	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/quickunion"
)

func New(n int) (qu models.Percolation, qf models.Percolation) {

	qf = quickfind.New(n)
	qu = quickunion.New(n)

	return qu, qf

}

func calculateMany(count int) float64 {
	probabilityResults := make(chan float64, count)
	var wg sync.WaitGroup
	for range count {
		wg.Add(1)
		go func() {
			defer wg.Done()
			gridSize := rand.Intn(96) + 3
			_, qf := New(gridSize)
			for !qf.Percolates() {
				// create a random row and column restricted to the size of the grid
				qf.Open(rand.Intn(gridSize), rand.Intn(gridSize))
			}
			probability := 0.0
			openSites := qf.NumberOfOpenSites()
			probability = float64(openSites) / float64(gridSize*gridSize)
			probabilityResults <- probability
		}()

	}
	go func() {
		wg.Wait()
		close(probabilityResults)
	}()

	totals := 0.0
	for val := range probabilityResults {
		totals += val
	}
	return totals / float64(count)
}

func main() {

	sampleSize := 50

	// open random sites until the system percolates
	// for !qu.Percolates() {
	// 	// create a random row and column restricted to the size of the grid
	// 	qu.Open(models.Node{Row: rand.Intn(gridSize), Col: rand.Intn(gridSize)})
	// 	count++
	// }

	// for !qf.Percolates() {
	// 	// create a random row and column restricted to the size of the grid
	// 	qf.Open(rand.Intn(gridSize), rand.Intn(gridSize))
	// }
	// probability := 0.0
	// openSites := qf.NumberOfOpenSites()
	// probability = float64(openSites) / float64(gridSize)
	probability := calculateMany(sampleSize)
	println(probability)
	// print the number of open sites

}
