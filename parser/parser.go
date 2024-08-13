package parser

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
	"strings"
)

// URLValue holds the URL and its associated value
type URLValue struct {
	URL   string
	Value int
}

// MinHeap is a min-heap of URLValue
type MinHeap []URLValue

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Value < h[j].Value }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(URLValue))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func ParseFile(filePath string, topN int) ([]string, error) {

	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	h := &MinHeap{}
	heap.Init(h)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if len(parts) != 2 {
			continue
		}

		url := parts[0]
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}

		if h.Len() < topN {
			heap.Push(h, URLValue{URL: url, Value: value})
		} else {
			if value > (*h)[0].Value {
				heap.Pop(h)
				heap.Push(h, URLValue{URL: url, Value: value})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Extract URLs from the heap
	h_result := make([]URLValue, 0, h.Len())
	for h.Len() > 0 {
		h_result = append(h_result, heap.Pop(h).(URLValue))
	}

	// Print in reverse order (largest value first)
	result := []string{}
	for i := len(h_result) - 1; i >= 0; i-- {
		result = append(result, h_result[i].URL)
	}

	return result, nil

}
