package game

import (
	"math"
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

type sysMove struct {
	dimension generic.Resource[Dimensions]
	movers    generic.Filter3[Position, Heading, Speed]

	KillSpeed Speed
	SpinSpeed Heading
}

func (s *sysMove) Initialize(world *ecs.World) {
	s.dimension = generic.NewResource[Dimensions](world)
	s.movers = *generic.NewFilter3[Position, Heading, Speed]()
}

func (s *sysMove) Update(world *ecs.World) error {
	// Get the dimensions
	dim := s.dimension.Get()

	// Get the query
	query := s.movers.Query(world)

	// Iterate over all entities
	for query.Next() {
		// Get the components
		pos, heading, speed := query.Get()

		// Update the position
		pos.X += math.Cos(float64(*heading)) * float64(*speed)
		pos.Y += math.Sin(float64(*heading)) * float64(*speed)

		// Wrap around the screen
		pos.X = math.Mod(pos.X+float64(dim.Width), float64(dim.Width))
		pos.Y = math.Mod(pos.Y+float64(dim.Height), float64(dim.Height))

		// Spiral
		*heading += s.SpinSpeed + Heading(rand.Float64()*0.002)
		heading.Wrap()

		// Slow down
		*speed *= 1 - s.KillSpeed
	}

	return nil
}
