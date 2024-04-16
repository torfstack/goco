package backend

import (
	"goco/internal/quantum"
	"reflect"
	"testing"
)

func TestLinearAlgebraBackend_Simulate(t *testing.T) {
	tests := []struct {
		name   string
		fields func() *quantum.System
		want   []float64
	}{
		{
			"one qbit into no gates",
			func() *quantum.System {
				s := quantum.NewSystem(1)
				return &s
			},
			[]float64{1, 0},
		},
		{
			"two qbits into no gates",
			func() *quantum.System {
				s := quantum.NewSystem(2)
				return &s
			},
			[]float64{1, 0, 0, 0},
		},
		{
			"one qbit into X gate",
			func() *quantum.System {
				s := quantum.NewSystem(1)
				s.X(0)
				return &s
			},
			[]float64{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewLinearAlgebraBackend(tt.fields())
			if got := b.Simulate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Simulate() = %v, want %v", got, tt.want)
			}
		})
	}
}
