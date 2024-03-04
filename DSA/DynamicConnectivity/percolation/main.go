package main

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/models"
	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/quickfind"
	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/quickunion"
)

func New(n int) (models.Percolation, models.Percolation, error) {

	qf, err := quickfind.New(n)
	if err != nil {
		return nil, nil, err
	}
	qu, err := quickunion.New(n)
	if err != nil {
		return nil, nil, err
	}
	return qu, qf, nil

}

func calculateMany(count int) float64 {
	probabilityResults := make(chan float64, count)
	var wg sync.WaitGroup
	for range count {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// make grid size between 5-100
			gridSize := rand.Intn(96) + 5
			_, qf, err := New(gridSize)
			if err != nil {
				return
			}
			for !qf.Percolates() {
				// create a random row and column restricted to the size of the grid
				// the first and last rows are vtop and vbottom
				_ = qf.Open(rand.Intn(gridSize)+1, rand.Intn(gridSize))
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
	resultCount := 0
	// in the case scenario that there are errors in the goroutines, dont rely on the buffer size as the source of truth
	for val := range probabilityResults {
		totals += val
		resultCount++
	}
	if resultCount != count {
		fmt.Println("there was an error creating some grids")
	}
	return totals / float64(resultCount)
}

func main() {

	sampleSize := 200

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
