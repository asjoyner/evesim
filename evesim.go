package main

import (
	"fmt"
	"math/rand"
)

type result struct {
	lines       []int64
	total       int64
	maximum     int64
	generations int
}

func evesim() result {
	var r result
	initialPopulation := 10 // number of initial lines
	ng := 4000              // number of generations
	maxGirls := 3.1         // maximum number of children per generation

	r.lines = make([]int64, initialPopulation)
	for i := range r.lines {
		r.lines[i] = 1
	}

	// Simulate as many generation as is takes for one lineage to get to 8
	// billion children.  Running too many generations results in overflowing an
	// int64.  :)
	var winner bool
	for i := 0; i < ng && !winner; i++ {
		for lineage, children := range r.lines {
			avgChildren := float64(maxGirls) * rand.Float64()
			r.lines[lineage] = int64(avgChildren * float64(children))
			if r.lines[lineage] >= 8000000000 {
				winner = true
				r.generations = i
			}
		}
	}

	// calculate some aggregate statistics about this result
	for _, children := range r.lines {
		r.total += children
		if children > r.maximum {
			r.maximum = children
		}
	}
	return r
}

func main() {
	numSims := 1000
	results := make([]result, numSims)
	for i := 0; i < numSims; i++ {
		results[i] = evesim()
	}

	var failed int
	var generations int
	var maxGeneration int
	for _, r := range results {
		if r.maximum == 0 {
			failed++
		}
		generations += r.generations
		if r.generations > maxGeneration {
			maxGeneration = r.generations
		}
	}
	if failed > 0 {
		f := float64(failed) / float64(numSims) * 100
		fmt.Printf("%2.0f%% of the populations died out.\n", f)
	}
	g := float64(generations) / float64(numSims)
	fmt.Printf("Averaged %f generations.\n", g)
	fmt.Printf("Last generation: %d.\n", maxGeneration)

	// Calculate how many of the simulations had a second mitochondrial linage
	// had greater than a certain percentage of the children.
	for _, threshold := range []float64{0.01, 0.1, 0.25} {
		var passing int
		for _, r := range results {
			for _, children := range r.lines {
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
