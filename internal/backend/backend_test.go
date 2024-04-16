package backend

import (
	"github.com/stretchr/testify/assert"
	"goco/internal/quantum"
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
		{
			"two qbits, second into X gate",
			func() *quantum.System {
				s := quantum.NewSystem(2)
				s.X(1)
				return &s
			},
			[]float64{0, 1, 0, 0},
		},
		{
			"one qbit into H gate",
			func() *quantum.System {
				s := quantum.NewSystem(1)
				s.H(0)
				return &s
			},
			[]float64{0.5, 0.5},
		},
		{
			"one qbit into X and H gates",
			func() *quantum.System {
				s := quantum.NewSystem(1)
				s.X(0)
				s.H(0)
				return &s
			},
			[]float64{0.5, 0.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewLinearAlgebraBackend(tt.fields())
			got := b.Simulate()

			for i, v := range got {
				assert.InDeltaf(t, tt.want[i], v, 0.0001,
					"Simulate(): expected %f, got %f (in arrays expected %v, got %v)", tt.want[i], v, tt.want, got)
			}
		})
	}
}
