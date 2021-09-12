package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/headfirstgo/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	candidates := make(map[string]int)

	for _, name := range lines {
		candidates[name]++
	}

	// due to the randomness of maps, it is good to rearrange via the use of a slice
	var names []string // to store the candidates name into a slice
	for candidate := range candidates {
		names = append(names, candidate)
	}

	sort.Strings(names) // rearranges the slice alphabetically

	for _, name := range names {
		fmt.Printf("Vote for %s: %d\n", name, candidates[name])
	}

	// // if randomness of the maps isn't important, use the following syntax
	// for candidate, vote := range candidates {
	// 	fmt.Printf("%s: %d\n", candidate, vote)
	// }
}
