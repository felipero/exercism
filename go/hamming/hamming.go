//Package hamming calculates the hamming distance of 2 dna strings
package hamming

import (
	"errors"
)

// Distance calculates the hamming distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings doesn't have the same length")
	}

	distance := 0
	brunes := []rune(b)
	for pos, char := range []rune(a) {
		if brunes[pos] != char {
			distance++
		}
	}

	return distance, nil
}
