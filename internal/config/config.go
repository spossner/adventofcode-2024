package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spossner/aoc2024/internal/utils"
)

const YEAR = 2024

type Config struct {
	Day         int    // stores the day as int - e.g. 12
	Dev         bool   // flag indicating development - enforces loading a local development example file (default dev.txt)
	Debug       bool   // flag indicating debug mode; NewConfig function sets it to Dev as default
	DevFile     string // the dev input file - default dev.txt
	Strip       bool   // strips whitespaces of the loaded data - default true
	SplitLines  bool   // specified whether or not the loaded data should be split per line; puzzle will store splitted rows in Rows property
	SplitFields bool   // splits each line by whitespaces (Fields) and stores each field in Rows (if no lines where splitted) or in Cells
	SplitWords  bool   // instead of splitting at whitespace you can specify a custom separator string (e.g. "\n\n" for 'at every empty line')
	SplitSep    string // the separator is stored in SplitSep while SplitWords indicates, that split should happen; note that you can also specify the empty string "" to split at every unicode character
	GetInts     bool   // specifies if the ints found in the data should be extracted into ParsedRows or ParsedCells
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
		Day:     day,
		Dev:     dev,
		Debug:   dev,
		DevFile: "dev.txt",
		Strip:   true,
	}
	for _, f := range fn {
		f(cfg)
	}
	return cfg
}
