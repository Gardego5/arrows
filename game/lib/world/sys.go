package world

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/mlange-42/arche/ecs"
)

type (
	sysID reflect.Type

	Sys interface {
		Update(*ecs.World) error
	}
	InitializeSys interface {
		Sys
		Initialize(*ecs.World)
	}
	DestroySys interface {
		Sys
		Destroy(*ecs.World)
	}

	Sim struct {
		ecs.World
		systems map[sysID]Sys
	}
)

func New() Sim {
	game := Sim{World: ecs.NewWorld(), systems: map[sysID]Sys{}}

	return game
}

func (game *Sim) AddSystem(sys Sys) sysID {
	id := sysID(reflect.TypeOf(sys))
	if _, ok := game.systems[id]; ok {
		panic(fmt.Errorf("system %T already exists", sys))
	} else {
		if isys, ok := sys.(InitializeSys); ok {
			isys.Initialize(&game.World)
		}
		game.systems[id] = sys
		return id
	}
}

func (game *Sim) RemoveSystem(id sysID) {
	if sys, ok := game.systems[id]; ok {
		if dsys, ok := sys.(DestroySys); ok {
			dsys.Destroy(&game.World)
		}
	}
	delete(game.systems, id)
}

func (game *Sim) Update() error {
	errs := []error{}

	for _, sys := range game.systems {
		if err := sys.Update(&game.World); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}
