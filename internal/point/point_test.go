package point

import (
	"reflect"
	"slices"
	"testing"
)

func TestPoint_RotateLeft(t *testing.T) {
	tests := []struct {
		name string
		p    Point
		want Point
	}{
		{"up", UP, LEFT},
		{"right", RIGHT, UP},
		{"down", DOWN, RIGHT},
		{"left", LEFT, DOWN},
		{"north", NORTH, LEFT},
		{"east", EAST, UP},
		{"south", SOUTH, RIGHT},
		{"west", WEST, DOWN},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.RotateLeft(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_RotateRight(t *testing.T) {
	tests := []struct {
		name string
		p    Point
		want Point
	}{
		{"up", UP, RIGHT},
		{"right", RIGHT, DOWN},
		{"down", DOWN, LEFT},
		{"left", LEFT, UP},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.RotateRight(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Translate(t *testing.T) {
	type args struct {
		dx int
		dy int
	}
	tests := []struct {
		name string
		p    Point
		args args
		want Point
	}{
		{"simple", Point{13, 7}, args{7, 3}, Point{20, 10}},
		{"negative", Point{13, 7}, args{-20, 0}, Point{-7, 7}},
		{"no", Point{13, 7}, args{0, 0}, Point{13, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Translate(tt.args.dx, tt.args.dy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DirectAdjacentPoints(t *testing.T) {
	if !(len(DIRECT_ADJACENT_POINTS) == 4 &&
		slices.Contains(DIRECT_ADJACENT_POINTS, NORTH) &&
		slices.Contains(DIRECT_ADJACENT_POINTS, SOUTH) &&
		slices.Contains(DIRECT_ADJACENT_POINTS, WEST) &&
		slices.Contains(DIRECT_ADJACENT_POINTS, EAST)) {
		t.Errorf("DIRECT_ADJACENT_POINTS() = %v, does not contain the desired direct adjacent points N,S,W and E", DIRECT_ADJACENT_POINTS)
	}
}

func Test_AdjacentPoints(t *testing.T) {
	if !(len(ADJACENT_POINTS) == 8 &&
		slices.Contains(ADJACENT_POINTS, NORTH) &&
		slices.Contains(ADJACENT_POINTS, SOUTH) &&
		slices.Contains(ADJACENT_POINTS, WEST) &&
		slices.Contains(ADJACENT_POINTS, EAST) &&
		slices.Contains(ADJACENT_POINTS, NORTH_EAST) &&
		slices.Contains(ADJACENT_POINTS, SOUTH_EAST) &&
		slices.Contains(ADJACENT_POINTS, SOUTH_WEST) &&
		slices.Contains(ADJACENT_POINTS, NORTH_WEST)) {
		t.Errorf("ADJACENT_POINTS() = %v, does not contain the desired adjacent points N,S,W,E,NE,SE,SW and NW", ADJACENT_POINTS)
	}
}

func Test_OppositeDirections(t *testing.T) {
	if !(len(OPPOSITE_DIRECTION) == 4 &&
		OPPOSITE_DIRECTION[NORTH] == SOUTH &&
		OPPOSITE_DIRECTION[SOUTH] == NORTH &&
		OPPOSITE_DIRECTION[WEST] == EAST &&
		OPPOSITE_DIRECTION[EAST] == WEST) {
		t.Errorf("OPPOSITE_DIRECTION() = %v, does not contain the desired opposite directions", OPPOSITE_DIRECTION)
	}
}

func Test_Directions(t *testing.T) {
	tests := []struct {
		c    string
		want Point
	}{
		{"w", WEST},
		{"W", WEST},
		{"l", WEST},
		{"L", WEST},
		{"<", WEST},
		{">", EAST},
		{"r", EAST},
		{"R", EAST},
		{"e", EAST},
		{"E", EAST},
		{"^", UP},
		{"U", UP},
		{"n", NORTH},
		{"N", NORTH},
		{"d", DOWN},
		{"v", DOWN},
		{"S", SOUTH},
		{"D", SOUTH},
	}
	for _, tt := range tests {
		t.Run(tt.c, func(t *testing.T) {
			if got := DIRECTIONS[tt.c]; got != tt.want {
				t.Errorf("DIRECTIONS[%s] = %v, want %v", tt.c, got, tt.want)
			}
		})
	}
}
