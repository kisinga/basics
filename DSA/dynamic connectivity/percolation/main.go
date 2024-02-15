package percolation

type percolationData struct {
	// the grid
	grid [][]int
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

func (p percolationData) Open(row int, col int) {

}

func (p percolationData) IsOpen(row int, col int) bool {
	return true
}

func (p percolationData) IsFull(row int, col int) bool {
	return true
}

func (p percolationData) NumberOfOpenSites() int {
	return 0
}

func (p percolationData) Percolates() bool {
	return true
}

func New(n int) percolationData {
	return percolationData{
		grid: make([][]int, n),
	}
}

func main() {
	var data Percolation = New(10)
	data.Open(1, 1)

}
