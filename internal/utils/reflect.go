package utils

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func GetProjectDir() string {
	_, fpath, _, ok := runtime.Caller(1)
	if !ok {
		err := errors.New("failed to get package directory")
		panic(err)
	}

	path := filepath.Dir(fpath)

	packageRoot, ok := strings.CutSuffix(path, "/internal/config")
	if !ok {
		log.Fatalf("error extracting package directory from %s", path)
	}
	return packageRoot
}

func GetPackageDir() int {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		err := errors.New("failed to get package directory")
		panic(err)
	}

	dir := filepath.Dir(path)
	vol := filepath.VolumeName(path)
	i := len(dir) - 1
	for i >= len(vol) && !os.IsPathSeparator(dir[i]) {
		i--
	}
	packageNumber, err := strconv.Atoi(dir[i+1:])
	if err != nil {
		log.Fatalf("error parsing puzzle package - must be the number of the day: %w", err)
	}
	return packageNumber
}
