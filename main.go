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

	var choice int
	fmt.Print("Enter: ")
	fmt.Scan(&choice)

	path := ""
	switch choice {
	case 1:
		path = "datasets/small.txt"
	case 2:
		path = "datasets/web-Stanford.txt"
	case 3:
		path = "datasets/web-Google.txt"
	default:
		fmt.Println("Invalid choice")
		return
	}

	graph, err := LoadGraph(path)
	if err != nil {
		panic(err)
	}

	VisualizeGraph(graph)

	fmt.Println("Running Serial PageRank...")
	start := time.Now()
	r1 := PageRankSerial(graph, 20, 0.85)
	fmt.Println("Time:", time.Since(start))
	PrintTop(r1, 10)

	fmt.Println("\nRunning Parallel PageRank (10 workers)...")
	start = time.Now()
	r2 := PageRankParallel(graph, 20, 0.85, 10)
	fmt.Println("Time:", time.Since(start))
	PrintTop(r2, 10)
}
