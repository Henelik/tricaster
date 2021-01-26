package matrix

import (
	"errors"
	"math"

	"github.com/Henelik/tricaster/tuple"
	"github.com/Henelik/tricaster/util"
)

var (
	Identity = &Matrix{
		Order: 4,
		Data: [][]float64{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}
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
	order := util.MinInt(m.Order, o.Order)
	data := make([][]float64, order)
	for i := 0; i < order; i++ {
		data[i] = make([]float64, order)
		for j := 0; j < order; j++ {
			for k := 0; k < order; k++ {
				data[i][j] += m.Data[i][k] * o.Data[k][j]
			}
		}
	}
	return &Matrix{order, data}
}

func (m *Matrix) MultTuple(t *tuple.Tuple) *tuple.Tuple {
	return tuple.New(
		m.Data[0][0]*t.X+m.Data[0][1]*t.Y+m.Data[0][2]*t.Z+m.Data[0][3]*t.W,
		m.Data[1][0]*t.X+m.Data[1][1]*t.Y+m.Data[1][2]*t.Z+m.Data[1][3]*t.W,
		m.Data[2][0]*t.X+m.Data[2][1]*t.Y+m.Data[2][2]*t.Z+m.Data[2][3]*t.W,
		m.Data[3][0]*t.X+m.Data[3][1]*t.Y+m.Data[3][2]*t.Z+m.Data[3][3]*t.W,
	)
}

func (m *Matrix) Transpose() *Matrix {
	data := make([][]float64, m.Order)
	for i := 0; i < m.Order; i++ {
		data[i] = make([]float64, m.Order)
		for j := 0; j < m.Order; j++ {
			data[i][j] = m.Data[j][i]
		}
	}
	return &Matrix{m.Order, data}
}

func (m *Matrix) Determinant() float64 {
	if m.Order == 2 {
		return m.Data[0][0]*m.Data[1][1] - m.Data[0][1]*m.Data[1][0]
	}
	det := 0.0
	for i := 0; i < m.Order; i++ {
		det += m.Cofactor(0, i) * m.Data[0][i]
	}
	return det
}

func (m *Matrix) Submatrix(x, y int) *Matrix {
	order := m.Order - 1
	data := make([][]float64, order)
	for i := 0; i < order; i++ {
		data[i] = make([]float64, order)
		for j := 0; j < order; j++ {
			iIndex := i
			jIndex := j
			if i >= x {
				iIndex += 1
			}
			if j >= y {
				jIndex += 1
			}
			data[i][j] = m.Data[iIndex][jIndex]
		}
	}
	return &Matrix{order, data}
}

func (m *Matrix) Minor(x, y int) float64 {
	return m.Submatrix(x, y).Determinant()
}

func (m *Matrix) Cofactor(x, y int) float64 {
	if (x+y)%2 != 0 {
		return m.Minor(x, y) * -1
	}
	return m.Minor(x, y)
}

func (m *Matrix) IsInvertible() bool {
	return m.Determinant() != 0.0
}

func (m *Matrix) Inverse() *Matrix {
	det := m.Determinant()
	data := make([][]float64, m.Order)
	for i := 0; i < m.Order; i++ {
		data[i] = make([]float64, m.Order)
		for j := 0; j < m.Order; j++ {
			data[i][j] = m.Cofactor(j, i) / det
		}
	}
	return &Matrix{m.Order, data}
}

func Translation(x, y, z float64) *Matrix {
	return &Matrix{
		Order: 4,
		Data: [][]float64{
			{1, 0, 0, x},
			{0, 1, 0, y},
			{0, 0, 1, z},
			{0, 0, 0, 1},
		},
	}
}

func Scaling(x, y, z float64) *Matrix {
	return &Matrix{
		Order: 4,
		Data: [][]float64{
			{x, 0, 0, 0},
			{0, y, 0, 0},
			{0, 0, z, 0},
			{0, 0, 0, 1},
		},
	}
}

func RotationX(r float64) *Matrix {
	return &Matrix{
		Order: 4,
		Data: [][]float64{
			{1, 0, 0, 0},
			{0, math.Cos(r), -math.Sin(r), 0},
			{0, math.Sin(r), math.Cos(r), 0},
			{0, 0, 0, 1},
		},
	}
}

func RotationY(r float64) *Matrix {
	return &Matrix{
		Order: 4,
		Data: [][]float64{
			{math.Cos(r), 0, math.Sin(r), 0},
			{0, 1, 0, 0},
			{-math.Sin(r), 0, math.Cos(r), 0},
			{0, 0, 0, 1},
		},
	}
}

func RotationZ(r float64) *Matrix {
	return &Matrix{
		Order: 4,
		Data: [][]float64{
			[]float64{math.Cos(r), -math.Sin(r), 0, 0},
			[]float64{math.Sin(r), math.Cos(r), 0, 0},
			[]float64{0, 0, 1, 0},
			[]float64{0, 0, 0, 1},
		},
	}
}
