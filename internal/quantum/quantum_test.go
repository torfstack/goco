package quantum

import "testing"

func TestState_valueOfBit(t *testing.T) {
	type args struct {
		bit    int
		length int
	}
	tests := []struct {
		name string
		s    State
		args args
		want int
	}{
		{
			"zero",
			State(0),
			args{0, 1},
			0,
		},
		{
			"one",
			State(1),
			args{0, 1},
			1,
		},
		{
			"one (bit 1, does not exist)",
			State(1),
			args{1, 2},
			1,
		},
		{
			"two",
			State(2),
			args{1, 2},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.valueOfBit(tt.args.bit, tt.args.length); got != tt.want {
				t.Errorf("valueOfBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystem_DoStatesOnlyDifferInPosition(t *testing.T) {
	type fields struct {
		NumberOfQbits int
	}
	type args struct {
		i int
		p State
		r State
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"same single qbit state",
			fields{1},
			args{0, State(0), State(0)},
			false,
		},
		{
			"differing single qbit state",
			fields{1},
			args{0, State(1), State(0)},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &System{
				NumberOfQbits: tt.fields.NumberOfQbits,
				Gates:         []GateApplication{},
			}
			if got := s.DoStatesOnlyDifferInPosition(tt.args.i, tt.args.p, tt.args.r); got != tt.want {
				t.Errorf("DoStatesOnlyDifferInPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
