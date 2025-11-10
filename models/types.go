package models

type Graph struct {
	Adj      map[int][]int
	Incoming map[int][]int
	OutDeg   map[int]int
	Nodes    []int
}