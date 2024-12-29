package _0

import (
	"fmt"
	"testing"
)

func Test_part1(t *testing.T) {
	fmt.Println(part1(false))
}

func Test_part2(t *testing.T) {
	fmt.Println(part2(false))
}

func Test_seller_mix(t *testing.T) {
	type fields struct {
		secret int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"simple", fields{42}, args{15}, 37},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := seller{
				secret: tt.fields.secret,
			}
			if got := mixAndPrune(s.secret, tt.args.value); got != tt.want {
				t.Errorf("mixAndPrune() = %v, want %v", got, tt.want)
			}
		})
	}
}
