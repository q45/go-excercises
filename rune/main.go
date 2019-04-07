package main

import "errors"

func main() {

	Distance("AAA", "AGAT")
}

func Distance(a, b string) (int, error) {

	s1 := []rune(a)
	s2 := []rune(b)

	if len(s1) != len(s2) {
		return 0, errors.New("mismatch length strings not accepted")
	}

	hammingDistance := 0

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
