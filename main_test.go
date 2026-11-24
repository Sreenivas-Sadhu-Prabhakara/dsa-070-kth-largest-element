package main

import "testing"

func TestKthLargest(t *testing.T) {
	cases := []struct {
		in   []int
		k    int
		want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{[]int{1}, 1, 1},
		{[]int{7, 7, 7}, 2, 7},
		{[]int{-1, -5, -3}, 1, -1},
	}
	for _, c := range cases {
		got, ok := KthLargest(c.in, c.k)
		if !ok || got != c.want {
			t.Errorf("KthLargest(%v,%d) = (%d,%v), want %d", c.in, c.k, got, ok, c.want)
		}
	}
	if _, ok := KthLargest([]int{1, 2}, 3); ok {
		t.Error("k beyond length should be not-ok")
	}
	if _, ok := KthLargest([]int{}, 1); ok {
		t.Error("empty input should be not-ok")
	}
}
