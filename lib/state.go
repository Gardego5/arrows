package lib

import "github.com/Gardego5/goutils/env"

type (
	envVars struct {
		StaticDir string `env:"STATIC_DIR"`
		Port      int    `env:"PORT=8080"`
	}
	State struct {
		Environment *envVars
	}
)

func LoadState() (state *State, err error) {
	state = &State{}

	state.Environment, err = env.Load[envVars]()
	if err != nil {
		return
	}

	return
}
