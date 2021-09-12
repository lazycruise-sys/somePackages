// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of a file
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	// open the data file for reading
	file, err := os.Open(fileName) // open the provided filename
	if err != nil {                // if there was an error openining the file, return it
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp, err := strconv.ParseFloat(scanner.Text(), 64) // convert the file line string to a float64
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, temp)
	}

	// "for i := 0; scanner.Scan(); i++" could have been used (learning purpose)

	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil // return the arrays of numbers and a "nil" error
}
