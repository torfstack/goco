package matrix

import (
	"math"
)

type ComplexNumber struct {
	Real      float64
	Imaginary float64
}

func NewComplexNumber(real, imaginary float64) ComplexNumber {
	return ComplexNumber{Real: real, Imaginary: imaginary}
}

func (a ComplexNumber) Add(b ComplexNumber) ComplexNumber {
	return ComplexNumber{Real: a.Real + b.Real, Imaginary: a.Imaginary + b.Imaginary}
}

func (a ComplexNumber) Sub(b ComplexNumber) ComplexNumber {
	return ComplexNumber{Real: a.Real - b.Real, Imaginary: a.Imaginary - b.Imaginary}
}

func (a ComplexNumber) Mul(b ComplexNumber) ComplexNumber {
	return ComplexNumber{
		Real:      a.Real*b.Real - a.Imaginary*b.Imaginary,
		Imaginary: a.Real*b.Imaginary + a.Imaginary*b.Real,
	}
}

func (a ComplexNumber) Length() float64 {
	return math.Sqrt(a.Real*a.Real + a.Imaginary*a.Imaginary)
}

type Matrix struct {
	Rows    int
	Columns int
	Data    [][]ComplexNumber
}

func NewMatrix(rows, columns int) *Matrix {
	data := make([][]ComplexNumber, rows)
	for i := range data {
		data[i] = make([]ComplexNumber, columns)
	}
	return &Matrix{Rows: rows, Columns: columns, Data: data}
}

func Tensor(a, b *Matrix) *Matrix {
	result := NewMatrix(a.Rows*b.Rows, a.Columns*b.Columns)
	for i := 0; i < a.Rows; i++ {
		for j := 0; j < a.Columns; j++ {
			for k := 0; k < b.Rows; k++ {
				for l := 0; l < b.Columns; l++ {
					result.Data[i*b.Rows+k][j*b.Columns+l] = a.Data[i][j].Mul(b.Data[k][l])
				}
			}
		}
	}
	return result
}

func Multiply(a, b *Matrix) *Matrix {
	if a.Columns != b.Rows {
		return nil
	}
	result := NewMatrix(a.Rows, b.Columns)
	for i := 0; i < a.Rows; i++ {
		for j := 0; j < b.Columns; j++ {
			for k := 0; k < a.Columns; k++ {
				result.Data[i][j] = result.Data[i][j].Add(a.Data[i][k].Mul(b.Data[k][j]))
			}
		}
	}
	return result
}
