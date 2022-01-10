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
		{"N", 359, false, "N"},
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

func TestTruncateWebString(t *testing.T) {
	type args struct {
		s      string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"nil string", args{"", 10}, ""},
		{"no truncate", args{"0123456789", 10}, "0123456789"},
		{"truncate at 5", args{"0123456789", 6}, "012345"},
		{"rune at pos 10", args{"012345678界", 10}, "012345678界"}, // len() would report 13 bytes
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wxo.TruncateWebString(tt.args.s, tt.args.length); got != tt.want {
				t.Errorf("TruncateWebString() = %v, want %v", got, tt.want)
			}
		})
	}
}
