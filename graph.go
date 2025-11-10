package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	Adj      map[int][]int
	Incoming map[int][]int
	OutDeg   map[int]int
	Nodes    []int
}

func LoadGraph(path string) (*Graph, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	g := &Graph{
		Adj:      make(map[int][]int),
		Incoming: make(map[int][]int),
		OutDeg:   make(map[int]int),
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || len(strings.TrimSpace(line)) == 0 {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}

		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])

		g.Adj[u] = append(g.Adj[u], v)
		g.Incoming[v] = append(g.Incoming[v], u)
		g.OutDeg[u]++
	}

	for node := range g.Adj {
		g.Nodes = append(g.Nodes, node)
	}

	return g, nil
}
