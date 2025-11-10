package main

import (
	"sync"
)

func PageRankParallel(g *Graph, iterations int, d float64, workers int) map[int]float64 {
	N := float64(len(g.Nodes))
	rank := make(map[int]float64)
	newRank := make(map[int]float64)

	for _, v := range g.Nodes {
		rank[v] = 1.0 / N
	}

	chunk := len(g.Nodes) / workers

	for iter := 0; iter < iterations; iter++ {

		var wg sync.WaitGroup
		wg.Add(workers)

		for w := 0; w < workers; w++ {
			start := w * chunk
			end := start + chunk
			if w == workers-1 {
				end = len(g.Nodes)
			}

			go func(start, end int) {
				defer wg.Done()
				for i := start; i < end; i++ {
					v := g.Nodes[i]
					newRank[v] = (1 - d) / N
					for _, u := range g.Incoming[v] {
						newRank[v] += d * rank[u] / float64(g.OutDeg[u])
					}
				}
			}(start, end)
		}

		wg.Wait()

		for _, v := range g.Nodes {
			rank[v] = newRank[v]
		}
	}

	return rank
}
