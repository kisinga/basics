package percolation

import "github.com/kisinga/basics/DSA/DynamicConnectivity/percolation/models"

func New(n int) models.Data {
	// n := models.Node{
	// 	Row: 0,
	// 	Col: 0,
	// }
	return models.Data{
		// Grid: make([][]models.Node, n),
	}
}

func main() {
	var data models.Percolation = New(10)
	data.Open(1, 1)

}
