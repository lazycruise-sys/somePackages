package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
	"github.com/headfirstgo/delete"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)
	sampleCount := float64(len(numbers))
	fmt.Printf("average: %.2f\n", sum/sampleCount)

	pack := delete.Pack(sum / sampleCount)
	fmt.Printf("pack value: %.3f\n", pack)
}
