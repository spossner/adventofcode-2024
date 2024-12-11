package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spossner/aoc2024/internal/utils"
)

const YEAR = 2024

type Config struct {
	Day         int
	Dev         bool
	DevFile     string
	Strip       bool
	SplitLines  bool
	SplitFields bool // splits each line by whitespaces (Fields
	SplitWords  bool
	SplitSep    string
	GetInts     bool
}

type ConfigFunc func(cfg *Config)

func NoLineSplit(cfg *Config) {
	cfg.SplitLines = false
}

/*func NoStrip(cfg *Config) {
	cfg.Strip = false
}*/

func SplitFields(cfg *Config) {
	cfg.SplitFields = true
}

func GetInts(cfg *Config) {
	cfg.GetInts = true
}

func WithDevFile(fileName string) ConfigFunc {
	return func(cfg *Config) {
		cfg.DevFile = fileName
	}
}

func WithSplitWords(sep string) ConfigFunc {
	return func(cfg *Config) {
		cfg.SplitWords = true
		cfg.SplitSep = sep
	}
}

func NewConfig(day int, dev bool, fn ...ConfigFunc) *Config {
	envFile := fmt.Sprintf("%s/.env", utils.GetProjectDir())
	if err := godotenv.Load(envFile); err != nil {
		fmt.Printf(".env file not found: %s\n", envFile)
	} else {
		fmt.Printf("using .env file: %s\n", envFile)
	}

	cfg := &Config{
		Day:         day,
		Dev:         dev,
		DevFile:     "dev.txt",
		Strip:       true,
		SplitLines:  true,
		SplitFields: false,
		SplitWords:  false,
		SplitSep:    "",
		GetInts:     false,
	}
	for _, f := range fn {
		f(cfg)
	}
	return cfg
}
