//go:build wasm

package main

import (
	"fmt"
	"image/color"
	"log"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"honnef.co/go/js/dom/v2"
)

type game struct{}

func (g *game) Draw(screen *ebiten.Image)                         { screen.Fill(color.RGBA{R: 255}) }
func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) { return 720, 480 }
func (g *game) Update() error {
	el := dom.GetWindow().Document().QuerySelector("canvas")
	canvas, ok := el.(*dom.HTMLCanvasElement)
	if !ok {
		return fmt.Errorf("unexpected element type %T", el)
	}

	height := canvas.Style().Get("height").String()
	percentString := strings.TrimSuffix(height, "%")
	percent, err := strconv.ParseFloat(percentString, 64)
	if err != nil {
		return err
	}

	percent -= 0.2
	canvas.Style().Set("height", fmt.Sprintf("%f%%", percent))

	return nil
}

func main() {
	g := &game{}

	ebiten.SetWindowSize(720, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
