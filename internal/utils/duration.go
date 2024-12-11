package utils

import (
	"fmt"
	"time"
)

// Duration returns a method which prints the time since calling this Duration method.
// Can be used as deferred method to track execution time of functions.
func Duration(name string) func() {
	start := time.Now()
	return func() {
		duration := time.Since(start)
		fmt.Printf("%s: took %v\n", name, duration)
	}
}
