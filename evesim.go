package main

import (
	"fmt"
	"math/rand"
)

func main() {
	remainingChildren := map[string]int{
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

	ng := 100     // number of generations
	maxGirls := 4 // maximum number of children per generation

	for i := 0; i < ng; i++ {
		for geneticLine, children := range remainingChildren {
			remainingChildren[geneticLine] = 0 // all mothers will die
			for _ = children; children > 0; children-- {
				remainingChildren[geneticLine] += rand.Intn(maxGirls)
			}
			//fmt.Printf("%s had %d children\n", geneticLine, remainingChildren[geneticLine])
		}
		if (i % 10) == 0 {
			fmt.Printf("%+v\n", remainingChildren)
		}

	}
	fmt.Printf("%+v\n", remainingChildren)

}
