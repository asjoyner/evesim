package main

import (
	"fmt"
	"math/rand"
)

func main() {
	remainingChildren := map[string]int64{
		"Eve":      1,
		"Deborah":  1,
		"Sue":      1,
		"Jane":     1,
		"Suzie":    1,
		"Sophia":   1,
		"Emma":     1,
		"Olivia":   1,
		"Ava":      1,
		"Isabella": 1,
	}

	ng := 4000      // number of generations
	maxGirls := 3.7 // maximum number of children per generation

	var winner bool
	for i := 0; i < ng && !winner; i++ {
		for geneticLine, children := range remainingChildren {
			avgChildren := float64(maxGirls) * rand.Float64()
			remainingChildren[geneticLine] = int64(avgChildren * float64(children))
			if remainingChildren[geneticLine] >= 1000000000 {
				fmt.Printf("%s wins with %d children in %d generations.\n", geneticLine, remainingChildren[geneticLine], i)
				winner = true
			}
		}
	}
	var total int64
	for _, children := range remainingChildren {
		total += children
	}
	for geneticLine, children := range remainingChildren {
		fmt.Printf("%s: %d (%0.2f%%)\n", geneticLine, children, float64(children)/float64(total)*100)
	}
}
