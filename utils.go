package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
    key   string
    value float64
}

type MinHeap []Item

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].value < h[j].value }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}


// PrintTop prints top 10 ranked nodes
func PrintTop(rank map[string]float64) {
    h := &MinHeap{}
    heap.Init(h)

    for k, v := range rank {
        if h.Len() < 10 {
            heap.Push(h, Item{k, v})
        } else if v > (*h)[0].value {
            heap.Pop(h)
            heap.Push(h, Item{k, v})
        }
    }

    fmt.Println("\nTop 10 Ranked Nodes:")
    for h.Len() > 0 {
        item := heap.Pop(h).(Item)
        fmt.Printf("%s → %.6f\n", item.key, item.value)
    }
}

