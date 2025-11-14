package main

import (
	"fmt"
	"time"
)

// Serial PageRank
func PageRankSerial(g *Graph, iterations int, damping float64) map[string]float64 {
	t0 := time.Now()

	n := len(g.Nodes)
	rank := make(map[string]float64)
	for node := range g.Nodes {
		rank[node] = 1.0 / float64(n)
	}

	for i := 0; i < iterations; i++ {
		newRank := make(map[string]float64)
		for node := range g.Nodes {
			newRank[node] = (1 - damping) / float64(n)
		}
		for node, neighbors := range g.Nodes {
			if len(neighbors) == 0 {
				continue
			}
			share := rank[node] * damping / float64(len(neighbors))
			for _, neighbor := range neighbors {
				newRank[neighbor] += share
			}
		}
		rank = newRank
	}

    fmt.Printf("Serial PageRank completed in %v\n", time.Since(t0))
	return rank
}
