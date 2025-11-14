package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Graph struct
type Graph struct {
	Nodes map[string][]string
}

// LoadGraph reads edge list file and creates adjacency list
func LoadGraph(filename string) (*Graph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	g := &Graph{Nodes: make(map[string][]string)}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		from, to := parts[0], parts[1]
		g.Nodes[from] = append(g.Nodes[from], to)
		if _, ok := g.Nodes[to]; !ok {
			g.Nodes[to] = []string{}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	fmt.Printf("\nLoaded graph with %d nodes\n", len(g.Nodes))
	return g, nil
}
