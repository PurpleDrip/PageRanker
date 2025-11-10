package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Select Dataset:")
	fmt.Println("1) small.txt")
	fmt.Println("2) web-Stanford.txt")
	fmt.Println("3) web-Google.txt")
	fmt.Print("Enter: ")

	var choice int
	fmt.Scan(&choice)

	files := map[int]string{
		1: "datasets/small.txt",
		2: "datasets/web-Stanford.txt",
		3: "datasets/web-Google.txt",
	}

	filename, ok := files[choice]
	if !ok {
		fmt.Println("Invalid choice")
		return
	}

	g, err := LoadGraph(filename)
	if err != nil {
		fmt.Println("Error loading graph:", err)
		return
	}

	if len(g.Nodes) < 50 {
		VisualizeGraph(g)
	} else {
		fmt.Println("Graph too large to display. Skipping visualization.")
	}

	// Serial PageRank
	fmt.Println("\nRunning Serial PageRank...")
	start := time.Now()
	rankSerial := PageRankSerial(g, 20, 0.85)
	fmt.Println("Time:", time.Since(start))
	PrintTop(rankSerial)

	// Parallel PageRank
	fmt.Println("\nRunning Parallel PageRank (10 workers)...")
	start = time.Now()
	rankParallel := PageRankParallel(g, 20, 0.85, 10)
	fmt.Println("Time:", time.Since(start))
	PrintTop(rankParallel)
}
