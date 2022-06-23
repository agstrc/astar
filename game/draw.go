package game

import (
	"fmt"
	"image/color"

	"github.com/agstrc/heuristic-search/game/plan"
	"github.com/agstrc/heuristic-search/images"
	"github.com/agstrc/heuristic-search/xy"
	"github.com/hajimehoshi/ebiten/v2"
)

func drawPlan(screen *ebiten.Image, gamePlan *plan.Plan) {
	grid := gamePlan.Grid

	drawMainGrid(screen, grid)
	drawDungeons(screen, gamePlan.Dungeons[:])
	drawImageAt(screen, images.MasterSword, gamePlan.Sword)
	drawImageAt(screen, images.TransparentDungeon, gamePlan.Gate)

}

func drawImageAt(screen, image *ebiten.Image, at xy.XY) {
	x, y := at.X*plan.TIS, at.Y*plan.TIS
	drawOpts := ebiten.DrawImageOptions{}
	drawOpts.GeoM.Translate(float64(x), float64(y))

	screen.DrawImage(image, &drawOpts)
}

func drawDungeons(screen *ebiten.Image, dungeons []plan.Dungeon) {
	for _, dungeon := range dungeons {
		x, y := dungeon.Entrance.X*plan.TIS, dungeon.Entrance.Y*plan.TIS
		drawOpts := ebiten.DrawImageOptions{}
		drawOpts.GeoM.Translate(float64(x), float64(y))

		screen.DrawImage(images.Dungeon, &drawOpts)
	}
}

func drawMainGrid(screen *ebiten.Image, grid [][]plan.Terrain) {
	for y := range grid {
		for x := range grid {
			drawX, drawY := x*plan.TIS, y*plan.TIS
			drawOpts := ebiten.DrawImageOptions{}
			drawOpts.GeoM.Translate(float64(drawX), float64(drawY))

			screen.DrawImage(grid[y][x].Image(), &drawOpts)
		}
	}
}

var whiteBlock = func() *ebiten.Image {
	image := ebiten.NewImage(32, 32)
	image.Fill(color.White)
	return image
}()

// drawCostAt draws the given cost starting at coordinates "at". Each digit is drawn
// one "X" further to the right.
func drawCostAt(screen *ebiten.Image, cost int, at xy.XY) {
	opts := ebiten.DrawImageOptions{}
	drawX, drawY := at.X*plan.TIS, at.Y*plan.TIS
	opts.GeoM.Translate(float64(drawX), float64(drawY))

	for _, char := range fmt.Sprint(cost) {
		screen.DrawImage(whiteBlock, &opts)

		var image *ebiten.Image
		switch char {
		case '0':
			image = images.Digit0
		case '1':
			image = images.Digit1
		case '2':
			image = images.Digit2
		case '3':
			image = images.Digit3
		case '4':
			image = images.Digit4
		case '5':
			image = images.Digit5
		case '6':
			image = images.Digit6
		case '7':
			image = images.Digit7
		case '8':
			image = images.Digit8
		case '9':
			image = images.Digit9
		}
		screen.DrawImage(image, &opts)

		opts.GeoM.Translate(float64(plan.TIS), 0)
	}
}
