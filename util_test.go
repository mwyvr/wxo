package wxo_test

import (
	"testing"

	"github.com/solutionroute/wxo"
)

func TestDirectionsFromDegrees(t *testing.T) {
	tests := []struct {
		name    string
		deg     int
		asArrow bool
		want    string
	}{
		{"0", 0, false, "N"},
		{"NE", 40, false, "NE"},
		{"NE", 40, true, "↗"},
		{"NE", 45, false, "NE"},
		{"NE", 50, false, "NE"},
		{"W", 270, false, "W"},
		{"W", 292, false, "WNW"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wxo.DirectionFromDegree(tt.deg, tt.asArrow); got != tt.want {
				t.Errorf("DirectionFromDegree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrowFromOrdinal(t *testing.T) {
	tests := []struct {
		name string
		ord  string
		want string
	}{
		{"NE", "NE", "↗"},
		{"ne", "ne", "↗"},
		{"foo", "FoO", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wxo.ArrowFromOrdinal(tt.ord); got != tt.want {
				t.Errorf("ArrowFromOrdinal() = %v, want %v", got, tt.want)
			}
		})
	}
}
