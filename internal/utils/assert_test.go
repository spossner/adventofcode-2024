package utils

import (
	"errors"
	"testing"
)

func TestAssertEqual(t *testing.T) {
	type args[T comparable] struct {
		a          T
		b          T
		msgAndArgs []any
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
	}
	tests := []testCase[int]{
		{"simple", args[int]{3, 3, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertEqual(tt.args.a, tt.args.b, tt.args.msgAndArgs...)
		})
	}
}

func TestAssertFalse(t *testing.T) {
	type args struct {
		value      bool
		msgAndArgs []any
	}
	tests := []struct {
		name string
		args args
	}{
		{"simple", args{false, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertFalse(tt.args.value, tt.args.msgAndArgs...)
		})
	}
}

func TestAssertNil(t *testing.T) {
	type args[T any] struct {
		value      T
		msgAndArgs []any
	}
	type testCase[T any] struct {
		name string
		args args[T]
	}
	tests := []testCase[error]{
		{"simple", args[error]{nil, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertNil(tt.args.value, tt.args.msgAndArgs...)
		})
	}
}

func TestAssertNotEqual(t *testing.T) {
	type args[T comparable] struct {
		a          T
		b          T
		msgAndArgs []any
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
	}
	tests := []testCase[string]{
		{"simple", args[string]{"seppo", "vera", nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertNotEqual(tt.args.a, tt.args.b, tt.args.msgAndArgs...)
		})
	}
}

func TestAssertNotNil(t *testing.T) {
	type args[T any] struct {
		value      T
		msgAndArgs []any
	}
	type testCase[T any] struct {
		name string
		args args[T]
	}
	tests := []testCase[error]{
		{"simple", args[error]{errors.New("test error"), nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertNotNil(tt.args.value, tt.args.msgAndArgs...)
		})
	}
}

func TestAssertTrue(t *testing.T) {
	type args struct {
		value      bool
		msgAndArgs []any
	}
	tests := []struct {
		name string
		args args
	}{
		{"simple", args{true, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertTrue(tt.args.value, tt.args.msgAndArgs...)
		})
	}
}
