package hamming

import "errors"

// Distance will calculate the minimum number of point mutations that could occur on two strands
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("string is longer that the other")
	}

	counter := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			counter++
		}
	}

	return counter, nil
}
