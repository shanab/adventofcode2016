package main

import (
	"fmt"
	"strconv"
	"strings"
)

const PROBLEM_STATEMENT = `
--- Day 1: No Time for a Taxicab ---

Santa's sleigh uses a very high-precision clock to guide its movements, and the clock's oscillator is regulated by stars. Unfortunately, the stars have been stolen... by the Easter Bunny. To save Christmas, Santa needs you to retrieve all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You're airdropped near Easter Bunny Headquarters in a city somewhere. "Near", unfortunately, is as close as you can get - the instructions on the Easter Bunny Recruiting Document the Elves intercepted start here, and nobody had time to work them out further.

The Document indicates that you should start at the given coordinates (where you just landed) and face North. Then, follow the provided sequence: either turn left (L) or right (R) 90 degrees, then walk forward the given number of blocks, ending at a new intersection.

There's no time to follow such ridiculous instructions on foot, though, so you take a moment and work out the destination. Given that you can only walk on the street grid of the city, how far is the shortest path to the destination?

For example:

Following R2, L3 leaves you 2 blocks East and 3 blocks North, or 5 blocks away.
R2, R2, R2 leaves you 2 blocks due South of your starting position, which is 2 blocks away.
R5, L5, R5, R3 leaves you 12 blocks away.
How many blocks away is Easter Bunny HQ?
`

const INPUT = `L4, L3, R1, L4, R2, R2, L1, L2, R1, R1, L3, R5, L2, R5, L4, L3, R2, R2, L5, L1, R4, L1, R3, L3, R5, R2, L5, R2, R1, R1, L5, R1, L3, L2, L5, R4, R4, L2, L1, L1, R1, R1, L185, R4, L1, L1, R5, R1, L1, L3, L2, L1, R2, R2, R2, L1, L1, R4, R5, R53, L1, R1, R78, R3, R4, L1, R5, L1, L4, R3, R3, L3, L3, R191, R4, R1, L4, L1, R3, L1, L2, R3, R2, R4, R5, R5, L3, L5, R2, R3, L1, L1, L3, R1, R4, R1, R3, R4, R4, R4, R5, R2, L5, R1, R2, R5, L3, L4, R1, L5, R1, L4, L3, R5, R5, L3, L4, L4, R2, R2, L5, R3, R1, R2, R5, L5, L3, R4, L5, R5, L3, R1, L1, R4, R4, L3, R2, R5, R1, R2, L1, R4, R1, L3, L3, L5, R2, R5, L1, L4, R3, R3, L3, R2, L5, R1, R3, L3, R2, L1, R4, R3, L4, R5, L2, L2, R5, R1, R2, L4, L4, L5, R3, L4`

const (
	NORTH int = iota
	EAST
	SOUTH
	WEST
)

type Position struct {
	direction  int
	coordinate *Coordinate
}

type Coordinate struct {
	x int
	y int
}

func NewPosition() *Position {
	return &Position{
		coordinate: &Coordinate{},
	}
}

func (p *Position) rotate(direction byte) {
	switch direction {
	case 'L':
		p.direction -= 1
		if p.direction < 0 {
			p.direction = 3
		}
	case 'R':
		p.direction += 1
		if p.direction > 3 {
			p.direction = 0
		}
	default:
		panic("Invalid turn!")
	}
}

func (p *Position) move(steps int) {
	switch p.direction {
	case NORTH:
		p.coordinate.y += steps
	case SOUTH:
		p.coordinate.y -= steps
	case EAST:
		p.coordinate.x += steps
	case WEST:
		p.coordinate.x -= steps
	}
}

func main() {
	instructions := strings.Split(INPUT, ", ")
	fmt.Println(BlocksAway(instructions))
}

func BlocksAway(instructions []string) int {
	position := NewPosition()
	coordinatesMap := make(map[Coordinate]bool)
OUTER:
	for _, instruction := range instructions {
		position.rotate(instruction[0])
		steps, err := strconv.Atoi(string(instruction[1:]))
		if err != nil {
			panic("Failed to convert steps to int")
		}

		for i := 0; i < steps; i++ {
			position.move(1)
			if coordinatesMap[*position.coordinate] {
				break OUTER
			}
			coordinatesMap[*position.coordinate] = true
		}
	}

	var result int
	if position.coordinate.x >= 0 {
		result += position.coordinate.x
	} else {
		result += -position.coordinate.x
	}

	if position.coordinate.y >= 0 {
		result += position.coordinate.y
	} else {
		result += -position.coordinate.y
	}

	return result
}
