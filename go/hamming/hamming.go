//Package hamming calculates the hamming distance of 2 dna strings
package hamming

import (
	"errors"
)

// Distance calculates the hamming distance between two strings
func Distance(a, b string) (int, error) {
	distance := 0
	brunes := []rune(b)
	var err error

	if len(a) != len(b) {
		err = errors.New("strings doesn't have the same length")
	} else {
		for pos, char := range a {
			if brunes[pos] != char {
				distance++
			}
		}
	}
	return distance, err
}
