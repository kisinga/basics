package quickfind

import (
	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/models"
)

type QuickFind struct {
	// the grid, where the arrangement is row, col
	grid models.Grid
}

func New(size int) models.Percolation {
	grid := models.EmptyGrid(size)

	// connect the first col to virtual top

	for i := range size {
		grid[0][i] = models.VirtualRootTop
		grid[size-1][i] = models.VirtualRootBottom
	}

	// connect the last row to virtual bottom

	return &QuickFind{
		grid: grid,
	}
}

func (qf *QuickFind) Connected(nodeA models.Node, nodeB models.Node) bool {
	// return array[p] == array[q]
	// two nodes are connected if they have the same root
	return qf.grid[nodeA.Row][nodeA.Col] == qf.grid[nodeA.Row][nodeA.Col]
}

func (qf *QuickFind) Open(node models.Node) {
	//opening a node means that it will be connected to all the neighbours that are also open
	qf.grid[node.Row][node.Col].IsOpen = true

	if node.Row == 0 || node.Row == len(qf.grid)-1 {
		return
	}

	// connect to top node
	if qf.IsOpen(models.Node{Row: node.Row - 1, Col: node.Col}) {
		qf.Connect(node, models.Node{Row: node.Row - 1, Col: node.Col})
	}

	// connect to bottom node
	if qf.IsOpen(models.Node{Row: node.Row + 1, Col: node.Col}) {
		qf.Connect(node, models.Node{Row: node.Row + 1, Col: node.Col})
	}

	// connect to left node
	if qf.IsOpen(models.Node{Row: node.Row, Col: node.Col - 1}) {
		qf.Connect(node, models.Node{Row: node.Row, Col: node.Col - 1})
	}

	// connect to right node
	if qf.IsOpen(models.Node{Row: node.Row, Col: node.Col + 1}) {
		qf.Connect(node, models.Node{Row: node.Row, Col: node.Col + 1})
	}

}

func (qf *QuickFind) IsOpen(node models.Node) bool {
	// a node is open if it's not connected to itself
	return qf.grid[node.Row][node.Col].IsOpen
}

func (qf *QuickFind) IsFull(node models.Node) bool {
	// check if a node is full
	return true
}

func (qf *QuickFind) NumberOfOpenSites() int {
	// return the number of open sites
	return 0
}

func (qf *QuickFind) Percolates() bool {
	// check if the system percolates
	return true
}

func (qf *QuickFind) Connect(nodeA models.Node, nodeB models.Node) {
	// connect the two nodes

	rootA := qf.grid[nodeA.Row][nodeA.Col]
	rootB := qf.grid[nodeB.Row][nodeB.Col]

	if rootA != rootB {
		for i := range len(qf.grid) {
			// for every row, check if any element has a parent equal to the rootA
			for j := range len(qf.grid) {
				if qf.grid[i][j] == rootA {
					qf.grid[i][j] = rootB
				}
			}
		}
	}

}
