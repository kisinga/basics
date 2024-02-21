package quickunion

import "github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/models"

type QuickUnion struct {
}

func New(size int) *QuickUnion {
	return &QuickUnion{}
}

func init() {
	// for i := 0; i < 10; i++ {
	// 	improvedQuickUnion.arr[i] = i
	// 	improvedQuickUnion.sizearr[i] = 1
	// }
}

// a function that return the largest element in connected component containing p
// func (imp ImprovedQuickUnion) find(id int) int {
// 	max := 0
// 	for i := 0; i < len(imp.arr); i++ {
// 		if imp.connected(i, id) && i > max {
// 			max = i
// 		}
// 	}
// 	return max
// }

// func (imp ImprovedQuickUnion) root(id int) int {
// 	if id == imp.arr[id] {
// 		return id
// 	} else {
// 		return imp.root(imp.arr[id])
// 	}
// }

func (qu *QuickUnion) Connected(nodeA models.Node, nodeB models.Node) bool {
	// return imp.root(p) == imp.root(q)
	return true
}

func (qu *QuickUnion) Union(nodeA models.Node, nodeB models.Node) {
	// rootP := imp.root(p)
	// rootQ := imp.root(q)
	// if improvedQuickUnion.sizearr[rootP] < improvedQuickUnion.sizearr[rootQ] {
	// 	improvedQuickUnion.arr[rootP] = rootQ
	// 	improvedQuickUnion.sizearr[rootQ] += improvedQuickUnion.sizearr[rootP]
	// } else {
	// 	improvedQuickUnion.arr[rootQ] = rootP
	// 	improvedQuickUnion.sizearr[rootP] += improvedQuickUnion.sizearr[rootQ]
	// }
	// improvedQuickUnion.arr[rootP] = rootQ
}

func (qu *QuickUnion) AddNode(node models.Node) {
	// add a node to the grid
}

func (qu *QuickUnion) Open(node models.Node) {
	// open a node
}

func (qu *QuickUnion) IsOpen(node models.Node) bool {
	// check if a node is open
	return true
}

func (qu *QuickUnion) IsFull(node models.Node) bool {
	// check if a node is full
	return true
}

func (qu *QuickUnion) NumberOfOpenSites() int {
	// return the number of open sites
	return 0
}

func (qu *QuickUnion) Percolates() bool {
	// check if the system percolates
	return true
}

func (qu *QuickUnion) Connect(node models.Node, rootNode models.Node) {
	// connect a node to the root node
}
