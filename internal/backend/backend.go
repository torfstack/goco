package backend

import (
	"goco/internal/matrix"
	"goco/internal/quantum"
)

type LinearAlgebraBackend struct {
	system *quantum.System
}

func NewLinearAlgebraBackend(system *quantum.System) *LinearAlgebraBackend {
	return &LinearAlgebraBackend{
		system: system,
	}
}

func (b *LinearAlgebraBackend) Simulate() []float64 {
	n := b.system.NumberOfQbits
	state := ZeroState()
	for range n - 1 {
		state = matrix.Tensor(state, ZeroState())
	}

	for _, gate := range b.system.Gates {
		switch gate.Type {
		case quantum.GateTypeX:
			state = matrix.Multiply(b.ConstructXGate(0), state)
		}
	}

	result := make([]float64, 1<<uint(n))
	for i := 0; i < 1<<uint(n); i++ {
		result[i] = state.Data[i][0].Length()
	}
	return result
}

func ZeroState() *matrix.Matrix {
	m := matrix.NewMatrix(2, 1)
	m.Data[0][0] = matrix.NewComplexNumber(1, 0)
	return m
}

func (b *LinearAlgebraBackend) ConstructXGate(target int) *matrix.Matrix {
	m := matrix.NewMatrix(1, 1)
	m.Data[0][0] = matrix.NewComplexNumber(1, 0)
	for i := range b.system.NumberOfQbits {
		if i == target {
			m = matrix.Tensor(m, XGate())
		} else {
			m = matrix.Tensor(m, IdentityGate())
		}
	}
	return m
}

func XGate() *matrix.Matrix {
	m := matrix.NewMatrix(2, 2)
	m.Data[0][0] = matrix.NewComplexNumber(0, 0)
	m.Data[0][1] = matrix.NewComplexNumber(1, 0)
	m.Data[1][0] = matrix.NewComplexNumber(1, 0)
	m.Data[1][1] = matrix.NewComplexNumber(0, 0)
	return m
}

func IdentityGate() *matrix.Matrix {
	m := matrix.NewMatrix(2, 2)
	m.Data[0][0] = matrix.NewComplexNumber(1, 0)
	m.Data[0][1] = matrix.NewComplexNumber(0, 0)
	m.Data[1][0] = matrix.NewComplexNumber(0, 0)
	m.Data[1][1] = matrix.NewComplexNumber(1, 0)
	return m
}
