package point

import (
	"reflect"
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
