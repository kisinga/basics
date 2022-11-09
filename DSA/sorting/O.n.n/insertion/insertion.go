package insertion

func Sort(arr []int) {
	for i, val := range arr {
		pos := i - 1
		for j := i - 1; j >= 0; j-- {
			if arr[j] > val {
				j = 0
				continue
			}
			pos = j
		}
		if pos > 0 {
			arr[pos], arr[i] = arr[i], arr[pos]
		}
	}
}
