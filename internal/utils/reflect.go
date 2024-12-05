package utils

import (
	"errors"
	"path/filepath"
	"runtime"
	"strings"
)

func GetFileName() string {
	_, fpath, _, ok := runtime.Caller(1)
	if !ok {
		err := errors.New("failed to get filename")
		panic(err)
	}
	filename := filepath.Base(fpath)
	filename = strings.Replace(filename, ".go", "", 1)
	return filename
}
