package heaps

import "github.com/eugene-shcherbo/cs-problems/ds"

type MedianStream struct {
	left  *ds.MinHeap[int]
	right *ds.MinHeap[int]
}

func NewMedianStream() *MedianStream {
	return &MedianStream{ds.NewHeapComp([]int{}), ds.NewHeapComp([]int{})}
}

func (s *MedianStream) Insert(num int) {
	s.right.Add(num)

	if s.right.Len() > s.left.Len() {
		r, _ := s.right.Extract()
		s.left.Add(-r)
	}

	if s.right.Len() == 0 {
		return
	}

	l, _ := s.left.Peak()
	r, _ := s.right.Peak()

	if -l > r {
		s.left.Extract()
		s.right.Extract()
		s.left.Add(-r)
		s.right.Add(-l)
	}
}

func (s *MedianStream) FindMedian() float64 {
	l, _ := s.left.Peak()

	if s.left.Len() == s.right.Len() {
		r, _ := s.right.Peak()

		return float64(-l+r) / 2.0
	}

	return float64(l) / 1.0
}
