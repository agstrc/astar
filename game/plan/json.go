package plan

import (
	"fmt"

	"github.com/agstrc/heuristic-search/images"
	"github.com/agstrc/heuristic-search/xy"
	"github.com/hajimehoshi/ebiten/v2"
)

// JSONPlan represents the game's map in a JSON format.
type JSONPlan struct {
	MasterSword xy.XY `json:"master_sword"`
	LostWoods   xy.XY `json:"lost_woods"`
	Start       xy.XY `json:"start"`

	MainMap  []string      `json:"main_map"`
	Dungeons []JSONDungeon `json:"dungeons"`
}

func (jp JSONPlan) ToPlan() Plan {
	plan := Plan{Sword: jp.MasterSword, Gate: jp.LostWoods}
	grid := [][]Terrain{}

	tm := jp.gridMap()
	for idx, str := range jp.MainMap {
		grid = append(grid, make([]Terrain, 0, MainMapSize))

		for _, rune := range str {
			grid[idx] = append(grid[idx], tm[rune])
		}
	}
	plan.Grid = grid

	dtm := jp.Dungeons[0].gridMap()
	images := [...]*ebiten.Image{images.Virtue1, images.Virtue2, images.Virtue3}
	for idx, jsonDungeon := range jp.Dungeons {
		dungeon := Dungeon{
			Entrance: jsonDungeon.Entrance,
			Start:    jsonDungeon.Start,
			GoalXY:   jsonDungeon.Goal,
			GoalImg:  images[idx],
		}

		for idx, str := range jsonDungeon.Grid {
			dungeon.Grid = append(dungeon.Grid, make([]DungeonTerrain, 0, DungeonSize))

			for _, rune := range str {
				dungeon.Grid[idx] = append(dungeon.Grid[idx], dtm[rune])
			}

		}

		plan.Dungeons[idx] = dungeon
	}

	return plan
}

func (jp JSONPlan) Validate() error {
	if err := jp.validateMainMap(); err != nil {
		return fmt.Errorf("main map is invalid: %w", err)
	}
	if err := jp.validateCoordinates(); err != nil {
		return fmt.Errorf("coordinates are invalid: %w", err)
	}
	if err := jp.validateDungeons(); err != nil {
		return fmt.Errorf("dungeons are invalid: %w", err)
	}

	return nil
}

func (jp JSONPlan) validateCoordinates() error {
	for _, coord := range [...]xy.XY{jp.MasterSword, jp.LostWoods, jp.Start} {
		if coord.X < 0 || coord.X >= MainMapSize {
			return fmt.Errorf("invalid coordinate pair: (%d, %d)", coord.X, coord.Y)
		}
		if coord.Y < 0 || coord.Y >= MainMapSize {
			return fmt.Errorf("invalid coordinate pair: (%d, %d)", coord.X, coord.Y)
		}
	}
	return nil
}

func (jp JSONPlan) validateMainMap() error {
	if len(jp.MainMap) != MainMapSize {
		return fmt.Errorf("invalid main map array length: %d", len(jp.MainMap))
	}

	tm := jp.gridMap()

	for idx, str := range jp.MainMap {
		if len(str) != MainMapSize {
			return fmt.Errorf("row (index %d) has invalid length: %d", idx, len(str))
		}

		for _, rune := range str {
			if _, inMap := tm[rune]; !inMap {
				return fmt.Errorf("row (index %d) has unknown character: %s", idx, string(rune))
			}
		}
	}

	return nil
}

func (jp JSONPlan) validateDungeons() error {
	if len(jp.Dungeons) != DungeonCount {
		return fmt.Errorf("invalid amount of dungeons (have %d, want %d)", len(jp.Dungeons), DungeonCount)
	}

	for idx, dungeon := range jp.Dungeons {
		if err := dungeon.validate(); err != nil {
			return fmt.Errorf("dungeon (index %d) is invalid: %w", idx, err)
		}
	}
	return nil
}

func (JSONPlan) gridMap() map[rune]Terrain {
	return map[rune]Terrain{
		'@': Forest, ' ': Grass,
		'%': Mountain, '_': Sand,
		'*': Water,
	}
}

// JSONDungeon represents the game's dungeon in a JSON format.
type JSONDungeon struct {
	Grid []string `json:"grid"`

	Entrance xy.XY `json:"entrance"`
	Start    xy.XY `json:"start"`
	Goal     xy.XY `json:"goal"`
}

func (jd JSONDungeon) validate() error {
	if err := jd.validateGrid(); err != nil {
		return fmt.Errorf("grid is invalid: %w", err)
	}
	if err := jd.validateCoordinates(); err != nil {
		return fmt.Errorf("coordinates are invalid: %w", err)
	}
	return nil
}

func (jd JSONDungeon) validateCoordinates() error {
	for _, coord := range [...]xy.XY{jd.Start, jd.Goal} {
		if coord.X < 0 || coord.X >= DungeonSize {
			return fmt.Errorf("invalid coordinate pair: (%d, %d)", coord.X, coord.Y)
		}
		if coord.Y < 0 || coord.Y >= DungeonSize {
			return fmt.Errorf("invalid coordinate pair: (%d, %d)", coord.X, coord.Y)
		}

		if jd.Grid[coord.Y][coord.X] == '#' {
			return fmt.Errorf("coordinate pair (%d, %d) is on non traversable block", coord.X, coord.Y)
		}
	}

	if jd.Entrance.X < 0 || jd.Entrance.X >= MainMapSize {
		return fmt.Errorf("invalid entrance coordinates: (%d, %d)", jd.Entrance.X, jd.Entrance.Y)
	}
	if jd.Entrance.Y < 0 || jd.Entrance.Y >= MainMapSize {
		return fmt.Errorf("invalid entrance coordinates: (%d, %d)", jd.Entrance.X, jd.Entrance.Y)
	}

	return nil
}

func (jd JSONDungeon) validateGrid() error {
	if len(jd.Grid) != DungeonSize {
		return fmt.Errorf("invalid dungeon map array length: %d", len(jd.Grid))
	}

	tm := jd.gridMap()

	for idx, str := range jd.Grid {
		if len(str) != DungeonSize {
			return fmt.Errorf("row (index %d) has invalid length: %d", idx, len(str))
		}

		for _, rune := range str {
			if _, inMap := tm[rune]; !inMap {
				return fmt.Errorf("row (index %d) has unknown character: %v", idx, rune)
			}
		}
	}

	return nil
}

func (JSONDungeon) gridMap() map[rune]DungeonTerrain {
	return map[rune]DungeonTerrain{
		'#': NonTraversable, ' ': Traversable,
	}
}
