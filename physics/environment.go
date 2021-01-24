package physics

import (
	"errors"

	"github.com/Henelik/tricaster/tuple"
)

type Environment struct {
	Gravity *tuple.Tuple
	Wind    *tuple.Tuple
}

func NewEnvironment(grav, wind *tuple.Tuple) (*Environment, error) {
	if !grav.IsVector() {
		return nil, errors.New("grav must be a vector")
	}
	if !wind.IsVector() {
		return nil, errors.New("wind must be a vector")
	}
	return &Environment{grav, wind}, nil
}
