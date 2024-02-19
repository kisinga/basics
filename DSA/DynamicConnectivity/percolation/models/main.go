package models

type Data struct {
	// the grid
	grid [][]RootNode
}

type Node struct {
	Row int
	Col int
}

type RootNode struct {
	*Node
}

var VirtualRootTop = RootNode{
	Node: &Node{
		Row: -1,
		Col: -1,
	},
}

var VirtualRootBottom = RootNode{
	Node: &Node{
		Row: -2,
		Col: -2,
	},
}

type Percolation interface {
	// opens the site (row, col) if it is not open already
	Open(row int, col int)

	// is the site (row, col) open?
	IsOpen(row int, col int) bool

	// is the site (row, col) full?
	IsFull(row int, col int) bool

	// returns the number of open sites
	NumberOfOpenSites() int

	// does the system percolate?
	Percolates() bool
}

func (p *Data) AddNode(row int, col int) {
	p.grid[row][col] = RootNode{
		Node: &Node{
			Row: row,
			Col: col,
		},
	}
}

func (p *Data) Connect(node Node, rootNode RootNode) {
	p.grid[node.Row][node.Col] = rootNode
}

func (p *Data) Open(row int, col int) {

}

func (p Data) IsOpen(row int, col int) bool {
	return true
}

func (p Data) IsFull(row int, col int) bool {
	return true
}

func (p Data) NumberOfOpenSites() int {
	return 0
}

func (p Data) Percolates() bool {
	return true
}
