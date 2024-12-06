package utils

import (
	"fmt"
	"log"
	"reflect"
)

func AssertZero[T any](value T, msgAndArgs ...any) {
	v := reflect.ValueOf(value)
	if !v.IsValid() {
		return
	}
	log.Fatalf(fmt.Sprintf("Expected nil value but got: %v", value), msgAndArgs)
}

func AssertNotZero[T any](value T, msgAndArgs ...any) {
	v := reflect.ValueOf(value)
	if v.IsValid() {
		kind := v.Kind()
		if !(kind == reflect.Ptr ||
			kind == reflect.Interface ||
			kind == reflect.Slice ||
			kind == reflect.Map ||
			kind == reflect.Chan ||
			kind == reflect.Func) ||
			v.IsNil() {
			return
		}
	}
	log.Fatalf(fmt.Sprintf("Expected nil value but got: %v", value), msgAndArgs)
}
