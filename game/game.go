// Package game interacts with ebiten in order to provide an interface
package game

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/agstrc/heuristic-search/game/plan"
	"github.com/agstrc/heuristic-search/xy"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten's game interface.
type Game struct {
	u func() error
	d func(screen *ebiten.Image)

	plan *plan.Plan
	cost int
}

var _ ebiten.Game = &Game{}

func (game *Game) Update() error {
	return game.u()
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.d(screen)
	drawCostAt(screen, game.cost, xy.XY{X: 0, Y: 0})
}

func (game *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	const size = plan.MainMapSize * plan.TIS
	return size, size
}

// DefaultGame returns a default value of Game.
func DefaultGame() *Game {
	// defaultJSONPlan must be valid
	jplan := plan.DefaultJSONPlan()
	plan := jplan.ToPlan()
	game := Game{plan: &plan, cost: 0}

	mapCrawler := mapCrawler{
		game: &game, agent: jplan.Start,
		paths: nil, visited: make(map[xy.XY]struct{}),
		lastStep: time.Time{},
	}

	game.u = mapCrawler.update
	game.d = mapCrawler.draw

	return &game
}

// GameFromJSON instantiates a new game by settings its map according to a well formatted
// JSON file.
func GameFromJSON(path string) (*Game, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var jplan plan.JSONPlan

	if err := json.Unmarshal(fileData, &jplan); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON file: %w", err)
	}
	if err := jplan.Validate(); err != nil {
		return nil, fmt.Errorf("invalid JSON file: %w", err)
	}
	plan := jplan.ToPlan()
	game := Game{plan: &plan, cost: 0}

	mapCrawler := mapCrawler{
		game: &game, agent: jplan.Start,
		paths: nil, visited: make(map[xy.XY]struct{}),
		lastStep: time.Time{},
	}

	game.u = mapCrawler.update
	game.d = mapCrawler.draw

	return &game, nil
}
