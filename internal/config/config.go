package config

const YEAR = 2024

type Config struct {
	Day        int
	Dev        bool
	Strip      bool
	SplitLines bool
	SplitWords bool
	ParseInts  bool
}

type ConfigFunc func(cfg *Config)

func NoLineSplit(cfg *Config) {
	cfg.SplitLines = false
}
func NoStrip(cfg *Config) {
	cfg.Strip = false
}
func SplitWords(cfg *Config) {
	cfg.SplitWords = true
}
func ParseInts(cfg *Config) {
	cfg.ParseInts = true
}

func NewConfig(day int, dev bool, fn ...ConfigFunc) *Config {
	cfg := &Config{
		Day:        day,
		Dev:        dev,
		Strip:      true,
		SplitLines: true,
		SplitWords: false,
		ParseInts:  false,
	}
	for _, f := range fn {
		f(cfg)
	}
	return cfg
}
