package main

import (
	"fmt"
	"os"

	"github.com/agstrc/heuristic-search/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(900, 900)
	ebiten.SetWindowTitle("A*")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	args := os.Args
	var g *game.Game
	if len(args) > 1 {
		var err error
		g, err = game.GameFromJSON(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create a game from JSON file:", err)
			os.Exit(1)
		}
	} else {
		g = game.DefaultGame()
	}

	if err := ebiten.RunGame(g); err != nil {
		fmt.Fprintln(os.Stderr, "An error occurred:", err)
		os.Exit(1)
	}
}
