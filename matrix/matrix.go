package matrix

import "errors"

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
