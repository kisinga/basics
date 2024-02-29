package quickfind

import (
	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/models"
)

type QuickFind struct {
	// the data, where the arrangement is row, col
	data *models.Data
}

func New(size int) models.Percolation {
	data := models.EmptyData(size)
	return &QuickFind{
		data: data,
	}
}

func (qf *QuickFind) Open(Row, Col int) {
	//opening a node means that it will be connected to all the neighbours that are also open
	qf.data.Grid[Row][Col].IsOpen = true

	// connect to top node
	if Row != 0 && qf.IsOpen(Row-1, Col) {
		qf.Connect(Row, Col, Row-1, Col)
	}

	// connect to bottom node
	if Row != qf.data.RealSize-1 && qf.IsOpen(Row+1, Col) {
		qf.Connect(Row, Col, Row+1, Col)
	}

	// escape early for first and last cols
	if Col == 0 || Col == qf.data.RealSize-1 {
		return
	}

	// connect to left node
	if qf.IsOpen(Row, Col-1) {
		qf.Connect(Row, Col, Row, Col-1)
	}

	// connect to right node
	if qf.IsOpen(Row, Col+1) {
		qf.Connect(Row, Col, Row, Col+1)
	}

}

func (qf *QuickFind) IsOpen(Row, Col int) bool {
	return qf.data.Grid[Row][Col].IsOpen
}

func (qf *QuickFind) IsFull(Row, Col int) bool {
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

func (qf *QuickFind) Connect(Row1, Col1, Row, Col int) {
	// connect the two nodes

	rootA := qf.data.Grid[Row1][Col1]
	rootB := qf.data.Grid[Row][Col]

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
