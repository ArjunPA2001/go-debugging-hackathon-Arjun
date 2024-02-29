package main

// find average given function parameters
// no return statments allowed
func average(a, b, c int, pavg *int) {

	*pavg = (a + b + c) / 3
}

// find min and maxmimum from the given function parameters
//
func min_max(a, b, c int) (int, int) {

	min := a
	max := a

	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}

	return min, max
}
