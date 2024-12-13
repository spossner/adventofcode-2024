package config

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		day int
		dev bool
		fn  []ConfigFunc
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		{"default", args{12, true, nil}, &Config{
			Day: 12, Dev: true, Debug: true, DevFile: "dev.txt", Strip: true,
		}},
		{"no dev", args{22, false, nil}, &Config{
			Day: 22, DevFile: "dev.txt", Strip: true,
		}},
		{"with split lines and split fields", args{12, true, []ConfigFunc{WithSplitFields(), WithSplitLines()}}, &Config{
			Day: 12, Dev: true, Debug: true, DevFile: "dev.txt", Strip: true, SplitLines: true, SplitFields: true,
		}},
		{"with split words and get ints", args{12, true, []ConfigFunc{WithSplitWords(":"), WithGetInts()}}, &Config{
			Day: 12, Dev: true, Debug: true, DevFile: "dev.txt", Strip: true, SplitWords: true, SplitSep: ":", GetInts: true,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.day, tt.args.dev, tt.args.fn...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
