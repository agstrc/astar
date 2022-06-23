package plan

import (
	"github.com/agstrc/heuristic-search/images"
	"github.com/hajimehoshi/ebiten/v2"
)

// TIS (terrain image size) is the expected length and width for every terrain image.
const TIS = 32

// Terrain is a type of terrain which may be traversed in the main map. Each terrain has
// an image and a traversal cost.
type Terrain struct {
	cost  int
	image *ebiten.Image
}

func (t Terrain) Cost() int {
	return t.cost
}

func (t Terrain) Image() *ebiten.Image {
	return t.image
}

// Predefined terrain types.
var (
	Forest   = Terrain{cost: 100, image: images.Forest}
	Grass    = Terrain{cost: 10, image: images.Grass}
	Mountain = Terrain{cost: 150, image: images.Mountain}
	Sand     = Terrain{cost: 20, image: images.Sand}
	Water    = Terrain{cost: 180, image: images.Water}
)

// DungeonTerrain is a terrain found in a dungeon. As dungeons are composed of either
// traversable or non traversable terrains, it may only be one of two values. Its inner
// type is a bool, which defines wheter its a traversable or non traversable terrain.
type DungeonTerrain bool

const (
	Traversable    DungeonTerrain = true
	NonTraversable DungeonTerrain = false
)

// Cost returns the terrain's cost. It is a constant value and only valid for traversable
// terrains. If called on a non traversable terrain, Cost panics.
func (dt DungeonTerrain) Cost() int {
	if !dt {
		panic("attempt to get cost of non traversable block")
	}
	return 10
}

func (dt DungeonTerrain) Image() *ebiten.Image {
	if dt {
		return images.Traversable
	} else {
		return images.NonTraversable
	}
}
