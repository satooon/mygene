// Generated by: gen
// TypeWriter: slice
// Directive: +gen on *Table

package schema

import (
	"errors"
	"math/rand"
)

// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.

// TableSlice is a slice of type *Table. Use it where you would use []*Table.
type TableSlice []*Table

// All verifies that all elements of TableSlice return true for the passed func. See: http://clipperhouse.github.io/gen/#All
func (rcv TableSlice) All(fn func(*Table) bool) bool {
	for _, v := range rcv {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Any verifies that one or more elements of TableSlice return true for the passed func. See: http://clipperhouse.github.io/gen/#Any
func (rcv TableSlice) Any(fn func(*Table) bool) bool {
	for _, v := range rcv {
		if fn(v) {
			return true
		}
	}
	return false
}

// Count gives the number elements of TableSlice that return true for the passed func. See: http://clipperhouse.github.io/gen/#Count
func (rcv TableSlice) Count(fn func(*Table) bool) (result int) {
	for _, v := range rcv {
		if fn(v) {
			result++
		}
	}
	return
}

// DistinctBy returns a new TableSlice whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
func (rcv TableSlice) DistinctBy(equal func(*Table, *Table) bool) (result TableSlice) {
Outer:
	for _, v := range rcv {
		for _, r := range result {
			if equal(v, r) {
				continue Outer
			}
		}
		result = append(result, v)
	}
	return result
}

// First returns the first element that returns true for the passed func. Returns error if no elements return true. See: http://clipperhouse.github.io/gen/#First
func (rcv TableSlice) First(fn func(*Table) bool) (result *Table, err error) {
	for _, v := range rcv {
		if fn(v) {
			result = v
			return
		}
	}
	err = errors.New("no TableSlice elements return true for passed func")
	return
}

// GroupByString groups elements into a map keyed by string. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv TableSlice) GroupByString(fn func(*Table) string) map[string]TableSlice {
	result := make(map[string]TableSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// Shuffle returns a shuffled copy of TableSlice, using a version of the Fisher-Yates shuffle. See: http://clipperhouse.github.io/gen/#Shuffle
func (rcv TableSlice) Shuffle() TableSlice {
	numItems := len(rcv)
	result := make(TableSlice, numItems)
	copy(result, rcv)
	for i := 0; i < numItems; i++ {
		r := i + rand.Intn(numItems-i)
		result[r], result[i] = result[i], result[r]
	}
	return result
}

// SortBy returns a new ordered TableSlice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv TableSlice) SortBy(less func(*Table, *Table) bool) TableSlice {
	result := make(TableSlice, len(rcv))
	copy(result, rcv)
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(result)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSortTableSlice(result, less, 0, n, maxDepth)
	return result
}

// Where returns a new TableSlice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv TableSlice) Where(fn func(*Table) bool) (result TableSlice) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Sort implementation based on http://golang.org/pkg/sort/#Sort, see top of this file

func swapTableSlice(rcv TableSlice, a, b int) {
	rcv[a], rcv[b] = rcv[b], rcv[a]
}

// Insertion sort
func insertionSortTableSlice(rcv TableSlice, less func(*Table, *Table) bool, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(rcv[j], rcv[j-1]); j-- {
			swapTableSlice(rcv, j, j-1)
		}
	}
}

// siftDown implements the heap property on rcv[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDownTableSlice(rcv TableSlice, less func(*Table, *Table) bool, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(rcv[first+child], rcv[first+child+1]) {
			child++
		}
		if !less(rcv[first+root], rcv[first+child]) {
			return
		}
		swapTableSlice(rcv, first+root, first+child)
		root = child
	}
}

func heapSortTableSlice(rcv TableSlice, less func(*Table, *Table) bool, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDownTableSlice(rcv, less, i, hi, first)
	}

	// Pop elements, largest first, into end of rcv.
	for i := hi - 1; i >= 0; i-- {
		swapTableSlice(rcv, first, first+i)
		siftDownTableSlice(rcv, less, lo, i, first)
	}
}

// Quicksort, following Bentley and McIlroy,
// Engineering a Sort Function, SP&E November 1993.

// medianOfThree moves the median of the three values rcv[a], rcv[b], rcv[c] into rcv[a].
func medianOfThreeTableSlice(rcv TableSlice, less func(*Table, *Table) bool, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if less(rcv[m1], rcv[m0]) {
		swapTableSlice(rcv, m1, m0)
	}
	if less(rcv[m2], rcv[m1]) {
		swapTableSlice(rcv, m2, m1)
	}
	if less(rcv[m1], rcv[m0]) {
		swapTableSlice(rcv, m1, m0)
	}
	// now rcv[m0] <= rcv[m1] <= rcv[m2]
}

func swapRangeTableSlice(rcv TableSlice, a, b, n int) {
	for i := 0; i < n; i++ {
		swapTableSlice(rcv, a+i, b+i)
	}
}

func doPivotTableSlice(rcv TableSlice, less func(*Table, *Table) bool, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's Ninther, median of three medians of three.
		s := (hi - lo) / 8
		medianOfThreeTableSlice(rcv, less, lo, lo+s, lo+2*s)
		medianOfThreeTableSlice(rcv, less, m, m-s, m+s)
		medianOfThreeTableSlice(rcv, less, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThreeTableSlice(rcv, less, lo, m, hi-1)

	// Invariants are:
	//	rcv[lo] = pivot (set up by ChoosePivot)
	//	rcv[lo <= i < a] = pivot
	//	rcv[a <= i < b] < pivot
	//	rcv[b <= i < c] is unexamined
	//	rcv[c <= i < d] > pivot
	//	rcv[d <= i < hi] = pivot
	//
	// Once b meets c, can swap the "= pivot" sections
	// into the middle of the slice.
	pivot := lo
	a, b, c, d := lo+1, lo+1, hi, hi
	for {
		for b < c {
			if less(rcv[b], rcv[pivot]) { // rcv[b] < pivot
				b++
			} else if !less(rcv[pivot], rcv[b]) { // rcv[b] = pivot
				swapTableSlice(rcv, a, b)
				a++
				b++
			} else {
				break
			}
		}
		for b < c {
			if less(rcv[pivot], rcv[c-1]) { // rcv[c-1] > pivot
				c--
			} else if !less(rcv[c-1], rcv[pivot]) { // rcv[c-1] = pivot
				swapTableSlice(rcv, c-1, d-1)
				c--
				d--
			} else {
				break
			}
		}
		if b >= c {
			break
		}
		// rcv[b] > pivot; rcv[c-1] < pivot
		swapTableSlice(rcv, b, c-1)
		b++
		c--
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := min(b-a, a-lo)
	swapRangeTableSlice(rcv, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRangeTableSlice(rcv, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

func quickSortTableSlice(rcv TableSlice, less func(*Table, *Table) bool, a, b, maxDepth int) {
	for b-a > 7 {
		if maxDepth == 0 {
			heapSortTableSlice(rcv, less, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivotTableSlice(rcv, less, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSortTableSlice(rcv, less, a, mlo, maxDepth)
			a = mhi // i.e., quickSortTableSlice(rcv, mhi, b)
		} else {
			quickSortTableSlice(rcv, less, mhi, b, maxDepth)
			b = mlo // i.e., quickSortTableSlice(rcv, a, mlo)
		}
	}
	if b-a > 1 {
		insertionSortTableSlice(rcv, less, a, b)
	}
}
