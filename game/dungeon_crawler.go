package game

import (
	"time"

	"github.com/agstrc/heuristic-search/astar"
	"github.com/agstrc/heuristic-search/game/plan"
	"github.com/agstrc/heuristic-search/images"
	"github.com/agstrc/heuristic-search/xy"
	"github.com/hajimehoshi/ebiten/v2"
)

// dungeonCrawler defines the game's state while the game agent is inside a dungeon.
type dungeonCrawler struct {
	game       *Game
	mapCrawler *mapCrawler

	dungeon plan.Dungeon
	agent   xy.XY
	path    []dtNode

	objectiveReached bool

	lastStep time.Time
}

func (dc *dungeonCrawler) initPath() {
	grid := dc.dungeon.Grid
	start := dtNode{XY: dc.dungeon.Start, grid: &grid}
	goal := dtNode{XY: dc.dungeon.GoalXY, grid: &grid}

	path, _ := astar.FindPath(start, goal, dtHeuristic)
	popped, path := path[0], path[1:]

	// appends the path from the goal back to the start, in order to exit the dungeon
	for i := len(path) - 2; i >= 0; i-- {
		path = append(path, path[i])
	}
	path = append(path, popped)

	dc.path = path
}

func (dc *dungeonCrawler) update() error {
	if dc.path == nil {
		dc.initPath()
	}
	if time.Since(dc.lastStep) < stepInterval {
		return nil
	}

	if dc.agent == dc.dungeon.GoalXY {
		dc.objectiveReached = true
	}
	if !dc.step() {
		dc.returnToMap()
		return nil
	}

	return nil
}

// returnToMap changes the game state back to the map crawling state in which it was
// before entering a dungeon.
func (dc *dungeonCrawler) returnToMap() {
	game := dc.game
	crawler := dc.mapCrawler
	game.u = crawler.update
	game.d = crawler.draw
}

// step moves the agent in accordance to the next movement in the crawler's path. If the
// path is over, it returns false.
func (dc *dungeonCrawler) step() bool {
	if len(dc.path) == 0 {
		return false
	}

	dc.agent = dc.path[0].XY
	dc.increaseCost()
	dc.path = dc.path[1:]
	dc.lastStep = time.Now()

	return true
}

func (dc *dungeonCrawler) increaseCost() {
	game := dc.game
	game.cost += dc.dungeon.Grid[dc.agent.Y][dc.agent.X].Cost()
}

func (dc *dungeonCrawler) draw(screen *ebiten.Image) {
	dc.drawDungeonGrid(screen)
	dc.drawDetails(screen)
	dc.drawAgent(screen)
}

// offsetOptsCenter applies a translation to opts that moves it further towards the
// center of the screen.
func offsetOptsCenter(opts *ebiten.DrawImageOptions) {
	// move calculates the necessary offset in order to center the images drawn by the
	// dungeon.
	const move float64 = (plan.DungeonSize / 4.0) * plan.TIS
	opts.GeoM.Translate(move, move)
}

// drawDungeonGrid draws the dungeon's blocks.
func (dc *dungeonCrawler) drawDungeonGrid(screen *ebiten.Image) {
	grid := dc.dungeon.Grid
	for y := range grid {
		for x := range grid {
			drawX, drawY := (x * plan.TIS), y*plan.TIS
			drawOpts := ebiten.DrawImageOptions{}
			offsetOptsCenter(&drawOpts)
			drawOpts.GeoM.Translate(float64(drawX), float64(drawY))

			screen.DrawImage(grid[y][x].Image(), &drawOpts)
		}
	}
}

// drawDetails draws the dungeon's goal and start point.
func (dc *dungeonCrawler) drawDetails(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	offsetOptsCenter(&opts)
	dungeonX, dungeonY := dc.dungeon.Start.X, dc.dungeon.Start.Y
	dungeonX, dungeonY = dungeonX*plan.TIS, dungeonY*plan.TIS
	opts.GeoM.Translate(float64(dungeonX), float64(dungeonY))
	screen.DrawImage(images.Dungeon, &opts)

	if !dc.objectiveReached {
		opts = ebiten.DrawImageOptions{} // reset opts
		offsetOptsCenter(&opts)
		goalX, goalY := dc.dungeon.GoalXY.X, dc.dungeon.GoalXY.Y
		goalX, goalY = goalX*plan.TIS, goalY*plan.TIS
		opts.GeoM.Translate(float64(goalX), float64(goalY))
		screen.DrawImage(dc.dungeon.GoalImg, &opts)
	}
}

// drawAgent draws the agent with the required offset.
func (dc *dungeonCrawler) drawAgent(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	offsetOptsCenter(&opts)
	agentX, agentY := dc.agent.X*plan.TIS, dc.agent.Y*plan.TIS

	// the agent is taller than a single block. Therefore, he is offset in order to have
	// its feet at the start of the correct block.
	agent := images.Agent
	_, agentHeight := agent.Size()
	agentY = agentY - (agentHeight - plan.TIS)
	opts.GeoM.Translate(float64(agentX), float64(agentY))
	screen.DrawImage(agent, &opts)
}
