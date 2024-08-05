package game

import (
	"math"
	"math/rand"

	"github.com/Gardego5/arrows/game/lib/world"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

func CreateWorld() world.Sim {
	game := world.New()

	w, h := 500, 500
	dimID := ecs.ResourceID[Dimensions](&game.World)
	game.World.Resources().Add(dimID, &Dimensions{w, h})

	mapper := generic.NewMap3[Position, Heading, Speed](&game.World)

	for i := 0; i < 400; i++ {
		pos, heading, speed := mapper.Get(mapper.New())

		pos.X, pos.Y = rand.Float64()*float64(w), rand.Float64()*float64(h)
		*heading = Heading(rand.Float64() * 2 * math.Pi)
		*speed = Speed(1 + rand.Float64()*2)
	}

	game.AddSystem(&sysMove{ KillSpeed: 0.0001, SpinSpeed: 0.03 })

	return game
}
