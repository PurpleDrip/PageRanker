package main

import (
	"fmt"
	"sort"
)

func PrintTop(rank map[int]float64, k int) {
	type Pair struct {
		Node int
		Score float64
	}
	var arr []Pair
	for n, v := range rank {
		arr = append(arr, Pair{n, v})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Score > arr[j].Score
	})

	fmt.Println("\nTop Ranked Nodes:")
	for i := 0; i < k && i < len(arr); i++ {
		fmt.Printf("%d → %.6f\n", arr[i].Node, arr[i].Score)
	}
}
