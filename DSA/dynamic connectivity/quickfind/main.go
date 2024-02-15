package quickfind

var arr [10]int

func main() {
	// connect several pairs of nodes and test if they are connected
	union(4, 3, arr[:])
}

func init() {
	arr = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func connected(p, q int, array []int) bool {
	return array[p] == array[q]
}

func union(p, q int, array []int) {
	pid := array[p]
	qid := array[q]

	for i := 0; i < len(array); i++ {
		if array[i] == pid {
			array[i] = qid
		}
	}
}
