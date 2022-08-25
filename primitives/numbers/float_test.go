package numbers

import (
	"testing"
)

func TestFloatDecPlaces(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  int
	}{
		{"test 1", 1.1233, 4},
		{"test 2", 1.1, 1},
		{"test 3", 1.0, 0},
		{"test 4", 1, 0},
		{"test 5", 0.012333, 6},
		{"test 5", 0.012333111, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatDecPlaces(tt.value); got != tt.want {
				t.Errorf("FloatDecPlaces() got = %v, want %v", got, tt.want)
			}
		})
	}
}
