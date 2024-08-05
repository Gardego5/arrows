package game

import (
	"math"

	"github.com/Gardego5/arrows/game/lib/vec"
)

type Position vec.Vec2f

type Speed float64

type Heading float64

func (h Heading) Direction() vec.Vec2f {
	a := float64(h)
	return vec.New(math.Cos(a), math.Sin(a))
}

func (h *Heading) Wrap() {
	*h = Heading(math.Mod(float64(*h), 2*math.Pi))
}
