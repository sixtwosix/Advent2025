package heap

import "cmp"

// root -> idx = 0
// parent -> idx = (i-1)/2 round down
// current -> idx = i
// left child -> idx = 2*i+1
// right child -> idx = 2*i+2
// last non leaf -> idx = (len(array)-2)/2

type MaxHeap[T cmp.Ordered] []T

// heapify an array
func (h *MaxHeap[T]) BuildHeap(array []T) {
	lastNondLeafNodeIdx := (len(array) - 2) / 2
	for currentIdx := lastNondLeafNodeIdx; currentIdx >= 0; currentIdx-- {
		endIdx := len(array) - 1
		h.siftDown(currentIdx, endIdx)
	}
}

//Remove and return the maximum value and update heap ordering
func (h *MaxHeap[T]) Remove() T {
	var zero T
	n := len(*h)
	// if the heap is empty
	if n == 0 {
		return zero
	}
	// swap first and last element in the array
	h.swap(0, n-1)
	valueToRemove := (*h)[n-1]
	// pop the last element in the array
	*h = (*h)[:n-1]
	// call siftDown to update heap ordering
	h.siftDown(0, n-2)
	return valueToRemove
}

// Insert a value to the end of the tree and update the heap ordering
func (h *MaxHeap[T]) Insert(value T) {
	(*h) = append((*h), value)
	h.siftUp()
}

// Siftdown
// Continously swap the current node with its big child node untill its in the correct position
func (h *MaxHeap[T]) siftDown(currentIdx, endIdx int) {
	leftChildIdx := currentIdx * 2 + 1
	for leftChildIdx <= endIdx {
		rightChildIdx := currentIdx * 2 + 2
		if rightChildIdx > endIdx {
			rightChildIdx = -1
		}

		// get the bigger node to swap
		idxToSwap := leftChildIdx
		if rightChildIdx != -1 && (*h)[rightChildIdx] > (*h)[leftChildIdx] {
			idxToSwap = rightChildIdx
		}

		// check if the value of swap node is more than the value at currentIdx
		if (*h)[idxToSwap] > (*h)[currentIdx] {
			h.swap(idxToSwap, currentIdx)
			currentIdx = idxToSwap
			leftChildIdx = currentIdx * 2 + 1
		} else {
			return
		}
	}
}

// SiftUp
// Continously swap the current node with its parent node untill it is in the correct position
func (h *MaxHeap[T]) siftUp() {
	currentIdx := len(*h) - 1
	parentIdx := (currentIdx - 1)/2
	for currentIdx > 0 && (*h)[currentIdx] > (*h)[parentIdx] {
		h.swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = (currentIdx - 1)/2
	}
}

// Swap 
// swap the two items in the heap
func (h MaxHeap[T]) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}