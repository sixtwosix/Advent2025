package main

import (
	"slices"
)

type Circuits struct {
	count    int
	circuits map[int][]Coordinate
}

func NewCircuits() *Circuits {
	return &Circuits{
		count:    0,
		circuits: make(map[int][]Coordinate),
	}
}

func (c *Circuits) Add(set CoordinateSet) {
	circuitsA := c.filter(func(i int, v []Coordinate) bool {
		return slices.Contains(v, set.boxA)
	})
	circuitsB := c.filter(func(i int, v []Coordinate) bool {
		return slices.Contains(v, set.boxB)
	})

	switch {
	// merge the two circuits
	case len(circuitsA) == 1 && len(circuitsB) == 1:
	
	// join the circuit A
	case len(circuitsA) == 1 && len(circuitsB) == 0:
	// join the circuit B
	case len(circuitsA) == 0 && len(circuitsB) == 0:
	// create a new circuit
	// > 1 -> shouldnt happen but error
	}
}

func (c *Circuits) filter(pred func(int, []Coordinate) bool) []int {
	result := make([]int, 0)

	for k, v := range c.circuits {
		if pred(k, v) {
			result = append(result, k)
		}
	}

	return result
}
