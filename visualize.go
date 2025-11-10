package main

import "fmt"

// VisualizeGraph prints adjacency list if graph is small
func VisualizeGraph(g *Graph) {
	fmt.Println("Graph visualization:")
	for node, neighbors := range g.Nodes {
		fmt.Printf("%s -> %v\n", node, neighbors)
	}
}
