package goco

import (
	"goco/pkg/matrix"
	goco "goco/pkg/quantum"
	"math"
)

type LinearAlgebraBackend struct {
	system *goco.System
}

func NewLinearAlgebraBackend(system *goco.System) *LinearAlgebraBackend {
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
		case goco.GateTypeX:
			state = matrix.Multiply(b.ConstructXGate(gate.Qbits[0]), state)
		case goco.GateTypeH:
			state = matrix.Multiply(b.ConstructHadamardGate(gate.Qbits[0]), state)
		case goco.GateTypeCNOT:
			state = matrix.Multiply(b.ConstructCNOTGate(gate.Qbits[0], gate.Qbits[1]), state)
		}
	}

	result := make([]float64, 1<<uint(n))
	for i := 0; i < 1<<uint(n); i++ {
		result[i] = math.Pow(state.Data[i][0].Length(), 2)
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

func (b *LinearAlgebraBackend) ConstructHadamardGate(target int) *matrix.Matrix {
	m := matrix.NewMatrix(1, 1)
	m.Data[0][0] = matrix.NewComplexNumber(1, 0)
	for i := range b.system.NumberOfQbits {
		if i == target {
			m = matrix.Tensor(m, HadamardGate())
		} else {
			m = matrix.Tensor(m, IdentityGate())
		}
	}
	return m
}

func (b *LinearAlgebraBackend) ConstructCNOTGate(control, target int) *matrix.Matrix {
	s := b.system
	n := s.NumberOfQbits
	m := matrix.NewMatrix(1<<n, 1<<n)
	for i := 0; i < 1<<n; i++ {
		for j := 0; j < 1<<n; j++ {
			stateI, errI := s.StateOf(i)
			stateJ, errJ := s.StateOf(j)
			if errI != nil || errJ != nil {
				panic("ConstructCNOTGate: called with invalid parameters " +
					"(control or target bits are most likely out of range")
			}
			if i == j && s.ValueOfBitInState(control, stateJ) == 0 {
				m.Data[i][j] = matrix.NewComplexNumber(1, 0)
			} else if s.ValueOfBitInState(control, stateJ) == 1 &&
				s.ValueOfBitInState(control, stateI) == 1 &&
				s.ValueOfBitInState(target, stateJ) != s.ValueOfBitInState(target, stateI) &&
				s.DoStatesOnlyDifferInPosition(target, stateI, stateJ) {
				m.Data[i][j] = matrix.NewComplexNumber(1, 0)
			} else {
				m.Data[i][j] = matrix.NewComplexNumber(0, 0)
			}
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

func HadamardGate() *matrix.Matrix {
	m := matrix.NewMatrix(2, 2)
	m.Data[0][0] = matrix.NewComplexNumber(1.0/math.Sqrt(2), 0)
	m.Data[0][1] = matrix.NewComplexNumber(1.0/math.Sqrt(2), 0)
	m.Data[1][0] = matrix.NewComplexNumber(1.0/math.Sqrt(2), 0)
	m.Data[1][1] = matrix.NewComplexNumber(-1.0/math.Sqrt(2), 0)
	return m
}
