package game

import (
	"github.com/agstrc/heuristic-search/game/plan"
	"github.com/agstrc/heuristic-search/xy"
)

// This file exports implementations of astar.Node on top of the game's terrains.

// tNode implements Node on a position within a terrain grid. The grid is defined through
// a pointer, so nodes used in the same context must point to the same address.
type tNode struct {
	xy.XY
	// grid uses a pointer in order to make this struct comparable.
	grid *[][]plan.Terrain
}

func (tn tNode) Neighbors() []tNode {
	neighbors := make([]tNode, 0, 4)
	rowCount := len(*tn.grid)
	colCount := len((*tn.grid)[tn.Y])

	if tn.X > 0 {
		dup := tn
		dup.X--
		neighbors = append(neighbors, dup)
	}
	if tn.X < colCount-1 {
		dup := tn
		dup.X++
		neighbors = append(neighbors, dup)
	}
	if tn.Y > 0 {
		dup := tn
		dup.Y--
		neighbors = append(neighbors, dup)
	}
	if tn.Y < rowCount-1 {
		dup := tn
		dup.Y++
		neighbors = append(neighbors, dup)
	}

	return neighbors
}

func (tn tNode) Cost() int {
	return (*tn.grid)[tn.Y][tn.X].Cost()
}

// dtNode implements Node on a position within a dungeon terrain grid. The grid is
// defined through a pointer, so nodes used in the same context must point to the same
// address.
type dtNode struct {
	xy.XY
	grid *[][]plan.DungeonTerrain
}

// Neighbors returns the node's neighbors. Non traversable nodes are not connected to any
// nodes, therefore Neighbors only returns traversable nodes.
func (dt dtNode) Neighbors() []dtNode {
	neighbors := make([]dtNode, 0, 4)
	rowCount := len(*dt.grid)
	colCount := len((*dt.grid)[dt.Y])

	if dt.X > 0 {
		duplicate := dt
		duplicate.X--
		if (*duplicate.grid)[duplicate.Y][duplicate.X] {
			neighbors = append(neighbors, duplicate)
		}
	}
	if dt.X < colCount-1 {
		duplicate := dt
		duplicate.X++
		if (*duplicate.grid)[duplicate.Y][duplicate.X] {
			neighbors = append(neighbors, duplicate)
		}
	}
	if dt.Y > 0 {
		duplicate := dt
		duplicate.Y--
		if (*duplicate.grid)[duplicate.Y][duplicate.X] {
			neighbors = append(neighbors, duplicate)
		}
	}
	if dt.Y < rowCount-1 {
		duplicate := dt
		duplicate.Y++
		if (*duplicate.grid)[duplicate.Y][duplicate.X] {
			neighbors = append(neighbors, duplicate)
		}
	}

	return neighbors
}

func (dt dtNode) Cost() int {
	return (*dt.grid)[dt.Y][dt.X].Cost()
}

// tHeuristic implements a heuristic on a pair of TNodes which may be used on the A*
// algorithm.
func tHeuristic(from, to tNode) int {
	xDistance := from.X - to.X
	yDistance := from.Y - to.Y
	if xDistance < 0 {
		xDistance = -xDistance
	}
	if yDistance < 0 {
		yDistance = -yDistance
	}
	return xDistance + yDistance
}

// dtHeuristic implements a heuristic on a pair of DTNodes which may be used on the A*
// algorithm.
func dtHeuristic(from, to dtNode) int {
	xDistance := from.X - to.X
	yDistance := from.Y - to.Y
	if xDistance < 0 {
		xDistance = -xDistance
	}
	if yDistance < 0 {
		yDistance = -yDistance
	}
	return xDistance + yDistance
}
