package quickunion

import "github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/models"

type QuickUnion struct {
	// the grid
	data models.Data
}

func New(size int) models.Percolation {
	data := models.EmptyData(size)

	return &QuickUnion{
		data: *data,
	}
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

func (qu *QuickUnion) Open(node models.Node) {
	// opening a node means connecting it to the nearest open node
	qu.data.Grid[node.Row][node.Col].IsOpen = true

	if node.Row == 0 || node.Row == qu.data.RealSize-1 {
		return
	}

	// connect to top node
	if qu.IsOpen(models.Node{Row: node.Row - 1, Col: node.Col}) {
		qu.Connect(node, models.Node{Row: node.Row - 1, Col: node.Col})
	}

	// connect to bottom node
	if qu.IsOpen(models.Node{Row: node.Row + 1, Col: node.Col}) {
		qu.Connect(node, models.Node{Row: node.Row + 1, Col: node.Col})
	}

	// connect to left node
	if qu.IsOpen(models.Node{Row: node.Row, Col: node.Col - 1}) {
		qu.Connect(node, models.Node{Row: node.Row, Col: node.Col - 1})
	}

	// connect to right node
	if qu.IsOpen(models.Node{Row: node.Row, Col: node.Col + 1}) {
		qu.Connect(node, models.Node{Row: node.Row, Col: node.Col + 1})
	}
}

func (qu *QuickUnion) IsOpen(node models.Node) bool {
	// check if a node is open
	return node.IsOpen
}

func (qu *QuickUnion) IsFull(node models.Node) bool {
	// check if a node is full
	if qu.IsOpen(node) {
		// check if the node is connected to the virtual top root

		// if the node is in the first row, it is connected to the virtual top root
		if node.Row == 0 {
			return true
		}

		return true

		// if the node is in the middle, check if it is connected to the virtual top root

	} else {
		return false
	}
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
