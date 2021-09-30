package light

import (
	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/tuple"
)

type PointLight struct {
	Pos   *tuple.Tuple
	Color *color.Color
}
