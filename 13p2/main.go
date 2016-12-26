package main

import (
	"container/list"
	"fmt"
	"strconv"
)

/*
--- Day 13: A Maze of Twisty Little Cubicles ---

You arrive at the first floor of this new building to discover a much less welcoming environment than the shiny atrium of the last one. Instead, you are in a maze of twisty little cubicles, all alike.

Every location in this area is addressed by a pair of non-negative integers (x,y). Each such coordinate is either a wall or an open space. You can't move diagonally. The cube maze starts at 0,0 and seems to extend infinitely toward positive x and y; negative values are invalid, as they represent a location outside the building. You are in a small waiting area at 1,1.

While it seems chaotic, a nearby morale-boosting poster explains, the layout is actually quite logical. You can determine whether a given x,y coordinate will be a wall or an open space using a simple system:

Find x*x + 3*x + 2*x*y + y + y*y.
Add the office designer's favorite number (your puzzle input).
Find the binary representation of that sum; count the number of bits that are 1.
If the number of bits that are 1 is even, it's an open space.
If the number of bits that are 1 is odd, it's a wall.
For example, if the office designer's favorite number were 10, drawing walls as # and open spaces as ., the corner of the building containing 0,0 would look like this:

  0123456789
0 .#.####.##
1 ..#..#...#
2 #....##...
3 ###.#.###.
4 .##..#..#.
5 ..##....#.
6 #...##.###
Now, suppose you wanted to reach 7,4. The shortest route you could take is marked as O:

  0123456789
0 .#.####.##
1 .O#..#...#
2 #OOO.##...
3 ###O#.###.
4 .##OO#OO#.
5 ..##OOO.#.
6 #...##.###
Thus, reaching 7,4 would take a minimum of 11 steps (starting from your current location, 1,1).

What is the fewest number of steps required for you to reach 31,39?
*/

const (
	INPUT   = 1350
	ROWS    = 50
	COLUMNS = 50
	STEPS   = 50
)

var grid [][]bool

type Coordinate struct {
	row int
	col int
}

func NewCoordinate(row, col int) Coordinate {
	return Coordinate{row, col}
}

type Node struct {
	Coordinate
	steps int
}

func NewNode(row, col, steps int) Node {
	return Node{
		NewCoordinate(row, col),
		steps,
	}
}

func (n Node) Children() []Node {
	children := make([]Node, 0)

	// Top
	if n.row-1 >= 0 && grid[n.row-1][n.col] {
		children = append(children, NewNode(n.row-1, n.col, n.steps+1))
	}
	// Bottom
	if n.row+1 < len(grid) && grid[n.row+1][n.col] {
		children = append(children, NewNode(n.row+1, n.col, n.steps+1))
	}
	// Left
	if n.col-1 >= 0 && grid[n.row][n.col-1] {
		children = append(children, NewNode(n.row, n.col-1, n.steps+1))
	}
	// Right
	if n.col+1 < len(grid[0]) && grid[n.row][n.col+1] {
		children = append(children, NewNode(n.row, n.col+1, n.steps+1))
	}
	return children
}

func main() {
	grid = MakeGrid(ROWS, COLUMNS)
	fmt.Println(LocationsCount(NewNode(1, 1, 0), STEPS))
}

func MakeGrid(rows, columns int) [][]bool {
	g := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		g[i] = make([]bool, columns)
		for j := 0; j < columns; j++ {
			num := j*j + 3*j + 2*i*j + i + i*i + INPUT
			bin := strconv.FormatInt(int64(num), 2)
			var oneCount int
			for i := 0; i < len(bin); i++ {
				if bin[i] == '1' {
					oneCount++
				}
			}

			if oneCount%2 == 0 {
				g[i][j] = true
			} else {
				g[i][j] = false
			}
		}
	}
	return g
}

func LocationsCount(src Node, steps int) int {
	queue := list.New()
	visited := make(map[Coordinate]bool)

	queue.PushBack(src)
	visited[src.Coordinate] = true
	visitedCount := 1

	for queue.Len() > 0 {
		elem := queue.Front()
		node := queue.Remove(elem).(Node)
		if !visited[node.Coordinate] {
			visited[node.Coordinate] = true
			visitedCount++
		}

		if node.steps < STEPS {
			for _, n := range node.Children() {
				if !visited[n.Coordinate] {
					queue.PushBack(n)
				}
			}
		}

	}

	return visitedCount
}
