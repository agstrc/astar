package game

import (
	"fmt"
	"time"

	"github.com/agstrc/heuristic-search/astar"
	"github.com/agstrc/heuristic-search/game/plan"
	"github.com/agstrc/heuristic-search/images"
	"github.com/agstrc/heuristic-search/xy"
	"github.com/hajimehoshi/ebiten/v2"
)

// stepInterval defines the waiting time between the agent's position updates.
const stepInterval = time.Millisecond * 100

// mapCrawler defines the game's state during the search for the all of the dungeons'
// goals.
type mapCrawler struct {
	game  *Game
	agent xy.XY
	paths [][]tNode

	// visited is a set that indicates if a dungeon has been visited. It can be checked
	// through the dungeon's coordinates
	visited map[xy.XY]struct{}
	// lastStep is the moment when the agent took its last step.
	lastStep time.Time
}

func (mc *mapCrawler) update() error {
	if mc.paths == nil {
		mc.initPaths()
	}
	if time.Since(mc.lastStep) < stepInterval {
		return nil
	}

	if len(mc.paths) == 0 {
		// path is done; nothing to do
		return nil
	}
	if !mc.step() {
		mc.nextPath()
		if len(mc.visited) == plan.DungeonCount {
			return nil
		}

		dungeon, err := dungeonByXY(mc.agent, mc.game.plan.Dungeons[:])
		if err != nil {
			panic(err)
		}
		mc.visited[dungeon.Entrance] = struct{}{}

		mc.transitionToDungeon(dungeon)
	}

	return nil
}

// transitionToDungeon changes the game's state to the dungeon state, in which the agent
// explores a dungeon in order to reach its goal.
func (mc *mapCrawler) transitionToDungeon(dungeon plan.Dungeon) {
	game := mc.game
	crawler := dungeonCrawler{
		game: game, mapCrawler: mc,
		dungeon: dungeon, agent: dungeon.Start,
		path: nil, objectiveReached: false, lastStep: time.Time{},
	}

	game.u = crawler.update
	game.d = crawler.draw
}

// step moves the agent one step further along the current path. If the path is over,
// it returns false.
func (mc *mapCrawler) step() bool {
	if len(mc.paths[0]) == 0 {
		return false
	}
	mc.agent = mc.paths[0][0].XY
	mc.increaseCost()
	mc.paths[0] = mc.paths[0][1:]
	mc.lastStep = time.Now()

	return true
}

func (mc *mapCrawler) increaseCost() {
	game := mc.game
	game.cost += game.plan.Grid[mc.agent.Y][mc.agent.X].Cost()
}

// nextPath changes the inner path being followed to the next one. If there are no more
// paths to follow, it returns false.
func (mc *mapCrawler) nextPath() bool {
	if len(mc.paths) == 0 {
		return false
	}
	mc.paths = mc.paths[1:]

	return true
}

// initPaths initiates the crawler's paths field. It sets it to the best route to reach
// the three dungeons and, after going through all of the dungeons, heading back to the
// starting point and then to the Lost Woods' gate.
func (mc *mapCrawler) initPaths() {
	grid := mc.game.plan.Grid
	paths, cost := [][]tNode(nil), 0
	visitOrders := permutations([]int{0, 1, 2})

	for _, order := range visitOrders {
		var objs []xy.XY
		for _, index := range order {
			addr := mc.game.plan.Dungeons[index].Entrance
			objs = append(objs, addr)
		}
		objs = append(objs, mc.agent, mc.game.plan.Gate)

		nPath, nCost := multiPath(mc.agent, grid, objs...)
		if paths == nil || nCost < cost {
			paths, cost = nPath, nCost
		}
	}

	mc.paths = paths
}

// multiPath calculates the path starting at start and moving through objs in order. The
// returned values indicate the sequence of paths plus the total cost.
func multiPath(start xy.XY, grid [][]plan.Terrain, objs ...xy.XY) ([][]tNode, int) {
	ps := [][]tNode(nil)
	totalCost := 0

	from := tNode{XY: start, grid: &grid}
	for _, obj := range objs {
		to := tNode{XY: obj, grid: &grid}
		path, cost := astar.FindPath(from, to, tHeuristic)
		path = path[1:] // skips the current position (same as start)
		totalCost += cost
		from.XY = obj
		ps = append(ps, path)
	}
	return ps, totalCost
}

type xyError xy.XY

func (xye xyError) Error() string {
	return fmt.Sprintf("failed to find dungeon with x%d y%d", xye.X, xye.Y)
}

// dungeonByXY filters the dungeon slice in order to find a dungeon with the given
// coordinates.
func dungeonByXY(xy xy.XY, dungeons []plan.Dungeon) (plan.Dungeon, error) {
	for _, dungeon := range dungeons {
		if dungeon.Entrance == xy {
			return dungeon, nil
		}
	}
	return plan.Dungeon{}, xyError(xy)
}

func (mc *mapCrawler) draw(screen *ebiten.Image) {
	drawPlan(screen, mc.game.plan)
	mc.drawAgent(screen)
}

// drawAgent draws the agent with the required offset.
func (mc *mapCrawler) drawAgent(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	agentX, agentY := mc.agent.X*plan.TIS, mc.agent.Y*plan.TIS

	// the agent is taller than a single block. Therefore, he is offset in order to have
	// its feet at the start of the correct block.
	agent := images.Agent
	_, agentHeight := agent.Size()
	agentY = agentY - (agentHeight - plan.TIS)
	opts.GeoM.Translate(float64(agentX), float64(agentY))
	screen.DrawImage(images.Agent, &opts)
}

// permutations returns all permutations of slice.
func permutations(slice []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(slice, len(slice))
	return res
}
