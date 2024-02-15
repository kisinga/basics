package quickunion

var arr [10]int

func main() {

}

func init() {
	arr = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func root(i int, array []int) int {
	if i == array[i] {
		return i
	} else {
		return root(array[i], array)
	}
}

func connected(p, q int, array []int) bool {
	return root(p, array) == root(q, array)
}

func union(p, q int, array []int) {
	rootP := root(p, array)
	rootQ := root(q, array)
	array[rootP] = rootQ
}
