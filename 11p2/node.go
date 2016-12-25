package main

import (
	"fmt"
)

type Node struct {
	index  int
	depth  int
	floors [][]byte
}

func NewNode(index, depth int, floors [][]byte) *Node {
	return &Node{
		index:  index,
		depth:  depth,
		floors: floors,
	}
}

func (n Node) String() string {
	return fmt.Sprintf(
		"%d-%s-%s-%s-%s",
		n.index,
		n.floors[0],
		n.floors[1],
		n.floors[2],
		n.floors[3])
}

func (n Node) nextFloors() []int {
	switch n.index {
	case 0:
		return []int{1}
	case 3:
		return []int{2}
	default:
		return []int{n.index - 1, n.index + 1}
	}
}

func (n Node) steps() [][]int {
	indices := make([]int, 0)
	for i, c := range n.floors[n.index] {
		if c == '1' {
			indices = append(indices, i)
		}
	}

	steps := make([][]int, 0)
	// Add 1 move steps
	for _, index := range indices {
		steps = append(steps, []int{index})
	}

	// Add 2 move steps
	for i := 0; i < len(indices)-1; i++ {
		for j := i + 1; j < len(indices); j++ {
			steps = append(steps, []int{indices[i], indices[j]})
		}
	}

	return steps
}

func (n Node) nextNode(floor int, step []int) *Node {
	newFloors := make([][]byte, 4)
	for i := 0; i < 4; i++ {
		newFloors[i] = make([]byte, len(n.floors[0]))
		copy(newFloors[i], n.floors[i])
	}

	for _, s := range step {
		newFloors[n.index][s] = '0'
		newFloors[floor][s] = '1'
	}

	if validFloor(newFloors[n.index]) && validFloor(newFloors[floor]) {
		return NewNode(floor, n.depth+1, newFloors)
	}

	return nil
}

func validFloor(floor []byte) bool {
	middle := len(floor) / 2

	noGenerators := true
	for i := middle; i < len(floor); i++ {
		if floor[i] == '1' {
			noGenerators = false
		}
	}

	if noGenerators {
		return true
	}

	for i := 0; i < len(floor)/2; i++ {
		if floor[i] == '1' && floor[i+middle] == '0' {
			return false
		}
	}

	return true
}

func (n Node) Children() []*Node {
	children := make([]*Node, 0)

	for _, f := range n.nextFloors() {
		for _, step := range n.steps() {
			if nextNode := n.nextNode(f, step); nextNode != nil {
				children = append(children, nextNode)
			}
		}
	}

	return children
}
