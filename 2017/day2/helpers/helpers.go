package helpers

func Min(b []int) int {
	// return minimum value from slice
	var m int
	for i, v := range b {
		if i == 0 || v < m {
			m = v
		}
	}
	return m
}

func Max(b []int) int {
	// return maximum value from slice
	var m int
	for i, v := range b {
		if i == 0 || v > m {
			m = v
		}
	}
	return m
}
