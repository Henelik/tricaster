package shading

import (
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/tuple"
)

type Pattern interface {
	Process(pos *tuple.Tuple) *color.Color
}
