package quickfind

import (
	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/models"
)

type QuickFind struct {
	// the data, where the arrangement is row, col
	data models.Data
}

func New(size int) models.Percolation {
	data := models.EmptyData(size)
	return &QuickFind{
		data: *data,
	}
}

func (qf *QuickFind) Connected(nodeA models.Node, nodeB models.Node) bool {
	// return array[p] == array[q]
	// two nodes are connected if they have the same root
	return qf.data.Grid[nodeA.Row][nodeA.Col] == qf.data.Grid[nodeA.Row][nodeA.Col]
}

func (qf *QuickFind) Open(node models.Node) {
	//opening a node means that it will be connected to all the neighbours that are also open
	qf.data.Grid[node.Row][node.Col].IsOpen = true

	if node.Row == 0 || node.Row == qf.data.RealSize-1 {
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
	return qf.data.Grid[node.Row][node.Col].IsOpen
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
	return qf.data.Grid[qf.data.VTopPos][0] == qf.data.Grid[qf.data.VBottomPos][0]
}

func (qf *QuickFind) Connect(nodeA models.Node, nodeB models.Node) {
	// connect the two nodes

	rootA := qf.data.Grid[nodeA.Row][nodeA.Col]
	rootB := qf.data.Grid[nodeB.Row][nodeB.Col]

	if rootA != rootB {
		for i := range qf.data.RealSize {
			// for every row, check if any element has a parent equal to the rootA
			for j := range qf.data.RealSize {
				if qf.data.Grid[i][j] == rootA {
					qf.data.Grid[i][j] = rootB
				}
			}
		}
	}

}
