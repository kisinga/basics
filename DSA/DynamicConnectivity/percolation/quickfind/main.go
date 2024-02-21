package quickfind

import "github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/models"

type QuickFind struct {
	// the grid
	grid [][]models.Node
}

func New(size int) models.Percolation {
	return &QuickFind{
		grid: make([][]models.Node, size),
	}
}

func (qf *QuickFind) Connected(nodeA models.Node, nodeB models.Node) bool {
	// return array[p] == array[q]

	return true
}

func (qf *QuickFind) Union(nodeA models.Node, nodeB models.Node) {
	// pid := array[p]
	// qid := array[q]

	// for i := 0; i < len(array); i++ {
	// 	if array[i] == pid {
	// 		array[i] = qid
	// 	}
	// }
}

func (qf *QuickFind) AddNode(node models.Node) {
	// add a node to the grid
}

func (qf *QuickFind) Open(node models.Node) {
	// open a node
}

func (qf *QuickFind) IsOpen(node models.Node) bool {
	// check if a node is open
	return true
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

func (qf *QuickFind) Connect(node models.Node, rootNode models.Node) {
	// connect a node to the root node
}
