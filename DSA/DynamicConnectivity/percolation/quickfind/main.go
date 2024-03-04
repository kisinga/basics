package quickfind

import (
	"errors"

	"github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/models"
)

type QuickFind struct {
	// the grid, where the arrangement is row, col
	models.Grid
}

func New(size int) (models.Percolation, error) {
	data, err := models.EmptyData(size)
	if err != nil {
		return nil, err
	}

	return &QuickFind{
		data,
	}, nil
}

func (qf *QuickFind) Open(Row, Col int) error {

	if Row < 1 || Row > len(qf.Grid)-2 || Col > len(qf.Grid[1]) {
		return errors.New("You cannot open a site outside boundaries, vtop or vbottom")
	}

	//opening a node means that it will be connected to all the neighbours that are also open
	qf.Grid[Row][Col].IsOpen = true

	// connect to immediate neighbor top
	if Row > 1 && qf.IsOpen(Row-1, Col) {
		qf.Connect(Row, Col, Row-1, Col)
	}

	// connect to immediate neighbor bottom
	if Row < (len(qf.Grid)-2) && qf.IsOpen(Row+1, Col) {
		qf.Connect(Row, Col, Row+1, Col)
	}

	// connect to immediate neighbor left
	if Col > 0 && qf.IsOpen(Row, Col-1) {
		qf.Connect(Row, Col, Row, Col-1)
	}

	// connect to immediate neighbor right
	if Col < (len(qf.Grid[1])-1) && qf.IsOpen(Row, Col+1) {
		qf.Connect(Row, Col, Row, Col+1)
	}
	return nil
}

func (qf *QuickFind) IsOpen(Row, Col int) bool {
	if Row == 0 || Row == len(qf.Grid)-1 {
		return true
	}
	return qf.Grid[Row][Col].IsOpen
}

func (qf *QuickFind) IsFull(Row, Col int) bool {
	// check if a node is full
	return true
}

func (qf *QuickFind) NumberOfOpenSites() int {
	// return the number of open sites
	count := 0
	for i := 1; i < len(qf.Grid)-1; i++ {
		// for every row, check if any element has a parent equal to the rootA
		for _, val := range qf.Grid[i] {
			if val.IsOpen {
				count++
			}
		}
	}
	return count
}

func (qf *QuickFind) Percolates() bool {
	return qf.Grid[0][0].Row == qf.Grid[len(qf.Grid)-1][0].Row && qf.Grid[0][0].Col == qf.Grid[len(qf.Grid)-1][0].Col

}

func (qf *QuickFind) Connect(rowA, colA, rowB, colB int) {

	var rootA models.Node
	var rootB models.Node

	rootA = qf.Grid[rowA][colA]

	rootB = qf.Grid[rowB][colB]

	// connect the two nodes
	for i := range len(qf.Grid) {
		// for every row, check if any element has a parent equal to the rootA
		for j := range len(qf.Grid[i]) {
			if qf.Grid[i][j].Row == rootA.Row && qf.Grid[i][j].Col == rootA.Col {
				// dont overwrite the entire root because if the node was closed (e.g nodes in the first row)
				// but it was connected to another node which also has childres (e.g the vtop), then we need to maintain the nodes isOpen state
				// so that we dont possibly create a false open node in a previously closed node just because the root changed
				qf.Grid[i][j].Row = rootB.Row
				qf.Grid[i][j].Col = rootB.Col
			}
		}
	}

}
