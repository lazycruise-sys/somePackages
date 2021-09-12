// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"os"
)

// GetFloats reads a string from each line of a file
func GetStrings(fileName string) ([]string, error) {
	var lines []string
	// open the data file for reading
	file, err := os.Open(fileName) // open the provided filename
	if err != nil {                // if there was an error openining the file, return it
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // convert the file line string to a float64
		lines = append(lines, line)
	}

	// "for i := 0; scanner.Scan(); i++" could have been used (learning purpose)

	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil // return the arrays of numbers and a "nil" error
}
