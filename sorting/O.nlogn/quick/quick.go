package quick

func Sort(arr []int) {
	recusrsionSort(arr, 0, len(arr)-1)
}

// recusrsionSort calls itself with the left and right partitions of the array until the index+1 is >= than the end
// this essentially means that the left part and right part is the same
func recusrsionSort(arr []int, start int, end int) {
	if start < end {
		pIndex := partition(arr, start, end)
		recusrsionSort(arr, start, pIndex-1)
		recusrsionSort(arr, pIndex+1, end)

	}
}

// partition returns the index of the pivot element
// it assumes that the pivot element is the last element in this implementation
func partition(arr []int, start int, end int) int {
	pivot := arr[end]
	pIndex := 0
	for i := start; i <= end; i++ {
		if arr[i] < pivot {
			arr[i], arr[pIndex] = arr[pIndex], arr[i]
			pIndex++
		}
	}
	arr[pIndex], arr[end] = arr[end], arr[pIndex]
	return pIndex
}
