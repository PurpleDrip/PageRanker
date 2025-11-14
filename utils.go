package main

import "fmt"

// PrintTop prints top 10 ranked nodes
func PrintTop(rank map[string]float64) {

	fmt.Println("\n\tSorting the Ranked Map...")
	type kv struct {
		Key   string
		Value float64
	}
	var sorted []kv
	for k, v := range rank {
		sorted = append(sorted, kv{k, v})
	}

	for i := 0; i < len(sorted)-1; i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j].Value > sorted[i].Value {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	fmt.Println("\nTop Ranked Nodes:")
	for i := 0; i < 10 && i < len(sorted); i++ {
		fmt.Printf("%s → %.6f\n", sorted[i].Key, sorted[i].Value)
	}
}
