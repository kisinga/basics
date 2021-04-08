package seletion

func Sort(arr []int) {

	// O(n)
	for i, val := range arr {
		// assume the index of the smallest value is the current array index
		smallestindex := i
		// O(n)
		for j := i + 1; j < len(arr); j++ {
			// compare all the values with the assumed smalled item
			if arr[smallestindex] > arr[j] {
				smallestindex = j
			}
		}
		// swap the smallest item to the current position
		arr[smallestindex], arr[i] = val, arr[smallestindex]
	}

}
