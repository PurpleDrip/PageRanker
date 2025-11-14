package main

import (
	"fmt"
	"sync"
	"time"
)

func PageRankParallel(g *Graph, iterations int, damping float64, workers int) map[string]float64 {
    t0 := time.Now()

    n := len(g.Nodes)

    rank := make(map[string]float64, n)
    for node := range g.Nodes {
        rank[node] = 1.0 / float64(n)
    }

    nodes := make([]string, 0, n)
    for node := range g.Nodes {
        nodes = append(nodes, node)
    }

    chunkSize := (n + workers - 1) / workers

    for it := 0; it < iterations; it++ {

        newRank := make(map[string]float64, n)
        base := (1 - damping) / float64(n)
        for _, node := range nodes {
            newRank[node] = base
        }

        var wg sync.WaitGroup
        wg.Add(workers)

        partial := make([]map[string]float64, workers)
        for i := 0; i < workers; i++ {
            partial[i] = make(map[string]float64)
        }

        for w := 0; w < workers; w++ {
            go func(w int) {
                defer wg.Done()

                start := w * chunkSize
                if start >= n {
                    return
                }
                end := start + chunkSize
                if end > n {
                    end = n
                }

                local := partial[w]

                for i := start; i < end; i++ {
                    node := nodes[i]
                    neighbors := g.Nodes[node]

                    if len(neighbors) == 0 {
                        continue
                    }

                    share := rank[node] * damping / float64(len(neighbors))

                    for _, nb := range neighbors {
                        local[nb] += share
                    }
                }

            }(w)
        }

        wg.Wait()

        for w := 0; w < workers; w++ {
            for node, val := range partial[w] {
                newRank[node] += val
            }
        }

        rank = newRank
    }

    fmt.Printf("Parallel PageRank completed in %v\n", time.Since(t0))
    return rank;
}

