/*
Package hamming provides Distance function to determin the hamming distance between two strings
*/
package hamming

import "errors"

//Distance calculated the hamming distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("strings are not the same length")
	}

	hammingDistance := 0

	for i := range a {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
