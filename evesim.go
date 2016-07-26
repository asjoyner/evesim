package main

import (
	"fmt"
	"math/rand"
)

type result struct {
	children map[string]int64
	total    int64
	maximum  int64
}

func evesim() map[string]int64 {
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
	maxGirls := 3.5 // maximum number of children per generation

	var winner bool
	for i := 0; i < ng && !winner; i++ {
		for lineage, children := range remainingChildren {
			avgChildren := float64(maxGirls) * rand.Float64()
			remainingChildren[lineage] = int64(avgChildren * float64(children))
			if remainingChildren[lineage] >= 1000000000 {
				//fmt.Printf("%s wins with %d children in %d generations.\n", lineage, remainingChildren[lineage], i)
				winner = true
			}
		}
	}
	return remainingChildren
}

func total(r map[string]int64) int64 {
	var total int64
	for _, children := range r {
		total += children
	}
	return total
}

func maximum(r map[string]int64) int64 {
	var max int64
	for _, children := range r {
		if children > max {
			max = children
		}
	}
	return max
}

func main() {
	numSims := 1000
	results := make([]result, numSims)
	for i := 0; i < numSims; i++ {
		r := evesim()
		results[i] = result{
			children: r,
			total:    total(r),
			maximum:  maximum(r),
		}
	}

	var failed int
	for _, r := range results {
		if r.maximum == 0 {
			failed++
		}
	}
	if failed > 0 {
		f := float64(failed) / float64(numSims) * 100
		fmt.Printf("%2.0f%% of the populations died out.\n", f)
	}

	for _, threshold := range []float64{0.01, 0.2, 0.5} {
		var passing int
		for _, r := range results {
			for _, children := range r.children {
				if children == r.maximum {
					continue // skip the winner
				}
				if (float64(children) / float64(r.total)) > threshold {
					passing++
					break
				}
			}
		}
		p := float64(passing) / float64(numSims) * 100
		fmt.Printf("Odds of a second lineage of >%2.0f%% of the population: %5.2f %%\n", threshold*100.0, p)
	}
}
