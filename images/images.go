// Package images exports some images which are loaded at compile time and decoded at
// runtime.
package images

import (
	"bytes"
	_ "embed"
	"fmt"
	"image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed assets/agent.png
	agentBytes []byte
	//go:embed assets/dungeon.png
	dungeonBytes []byte
	//go:embed assets/forest.png
	forestBytes []byte
	//go:embed assets/grass.png
	grassBytes []byte
	//go:embed assets/master_sword.png
	masterSwordBytes []byte
	//go:embed assets/mountain.png
	mountainBytes []byte
	//go:embed assets/non_traversable.png
	nonTraversableBytes []byte
	//go:embed assets/sand.png
	sandBytes []byte
	//go:embed assets/transparent_dungeon.png
	transparentDungeonBytes []byte
	//go:embed assets/traversable.png
	traversableBytes []byte
	//go:embed assets/virtue01.png
	virtue1Bytes []byte
	//go:embed assets/virtue02.png
	virtue2Bytes []byte
	//go:embed assets/virtue03.png
	virtue3Bytes []byte
	//go:embed assets/water.png
	waterBytes []byte

	//go:embed assets/00.png
	digit0Bytes []byte
	//go:embed assets/01.png
	digit1Bytes []byte
	//go:embed assets/02.png
	digit2Bytes []byte
	//go:embed assets/03.png
	digit3Bytes []byte
	//go:embed assets/04.png
	digit4Bytes []byte
	//go:embed assets/05.png
	digit5Bytes []byte
	//go:embed assets/06.png
	digit6Bytes []byte
	//go:embed assets/07.png
	digit7Bytes []byte
	//go:embed assets/08.png
	digit8Bytes []byte
	//go:embed assets/09.png
	digit9Bytes []byte
)

var (
	Agent              = mustDecode(agentBytes)
	Dungeon            = mustDecode(dungeonBytes)
	Forest             = mustDecode(forestBytes)
	Grass              = mustDecode(grassBytes)
	MasterSword        = mustDecode(masterSwordBytes)
	Mountain           = mustDecode(mountainBytes)
	NonTraversable     = mustDecode(nonTraversableBytes)
	Sand               = mustDecode(sandBytes)
	TransparentDungeon = mustDecode(transparentDungeonBytes)
	Traversable        = mustDecode(traversableBytes)
	Virtue1            = mustDecode(virtue1Bytes)
	Virtue2            = mustDecode(virtue2Bytes)
	Virtue3            = mustDecode(virtue3Bytes)
	Water              = mustDecode(waterBytes)

	Digit0 = mustDecode(digit0Bytes)
	Digit1 = mustDecode(digit1Bytes)
	Digit2 = mustDecode(digit2Bytes)
	Digit3 = mustDecode(digit3Bytes)
	Digit4 = mustDecode(digit4Bytes)
	Digit5 = mustDecode(digit5Bytes)
	Digit6 = mustDecode(digit6Bytes)
	Digit7 = mustDecode(digit7Bytes)
	Digit8 = mustDecode(digit8Bytes)
	Digit9 = mustDecode(digit9Bytes)
)

// mustDecode decodes b as a PNG image and panics if any errors occur.
func mustDecode(b []byte) *ebiten.Image {
	image, err := png.Decode(bytes.NewReader(b))
	if err != nil {
		message := fmt.Sprintf("Failed to decode image: %v", err)
		panic(message)
	}
	return ebiten.NewImageFromImage(image)
}
