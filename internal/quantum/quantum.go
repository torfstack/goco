package quantum

type GateType string

const (
	GateTypeI    GateType = "I"
	GateTypeX    GateType = "X"
	GateTypeY    GateType = "Y"
	GateTypeZ    GateType = "Z"
	GateTypeH    GateType = "H"
	GateTypeCNOT GateType = "CNOT"
)

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
