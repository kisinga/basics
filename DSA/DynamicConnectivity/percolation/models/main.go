package models

type Node struct {
	Row    int
	Col    int
	IsOpen bool
}

type Data struct {
	Grid
	RealSize   int
	VTopPos    int
	VBottomPos int
}

type Grid [][]Node

func EmptyData(size int) *Data {
	model := &Data{
		RealSize:   size,
		VTopPos:    size,
		VBottomPos: size + 1,
	}
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

	grid[model.VTopPos] = make([]Node, 1)
	grid[model.VBottomPos] = make([]Node, 1)
	// create the vtop and vbottom and connect each to itself
	grid[model.VTopPos][0] = Node{
		Row:    model.VTopPos,
		Col:    0,
		IsOpen: false,
	}
	grid[model.VBottomPos][0] = Node{
		Row:    model.VBottomPos,
		Col:    0,
		IsOpen: false,
	}

	// connect the first row to the vTop and the last row to the VBottom
	for i := 0; i < size; i++ {
		grid[0][i] = grid[model.VTopPos][0]
		grid[size-1][i] = grid[model.VBottomPos][0]
	}

	model.Grid = grid

	return model
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

	Connect(Row1, Col1, Row, Col int)
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
