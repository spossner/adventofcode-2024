package _0

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {
	fmt.Println(part1(false))
}

func Test_part2(t *testing.T) {
	fmt.Println(part2(false))
}

func TestCPU_exec(t *testing.T) {
	type fields struct {
		ptr    int
		a      int
		b      int
		c      int
		output []int
	}
	type args struct {
		op  int
		arg int
	}
	tests := []struct {
		name   string
		fields fields
		args   []args
		want   *CPU
	}{
		{"bst", fields{c: 9}, []args{{2, 6}}, &CPU{ptr: 2, b: 1, c: 9}},
		{"bxl", fields{b: 29}, []args{{1, 7}}, &CPU{ptr: 2, b: 26}},
		{"bxc", fields{b: 2024, c: 43690}, []args{{4, 0}}, &CPU{ptr: 2, b: 44354, c: 43690}},
		//{"out", fields{a: 2024, output: make([]string, 0)}, []args{{4, 0}}, &CPU{ptr: 2, b: 44354, c: 43690, output: []string{
		//	"4", "2", "5", "6", "7", "7", "7", "7", "3", "1", "0",
		//}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CPU{
				ptr:    tt.fields.ptr,
				a:      tt.fields.a,
				b:      tt.fields.b,
				c:      tt.fields.c,
				output: tt.fields.output,
			}

			for _, p := range tt.args {
				c = c.exec(p.op, p.arg)
			}

			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("exec() = %+v, want %+v", c, tt.want)
			}
		})
	}
}
