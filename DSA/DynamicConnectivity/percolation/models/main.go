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
