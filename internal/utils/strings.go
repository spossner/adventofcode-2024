package utils

import (
	"fmt"
	"iter"
	"regexp"
	"strconv"
)

func GetInts(s string) ([]int, error) {
	re := regexp.MustCompile(`([-+]?\d+)`)
	matches := re.FindAllSubmatch([]byte(s), -1)
	ints := make([]int, 0)
	for _, m := range matches {
		n, err := strconv.Atoi(string(m[0]))
		if err != nil {
			return nil, err
		}
		ints = append(ints, n)
	}
	return ints, nil
}

func FormatMsgAndArgs(defaultMessage string, msgAndArgs ...any) string {
	if len(msgAndArgs) == 0 {
		return defaultMessage
	}
	format, ok := msgAndArgs[0].(string)
	if !ok {
		panic("message argument to assert function must be a fmt string")
	}
	return fmt.Sprintf(format, msgAndArgs[1:]...)
}

func BatchedStrings(text string, n int) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		loop := 0
		for i := 0; i < len(text); i += n {
			if !yield(loop, text[i:min(i+n, len(text))]) {
				break
			}
			loop++
		}
	}
}
