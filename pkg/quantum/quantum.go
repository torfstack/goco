package goco

import (
	"errors"
	"strconv"
)

type GateType string

const (
	GateTypeI    GateType = "I"
	GateTypeX    GateType = "X"
	GateTypeY    GateType = "Y"
	GateTypeZ    GateType = "Z"
	GateTypeH    GateType = "H"
	GateTypeCNOT GateType = "CNOT"
)

var ErrCreatingStateFromInt = errors.New("cannot create state from int value, int value too large")

type GateApplication struct {
	Type  GateType
	Qbits []int
}

type System struct {
	NumberOfQbits int
	Gates         []GateApplication
}

func NewSystem(numberOfQbits int) System {
	return System{NumberOfQbits: numberOfQbits}
}

func (s *System) I(qbit int) {
	if qbit >= s.NumberOfQbits {
		panic("qbit index out of range")
	}
	s.Gates = append(s.Gates, GateApplication{Type: GateTypeI, Qbits: []int{qbit}})
}

func (s *System) X(qbit int) {
	if qbit >= s.NumberOfQbits {
		panic("qbit index out of range")
	}
	s.Gates = append(s.Gates, GateApplication{Type: GateTypeX, Qbits: []int{qbit}})
}

func (s *System) Y(qbit int) {
	if qbit >= s.NumberOfQbits {
		panic("qbit index out of range")
	}
	s.Gates = append(s.Gates, GateApplication{Type: GateTypeY, Qbits: []int{qbit}})
}

func (s *System) Z(qbit int) {
	if qbit >= s.NumberOfQbits {
		panic("qbit index out of range")
	}
	s.Gates = append(s.Gates, GateApplication{Type: GateTypeZ, Qbits: []int{qbit}})
}

func (s *System) H(qbit int) {
	if qbit >= s.NumberOfQbits {
		panic("qbit index out of range")
	}
	s.Gates = append(s.Gates, GateApplication{Type: GateTypeH, Qbits: []int{qbit}})
}

func (s *System) CNOT(control, target int) {
	if control >= s.NumberOfQbits || target >= s.NumberOfQbits {
		panic("qbit index out of range")
	}
	s.Gates = append(s.Gates, GateApplication{Type: GateTypeCNOT, Qbits: []int{control, target}})
}

func (s *System) States() []State {
	bits := make([]State, s.NumberOfQbits)
	for i := range bits {
		bits[i] = State(i)
	}
	return bits
}

func (s *System) ValueOfBitInState(bit int, state State) int {
	return state.valueOfBit(bit, s.NumberOfQbits)
}

func (s *System) DoStatesOnlyDifferInPosition(i int, p, r State) bool {
	n := s.NumberOfQbits
	for j := 0; j < n; j++ {
		if p.valueOfBit(j, n) != r.valueOfBit(j, n) && j != i ||
			p.valueOfBit(j, n) == r.valueOfBit(j, n) && j == i {
			return false
		}
	}
	return true
}

func (s *System) StateOf(i int) (State, error) {
	state := State(i)
	if i < 0 || state.numberOfBits() > s.NumberOfQbits {
		return 0, ErrCreatingStateFromInt
	}
	return state, nil
}

type State int

func (s State) numberOfBits() int {
	return len(strconv.FormatInt(int64(s), 2))
}

// ValueOfBit returns the value of the bit at the given position.
// The most significant bit is at position 0.
func (s State) valueOfBit(bit, length int) int {
	bits := strconv.FormatInt(int64(s), 2)
	zeroPrefixLength := length - len(bits)
	for i := 0; i < zeroPrefixLength; i++ {
		bits = "0" + bits
	}
	parsed, err := strconv.ParseInt(string(bits[bit]), 2, 0)
	if err != nil {
		panic(err)
	}
	return int(parsed)
}
