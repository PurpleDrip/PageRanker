package main

func PageRankSerial(g *Graph, iterations int, d float64) map[int]float64 {
	N := float64(len(g.Nodes))
	rank := make(map[int]float64)
	newRank := make(map[int]float64)

	for _, v := range g.Nodes {
		rank[v] = 1.0 / N
	}

	for i := 0; i < iterations; i++ {
		for _, v := range g.Nodes {
			newRank[v] = (1 - d) / N
			for _, u := range g.Incoming[v] {
				newRank[v] += d * rank[u] / float64(g.OutDeg[u])
			}
		}
		for _, v := range g.Nodes {
			rank[v] = newRank[v]
		}
	}

	return rank
}
