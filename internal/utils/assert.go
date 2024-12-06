package utils

import (
	"fmt"
	"log"
	"reflect"
)

func formatMessage(defaultMsg string, msgAndArgs ...any) string {
	if len(msgAndArgs) == 0 {
		return defaultMsg
	}
	format, ok := msgAndArgs[0].(string)
	if !ok {
		panic("message argument to assert function must be a fmt string")
	}
	return fmt.Sprintf(format, msgAndArgs[1:]...)
}

func AssertTrue(value bool, msgAndArgs ...any) {
	if !value {
		log.Fatalln(formatMessage("Expected true but got false", msgAndArgs))
	}
}

func AssertFalse(value bool, msgAndArgs ...any) {
	if value {
		log.Fatalln(formatMessage("Expected false but got true", msgAndArgs))
	}
}

func AssertNil[T any](value T, msgAndArgs ...any) {
	v := reflect.ValueOf(value)
	if !v.IsValid() {
		return
	}
	log.Fatalf(formatMessage(fmt.Sprintf("Expected nil value but got: %v", value), msgAndArgs))
}

func AssertNotNil[T any](value T, msgAndArgs ...any) {
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
	log.Fatalf(formatMessage(fmt.Sprintf("Expected non nil value but got: %v", value), msgAndArgs))
}
