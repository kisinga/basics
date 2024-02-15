package quickunion

type ImprovedQuickUnion struct {
	arr     [10]int
	sizearr [10]int
}

var improvedQuickUnion ImprovedQuickUnion

func improvedmain() {

}

func init() {
	for i := 0; i < 10; i++ {
		improvedQuickUnion.arr[i] = i
		improvedQuickUnion.sizearr[i] = 1
	}
}

// a function that return the largest element in connected component containing p
func (imp ImprovedQuickUnion) find(id int) int {
	max := 0
	for i := 0; i < len(imp.arr); i++ {
		if imp.connected(i, id) && i > max {
			max = i
		}
	}
	return max
}

func (imp ImprovedQuickUnion) root(id int) int {
	if id == imp.arr[id] {
		return id
	} else {
		return imp.root(imp.arr[id])
	}
}

func (imp ImprovedQuickUnion) connected(p, q int) bool {
	return imp.root(p) == imp.root(q)
}

func (imp ImprovedQuickUnion) union(p, q int) {
	rootP := imp.root(p)
	rootQ := imp.root(q)
	if improvedQuickUnion.sizearr[rootP] < improvedQuickUnion.sizearr[rootQ] {
		improvedQuickUnion.arr[rootP] = rootQ
		improvedQuickUnion.sizearr[rootQ] += improvedQuickUnion.sizearr[rootP]
	} else {
		improvedQuickUnion.arr[rootQ] = rootP
		improvedQuickUnion.sizearr[rootP] += improvedQuickUnion.sizearr[rootQ]
	}
	improvedQuickUnion.arr[rootP] = rootQ
}
