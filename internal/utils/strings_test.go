package utils

import (
	"reflect"
	"testing"
)

func TestGetInts(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"simple", args{"13 multiplied by 2 is 26"}, []int{13, 2, 26}, false},
		{"no ints found", args{"Hello, World!"}, []int{}, false},
		{"with plus sign", args{"13 multiplied by 2 is +26"}, []int{13, 2, 26}, false},
		{"with minus sign", args{"13 multiplied by -2 is -26"}, []int{13, -2, -26}, false},
		{"tons of numbers", args{"1,2,3,4,5,6,7,8,9,10,11,12"}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetInts(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBatchedStrings(t *testing.T) {
	type args struct {
		text string
		n    int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"simple", args{"Sebastian", 3}, []string{"Seb", "ast", "ian"}},
		{"unaligned", args{"Sebastian", 4}, []string{"Seba", "stia", "n"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, b := range BatchedStrings(tt.args.text, tt.args.n) {
				if !reflect.DeepEqual(b, tt.want[i]) {
					t.Errorf("BatchedStrings()[%d] = %v, want %v", i, b, tt.want)
				}
			}
		})
	}
}
