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

func (qu *QuickUnion) Open(Row, Col int) {

}

func (qu *QuickUnion) IsOpen(Row, Col int) bool {
	// check if a node is open
	return qu.data.Grid[Row][Col].IsOpen
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

func (qu *QuickUnion) Connect(Row1, Col1, Row, Col int) {
	// connect a node to the root node
}
