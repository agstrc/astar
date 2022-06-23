package astar

import (
	"reflect"
	"testing"
)

func TestFindPath(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		grid := [][]int{
			{0, 1},
			{3, 1},
			{0, 1},
			{0, 1},
			{0, 1},
		}

		start := intNode{x: 0, y: 0, grid: &grid}
		goal := intNode{x: 1, y: 4, grid: &grid}

		// "empty" heuristic.
		path, cost := FindPath(start, goal, func(_, _ intNode) int { return 0 })
		if path == nil {
			t.Fatal("Expected non nil path")
		}
		if cost != 4 {
			t.Fatal("Cost differs from expected")
		}

		if !reflect.DeepEqual(
			path,
			[]intNode{{0, 0, &grid}, {0, 1, &grid}, {0, 2, &grid}, {0, 3, &grid}, {0, 4, &grid}, {1, 4, &grid}},
		) {
			t.Fatal("Returned path differs from expected")
		}
	})

	t.Run("bool", func(t *testing.T) {
		grid := [][]bool{
			{true, false, false},
			{true, true, true},
			{false, false, false},
			{true, true, true},
		}

		start := boolNode{x: 1, y: 1, grid: &grid}
		goal := boolNode{x: 0, y: 3, grid: &grid}

		// "empty" heuristic
		path, _ := FindPath(start, goal, func(_, _ boolNode) int { return 0 })
		if path != nil {
			t.Fatal("Expected nil path")
		}

		goal = boolNode{x: 0, y: 0, grid: &grid}
		path, cost := FindPath(start, goal, func(_, _ boolNode) int { return 0 })
		if path == nil {
			t.Fatal("Expected non nil path")
		}
		if cost != 2 {
			t.Fatal("Cost differs from expected")
		}
		if !reflect.DeepEqual(
			path,
			[]boolNode{start, {x: 0, y: 1, grid: &grid}, goal},
		) {
			t.Fatal("Returned path differs from expected")
		}
	})
}

// ======================================================================================
// interface implementations

// intNode implements Node on a 2D slice of ints.
type intNode struct {
	x, y int
	// a pointer is used in order to make the type comparable
	grid *[][]int
}

func (i intNode) Neighbors() []intNode {
	neighborSlice := make([]intNode, 0, 4)
	rowCount := len(*i.grid)
	columnCount := len((*i.grid)[i.y])

	if i.x > 0 {
		duplicate := i
		duplicate.x--
		neighborSlice = append(neighborSlice, duplicate)
	}
	if i.x < columnCount-1 {
		duplicate := i
		duplicate.x++
		neighborSlice = append(neighborSlice, duplicate)
	}
	if i.y > 0 {
		duplicate := i
		duplicate.y--
		neighborSlice = append(neighborSlice, duplicate)
	}
	if i.y < rowCount-1 {
		duplicate := i
		duplicate.y++
		neighborSlice = append(neighborSlice, duplicate)
	}

	return neighborSlice
}

func (i intNode) Cost() int {
	return (*i.grid)[i.y][i.x]
}

// boolNode implements Node on a 2D slice of bools. The costs are constant and false
// values are not returned as neighbors.
type boolNode struct {
	x, y int
	grid *[][]bool
}

// The traversal cost is constant.
func (b boolNode) Cost() int {
	return 1
}

// Returned neighbors are always true values.
func (b boolNode) Neighbors() []boolNode {
	neighborSlice := make([]boolNode, 0, 4)
	rowCount := len(*b.grid)
	columnCount := len((*b.grid)[b.y])

	if b.x > 0 {
		duplicate := b
		duplicate.x--
		if (*b.grid)[b.y][b.x] {
			neighborSlice = append(neighborSlice, duplicate)
		}
	}
	if b.x < columnCount-1 {
		duplicate := b
		duplicate.x++
		if (*b.grid)[b.y][b.x] {
			neighborSlice = append(neighborSlice, duplicate)
		}
	}
	if b.y > 0 {
		duplicate := b
		duplicate.y--
		if (*b.grid)[b.y][b.x] {
			neighborSlice = append(neighborSlice, duplicate)
		}
	}
	if b.y < rowCount-1 {
		duplicate := b
		duplicate.y++
		if (*b.grid)[b.y][b.x] {
			neighborSlice = append(neighborSlice, duplicate)
		}
	}

	return neighborSlice
}
