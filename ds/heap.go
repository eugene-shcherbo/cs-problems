package ds

import (
	"cmp"
	"fmt"
)

type MinHeap[T any] struct {
	items []T
	zero  T
	less  func(a, b T) bool
}

func NewHeap[T any](items []T, less func(a, b T) bool) *MinHeap[T] {
	itemsCopy := make([]T, len(items))
	copy(itemsCopy, items)

	var zero T
	heap := &MinHeap[T]{itemsCopy, zero, less}
	for i := len(heap.items)/2 - 1; i >= 0; i-- {
		heapify(heap, 0)
	}

	return heap
}

func NewHeapComp[T cmp.Ordered](items []T) *MinHeap[T] {
	return NewHeap(items, func(a, b T) bool { return a < b })
}

func (h *MinHeap[T]) Add(el T) {
	h.items = append(h.items, el)

	idx := len(h.items) - 1
	for h.less(h.items[idx], h.items[parentIdx(idx)]) {
		h.items[idx], h.items[parentIdx(idx)] = h.items[parentIdx(idx)], h.items[idx]
		idx = parentIdx(idx)
	}
}

func (h *MinHeap[T]) Peak() (T, error) {
	if len(h.items) == 0 {
		return h.zero, fmt.Errorf("heap is empty")
	}

	return h.items[0], nil
}

func (h *MinHeap[T]) Extract() (T, error) {
	minEl, err := h.Peak()

	if err != nil {
		return minEl, err
	}

	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]

	heapify(h, 0)

	return minEl, nil
}

func (h *MinHeap[T]) Len() int {
	return len(h.items)
}

func heapify[T any](heap *MinHeap[T], idx int) {
	left := leftChildIdx(idx)
	right := rightChildIdx(idx)
	minIdx := idx

	if left < len(heap.items) && heap.less(heap.items[left], heap.items[idx]) {
		minIdx = left
	}

	if right < len(heap.items) && heap.less(heap.items[right], heap.items[minIdx]) {
		minIdx = right
	}

	if minIdx != idx {
		heap.items[minIdx], heap.items[idx] = heap.items[idx], heap.items[minIdx]
		heapify(heap, minIdx)
	}
}

func leftChildIdx(idx int) int {
	return 2*idx + 1
}

func rightChildIdx(idx int) int {
	return 2*idx + 2
}

func parentIdx(childIdx int) int {
	return (childIdx - 1) / 2
}
