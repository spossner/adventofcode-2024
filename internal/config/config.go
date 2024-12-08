package config

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
