package bubble

func Sort(arr []int) {
	// create the outer loop
	for i, val := range arr {
		// create an inner for loop
		for j := 0; j < len(arr); j++ {
			// compare the values
			if val < arr[j] {
				// swap the values
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
