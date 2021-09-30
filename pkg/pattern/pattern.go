package pattern

import (
	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/tuple"
)

type Pattern interface {
	Process(pos *tuple.Tuple) *color.Color
}
