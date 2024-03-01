package models

type Node struct {
	Row    int
	Col    int
	IsOpen bool
}

type Grid [][]Node

func EmptyData(size int) Grid {

	// we add 2 rows, one for vTop and the other for vBottom
	grid := make([][]Node, size+2)

	for i := range size {
		grid[i] = make([]Node, size)
		for j := range size {
			grid[i][j] = Node{
				Row:    i,
				Col:    j,
				IsOpen: false,
			}
		}
	}
	// create the vtop and vbottom and connect each to itself
	grid[0][0] = Node{
		Row:    0,
		Col:    0,
		IsOpen: false,
	}
	grid[size+1][0] = Node{
		Row:    0,
		Col:    0,
		IsOpen: false,
	}

	// connect the first row to the vTop and the last row to the VBottom
	for i := 0; i < size; i++ {
		grid[0][i] = grid[0][0]
		grid[size-1][i] = grid[size+1][0]
	}

	return grid
}

type Percolation interface {

	// opens the site (row, col) if it is not open already
	Open(Row, Col int)

	// is the site (row, col) open?
	IsOpen(Row, Col int) bool

	// is the site (row, col) full?
	IsFull(Row, Col int) bool

	// returns the number of open sites
	NumberOfOpenSites() int

	// does the system percolate?
	Percolates() bool

	Connect(RowA, ColA, RowB, ColB int)
}

type PercolationStats interface {
	// perform trials independent experiments on an n-by-n grid
	PercolationStats(n, trials int) float64

	// sample mean of percolation threshold
	Mean() float64

	// sample standard deviation of percolation threshold
	Stddev() float64

	// low endpoint of 95% confidence interval
	ConfidenceLo() float64

	// high endpoint of 95% confidence interval
	ConfidenceHi() float64
}
