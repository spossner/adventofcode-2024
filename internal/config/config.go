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
	Debug       bool
	DevFile     string
	Strip       bool
	SplitLines  bool
	SplitFields bool // splits each line by whitespaces (Fields
	SplitWords  bool
	SplitSep    string
	GetInts     bool
}

type ConfigFunc func(cfg *Config)

func WithSplitLines() ConfigFunc {
	return func(cfg *Config) {
		cfg.SplitLines = true
	}
}

func WithSplitFields() ConfigFunc {
	return func(cfg *Config) {
		cfg.SplitFields = true
	}
}

func WithGetInts() ConfigFunc {
	return func(cfg *Config) {
		cfg.GetInts = true
	}
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

func WithDebug() ConfigFunc {
	return func(cfg *Config) {
		cfg.Debug = true
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
		Debug:       dev,
		DevFile:     "dev.txt",
		Strip:       true,
		SplitLines:  false,
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
