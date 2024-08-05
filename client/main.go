//go:build wasm

package main

import (
	"fmt"
	"log"

	"github.com/Gardego5/arrows/game"
	"github.com/Gardego5/arrows/game/lib/world"
	"github.com/Gardego5/arrows/lib/palette"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/mlange-42/arche/generic"
)

type Game struct {
	world.Sim
}

var (
	w           = game.CreateWorld()
	dimResource = generic.NewResource[game.Dimensions](&w.World)
	g           = &Game{w}
)

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	fmt.Printf("Layout: %d %d\n", outsideWidth, outsideHeight)

	if dimResource.Has() {
		dim := dimResource.Get()
		dim.Width, dim.Height = outsideWidth, outsideHeight
	}

	return outsideWidth, outsideHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	query := generic.NewFilter1[game.Position]().Query(&g.Sim.World)
	for query.Next() {
		pos := query.Get()

		vector.StrokeLine(screen,
			float32(pos.X-2), float32(pos.Y-2),
			float32(pos.X+2), float32(pos.Y+2),
			2, palette.Azul, true)
		vector.StrokeLine(screen,
			float32(pos.X-2), float32(pos.Y+2),
			float32(pos.X+2), float32(pos.Y-2),
			2, palette.Azul, true)
	}
}

func init() {
	ebiten.SetWindowSize(720, 480)
	ebiten.SetWindowTitle("Hello, World!")
}

func main() {
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
