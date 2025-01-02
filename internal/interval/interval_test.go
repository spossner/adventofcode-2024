package interval

import (
	"reflect"
	"slices"
	"testing"
)

func TestInterval_All(t *testing.T) {
	type fields struct {
		From int
		To   int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{"simple", fields{3, 16}, []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
		{"empty", fields{3, 3}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Interval{
				From: tt.fields.From,
				To:   tt.fields.To,
			}
			if got := slices.Collect(i.All()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_Contains(t *testing.T) {
	type fields struct {
		From int
		To   int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"simple", fields{3, 16}, args{value: 8}, true},
		{"outside", fields{3, 16}, args{value: 18}, false},
		{"lower end inside", fields{3, 16}, args{value: 3}, true},
		{"upper end inside", fields{3, 16}, args{value: 15}, true},
		{"lower end outside", fields{3, 16}, args{value: 2}, false},
		{"upper end outside", fields{3, 16}, args{value: 16}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Interval{
				From: tt.fields.From,
				To:   tt.fields.To,
			}
			if got := i.Contains(tt.args.value); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_Intersect(t *testing.T) {
	type fields struct {
		From int
		To   int
	}
	type args struct {
		other Interval
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Interval
	}{
		{"simple", fields{3, 16}, args{Interval{7, 20}}, Interval{7, 16}},
		{"included", fields{3, 16}, args{Interval{7, 12}}, Interval{7, 12}},
		{"left from this interval", fields{13, 26}, args{Interval{7, 20}}, Interval{13, 20}},
		{"no overlap", fields{3, 16}, args{Interval{17, 20}}, Interval{}},
		{"on the edge", fields{3, 16}, args{Interval{16, 20}}, Interval{}},
		{"on the edge plus one", fields{3, 17}, args{Interval{16, 20}}, Interval{16, 17}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Interval{
				From: tt.fields.From,
				To:   tt.fields.To,
			}
			if got := i.Intersect(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_Len(t *testing.T) {
	type fields struct {
		From int
		To   int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"simple", fields{3, 16}, 13},
		{"smaller", fields{7, 12}, 5},
		{"empty", fields{}, 0},
		{"single", fields{16, 17}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Interval{
				From: tt.fields.From,
				To:   tt.fields.To,
			}
			if got := i.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_Overlaps(t *testing.T) {
	type fields struct {
		From int
		To   int
	}
	type args struct {
		other Interval
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"simple", fields{3, 16}, args{Interval{7, 20}}, true},
		{"included", fields{3, 16}, args{Interval{7, 12}}, true},
		{"left from this interval", fields{13, 26}, args{Interval{7, 20}}, true},
		{"no overlap", fields{3, 16}, args{Interval{17, 20}}, false},
		{"on the edge", fields{3, 16}, args{Interval{16, 20}}, false},
		{"on the edge plus one", fields{3, 17}, args{Interval{16, 20}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Interval{
				From: tt.fields.From,
				To:   tt.fields.To,
			}
			if got := i.Overlaps(tt.args.other); got != tt.want {
				t.Errorf("Overlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
