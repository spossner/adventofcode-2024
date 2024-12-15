package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spossner/aoc2024/internal/utils"
)

// YEAR is expected to reflect the year of the puzzles
const YEAR = 2024

// Config is a set of flags and settings used across the puzzles. Especially when parsing the input data, configurations can pre-process several steps in the input data.
// That includes stripping the data in the first step, splitting each line or even extracting all numbers from each row.
type Config struct {
	Day         int    // stores the day as int - e.g. 12
	Dev         bool   // flag indicating development - enforces loading a local development example file (default dev.txt)
	Debug       bool   // flag indicating debug mode; NewConfig function sets it to Dev as default
	DevFile     string // the dev input file - default dev.txt
	Strip       bool   // strips whitespaces of the loaded data - default true
	SplitLines  bool   // specifies whether the loaded data should be split per line; puzzle will store splitted rows in Rows property
	SplitFields bool   // splits each line by whitespaces (Fields) and stores each field in Rows (if no lines where splitted) or in Cells
	SplitWords  bool   // instead of splitting at whitespace you can specify a custom separator string (e.g. "\n\n" for 'at every empty line')
	SplitSep    string // the separator is stored in SplitSep while SplitWords indicates, that split should happen; note that you can also specify the empty string "" to split at every unicode character
	GetInts     bool   // specifies if the ints found in the data should be extracted into ParsedRows or ParsedCells
}

// ConfigFunc type used to specify optional configuration modifications in NewConfig
type ConfigFunc func(cfg *Config)

// WithSplitLines config func will set the split lines flag in the configuration.
// This is used in puzzle to split the puzzle input by lines automatically ready to use in your puzzle implementation.
func WithSplitLines() ConfigFunc {
	return func(cfg *Config) {
		cfg.SplitLines = true
	}
}

// WithSplitFields config func will set the split fields flag in the configuration.
// This is used (also after splitting the data into several lines) on every line to split the
// input at whitespaces. The strings.Fields function is used on every line or the single data input block.
func WithSplitFields() ConfigFunc {
	return func(cfg *Config) {
		cfg.SplitFields = true
	}
}

// WithGetInts config func will set the get ints flag in the configuration.
// This is a kind of fire and forget method to parse the puzzle data by extracting every number from the
// input data or from every line when data was splitted by line.
func WithGetInts() ConfigFunc {
	return func(cfg *Config) {
		cfg.GetInts = true
	}
}

// WithDevFile config func will set the given file name where local development input data is expected.
func WithDevFile(fileName string) ConfigFunc {
	return func(cfg *Config) {
		cfg.DevFile = fileName
	}
}

// WithSplitWords config func will set the split words flag and sets the specified separator in the configuration.
// Split words is used to split the data or every line not necessarily at whitespaces but at any given
// separator. If separator is the empty string, data is splitted after each utf-8 sequence. See strings.Split.
func WithSplitWords(sep string) ConfigFunc {
	return func(cfg *Config) {
		cfg.SplitWords = true
		cfg.SplitSep = sep
	}
}

// WithDebug config func will set the debug flag in the configuration.
func WithDebug() ConfigFunc {
	return func(cfg *Config) {
		cfg.Debug = true
	}
}

// NewConfig creates a new configuration for the given day and with dev flag as specified.
// It sets strip flag, dev file to "dev.txt" as default and debug flag same as dev.
// A variadic amount of ConfigFunc can be specified to overwrite any of such or other configurations.
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
