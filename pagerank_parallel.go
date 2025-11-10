package main

import (
	"sync"
)

// Parallel PageRank with mutex
func PageRankParallel(g *Graph, iterations int, damping float64, workers int) map[string]float64 {
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

		var wg sync.WaitGroup
		var mu sync.Mutex

		nodes := make([]string, 0, len(g.Nodes))
		for node := range g.Nodes {
			nodes = append(nodes, node)
		}

		chunkSize := (len(nodes) + workers - 1) / workers
		for w := 0; w < workers; w++ {
			start := w * chunkSize
			end := start + chunkSize
			if end > len(nodes) {
				end = len(nodes)
			}
			wg.Add(1)
			go func(subnodes []string) {
				defer wg.Done()
				for _, node := range subnodes {
					neighbors := g.Nodes[node]
					if len(neighbors) == 0 {
						continue
					}
					share := rank[node] * damping / float64(len(neighbors))
					for _, neighbor := range neighbors {
						mu.Lock()
						newRank[neighbor] += share
						mu.Unlock()
					}
				}
			}(nodes[start:end])
		}
		wg.Wait()
		rank = newRank
	}

	return rank
}
