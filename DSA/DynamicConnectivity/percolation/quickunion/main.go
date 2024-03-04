package quickunion

import "github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/models"

type QuickUnion struct {
	// the grid
	models.Grid
}

func New(size int) (models.Percolation, error) {
	data, err := models.EmptyData(size)
	if err != nil {
		return nil, err
	}

	return &QuickUnion{
		data,
	}, nil
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

func (qu *QuickUnion) Open(Row, Col int) error {
	return nil
}

func (qf *QuickUnion) FindRoot(row, col int) models.Node {
	// Starting at the given site
	current := qf.Grid[row][col]

	// Follow the parent nodes until finding a root
	// A root node is indicated by it pointing to itself
	// Assuming each Node knows its parent, adjust accordingly if your structure is different
	for !(current.Row == row && current.Col == col) {
		row = current.Row
		col = current.Col
		current = qf.Grid[row][col]
	}

	return current
}

func (qu *QuickUnion) IsOpen(Row, Col int) bool {
	// check if a node is open
	return qu.Grid[Row][Col].IsOpen
}

func (qu *QuickUnion) IsFull(Row, Col int) bool {

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

func (qu *QuickUnion) Connect(RowA, ColA, RowB, ColB int) {
	// connect a node to the root node
}
