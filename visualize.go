package main

import "fmt"

func VisualizeGraph(g *Graph) {
	if len(g.Nodes) > 50 {
		fmt.Println("Graph too large to display. Skipping visualization.")
		return
	}

	fmt.Println("\n--- Graph (Adjacency List) ---")
	for _, u := range g.Nodes {
		fmt.Printf("%d → %v\n", u, g.Adj[u])
	}
	fmt.Println("------------------------------\n")
}
