// Package plan specifies details regarding the game's map.
// "Plan" is used as synonym to map, as "map" is a reserved word.
package plan

import (
	"github.com/agstrc/heuristic-search/xy"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// DungeonCount is the number of dungeons in the game's map.
	DungeonCount = 3
	// The main map size in number of blocks.
	MainMapSize = 42
	// The dungeon size in number of blocks.
	DungeonSize = 28
)

// Plan is the game's map. It consists of a "main" map, which is all areas outside of a
// dungeon, and three dungeons, each with its own inner map.
type Plan struct {
	// Grid is the main map's terrain grid.
	Grid [][]Terrain

	// Sword defines the Master Sword's coordinates.
	Sword xy.XY
	// Gate defines the Lost Wood's gate coordinates.
	Gate xy.XY

	// Dungeons are all the dungeons in the game.
	Dungeons [DungeonCount]Dungeon
}

// Dungeon contains all data required to represent a dungeon within the main map.
type Dungeon struct {
	// Grid is the dungeon's inner grid. It reports whether a position is traversable.
	Grid [][]DungeonTerrain

	// Entrance is the dungeon's entrance in the main map.
	Entrance xy.XY

	// Start is the dungeon's inner start point. After "entering" the dungeon, the agent
	// should be at Start.
	Start xy.XY
	// GoalXY is the goal's position in the inner grid.
	GoalXY xy.XY
	// GoalImg is the image which will be used to visually represent the goal inside the
	// dungeon.
	GoalImg *ebiten.Image
}
