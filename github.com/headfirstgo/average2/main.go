package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(nums ...float64) float64 {
	var sum float64
	length := float64(len(nums))
	for _, num := range nums {
		sum += num
	}
	return sum/length
}

func main() {
	// get the command line argument via slice
	arguments := os.Args[1:] // because the first element is the cli path. also all elements are strings

	// initialize the sum variable
	var numbers []float64

	// loop through the cli slice using for loop and update sum variable
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		// error handling using log.Fatal
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	// use the printf() function to print the average
	fmt.Printf("average: %.2f\n", average(numbers...))
}
