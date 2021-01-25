package matrix

import (
	"errors"

	"github.com/Henelik/tricaster/tuple"
	"github.com/Henelik/tricaster/util"
)

type Matrix struct {
	Order int
	Data  [][]float64
}

// NewMatrix makes a matrix.
// only 2x2, 3x3, and 4x4 matrices are allowed for now
func NewMatrix(ns ...float64) (*Matrix, error) {
	var o int
	switch len(ns) {
	case 4:
		o = 2
	case 9:
		o = 3
	case 16:
		o = 4
	default:
		return nil, errors.New("matrix must be a square")
	}
	m := Matrix{
		Order: o,
		Data:  make([][]float64, o),
	}
	for i := 0; i < o; i++ {
		row := make([]float64, o)
		for j := 0; j < o; j++ {
			row[j] = ns[i*o+j]
		}
		m.Data[i] = row
	}
	return &m, nil
}

func (m *Matrix) Equal(o *Matrix) bool {
	if m.Order != o.Order {
		return false
	}
	for i := 0; i < m.Order; i++ {
		for j := 0; j < m.Order; j++ {
			if !util.Equal(m.Data[i][j], o.Data[i][j]) {
				return false
			}
		}
	}
	return true
}

func (m *Matrix) Mult(o *Matrix) *Matrix {
	if m.Order != o.Order {
		return nil
	}
	data := make([][]float64, m.Order)
	for i := 0; i < m.Order; i++ {
		row := make([]float64, m.Order)
		for j := 0; j < m.Order; j++ {
			val := 0.0
			for k := 0; k < m.Order; k++ {
				val += m.Data[i][k] * o.Data[k][j]
			}
			row[j] = val
		}
		data[i] = row
	}
	return &Matrix{m.Order, data}
}

func (m *Matrix) MultTuple(t *tuple.Tuple) *tuple.Tuple {
	return tuple.New(
		m.Data[0][0]*t.X+m.Data[0][1]*t.Y+m.Data[0][2]*t.Z+m.Data[0][3]*t.W,
		m.Data[1][0]*t.X+m.Data[1][1]*t.Y+m.Data[1][2]*t.Z+m.Data[1][3]*t.W,
		m.Data[2][0]*t.X+m.Data[2][1]*t.Y+m.Data[2][2]*t.Z+m.Data[2][3]*t.W,
		m.Data[3][0]*t.X+m.Data[3][1]*t.Y+m.Data[3][2]*t.Z+m.Data[3][3]*t.W,
	)
}
