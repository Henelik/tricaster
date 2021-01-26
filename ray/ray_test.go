package ray

import (
	"testing"

	"github.com/Henelik/tricaster/tuple"
	"github.com/stretchr/testify/assert"
)

func TestPosition(t *testing.T) {
	r := NewRay(
		tuple.NewPoint(2, 3, 4),
		tuple.NewVector(1, 0, 0),
	)
	assert.Equal(t,
		tuple.NewPoint(2, 3, 4),
		r.Position(0),
	)
	assert.Equal(t,
		tuple.NewPoint(3, 3, 4),
		r.Position(1),
	)
	assert.Equal(t,
		tuple.NewPoint(1, 3, 4),
		r.Position(-1),
	)
	assert.Equal(t,
		tuple.NewPoint(4.5, 3, 4),
		r.Position(2.5),
	)
}
