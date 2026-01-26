package main

import (
	"slices"

	"github.com/google/uuid"
)

type Circuits struct {
	count    []string
	circuits map[string][]Coordinate
}

func NewCircuits() *Circuits {
	return &Circuits{
		count:    []string{},
		circuits: make(map[string][]Coordinate),
	}
}

func (c *Circuits) Add(set CoordinateSet) {
	circuitsA := c.filter(func(i string, v []Coordinate) bool {
		return slices.Contains(v, set.boxA)
	})
	circuitsB := c.filter(func(i string, v []Coordinate) bool {
		return slices.Contains(v, set.boxB)
	})

	switch {
	
	case len(circuitsA) == 1 && len(circuitsB) == 1:
		if circuitsA[0] == circuitsB[0] {
			break
		}
		// merge the two circuits
		c.merge(circuitsA[0], circuitsB[0])
		
	case len(circuitsA) == 1 && len(circuitsB) == 0:
		// join the circuit A
		val := append(c.circuits[circuitsA[0]], set.boxB)
		// val = slices.Compact(val)
		c.circuits[circuitsA[0]] = val
			
		
	case len(circuitsA) == 0 && len(circuitsB) == 1:
		// join the circuit B
		val := append(c.circuits[circuitsB[0]], set.boxA)
		// val = slices.Compact(val)
		c.circuits[circuitsB[0]] = val
		

	case len(circuitsA) == 0 && len(circuitsB) == 0:
		// create a new circuit
		id, err := uuid.NewRandom()
		if err != nil {
			checkErr(err)
		}
		c.circuits[id.String()] = append(c.circuits[id.String()], set.boxA, set.boxB)

		
		// > 1 -> shouldnt happen but error
	}
}

func (c *Circuits) merge(idA, idB string) {
	arrA := make([]Coordinate, len(c.circuits[idA]))
	arrB := make([]Coordinate, len(c.circuits[idB]))
	copy(arrA, c.circuits[idA])
	copy(arrB, c.circuits[idB])
	delete(c.circuits, idB)
	val := append(arrA, arrB...)	
	c.circuits[idA] = val
	
}

func (c *Circuits) filter(pred func(string, []Coordinate) bool) []string {
	result := make([]string, 0)

	for k, v := range c.circuits {
		if pred(k, v) {
			result = append(result, k)
		}
	}

	return result
}
