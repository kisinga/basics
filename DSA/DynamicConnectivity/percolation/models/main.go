package models

type Node struct {
	Row int
	Col int
}

var VirtualRootTop = Node{
	Row: -1,
	Col: -1,
}

var VirtualRootBottom = Node{
	Row: -2,
	Col: -2,
}

type Percolation interface {

	// opens the site (row, col) if it is not open already
	Open(node Node)

	// is the site (row, col) open?
	IsOpen(node Node) bool

	// is the site (row, col) full?
	IsFull(node Node) bool

	// returns the number of open sites
	NumberOfOpenSites() int

	// does the system percolate?
	Percolates() bool

	AddNode(nd Node)

	Connect(node Node, rootNode Node)
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
