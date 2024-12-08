package puzzle

import (
	"errors"
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/utils"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct {
	Data        string
	Rows        []string
	Cells       [][]string
	ParsedRows  []int
	ParsedCells [][]int
}

func fileNameByDay(cfg *config.Config) string {
	if cfg.Dev {
		return fmt.Sprintf(cfg.DevFile)
	}
	return fmt.Sprintf("%02d.txt", cfg.Day)
}

func loadPuzzleInput(cfg *config.Config) ([]byte, error) {
	fileName := fileNameByDay(cfg)
	fmt.Println(fileName)
	return os.ReadFile(fileName)
}

func writePuzzleInput(cfg *config.Config, data []byte) error {
	return os.WriteFile(fileNameByDay(cfg), data, 0644)
}

func downloadPuzzleInput(day int) ([]byte, error) {
	aocSession := os.Getenv("AOC_SESSION")
	if aocSession == "" {
		log.Fatalln(`missing AOC_SESSION in env
Hint: if using goland you can set the AOC_SESSION env variable in Settings > Go > Go Modules > Environment globally`)
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", config.YEAR, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Cookie", fmt.Sprintf("session=%s", aocSession))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewPuzzle(cfg *config.Config) *Puzzle {
	fmt.Println("NEW PUZZlE FOR DAY", cfg.Day)
	data, err := loadPuzzleInput(cfg)
	if errors.Is(err, os.ErrNotExist) {
		if cfg.Dev {
			log.Fatalln("can not download dev test data")
		}
		data, err = downloadPuzzleInput(cfg.Day)
		if err := writePuzzleInput(cfg, data); err != nil {
			fmt.Printf("error caching puzzle input: %v\n", err)
		}
	}
	raw := string(data)
	if cfg.Strip {
		raw = strings.TrimSpace(raw)
	}

	var rows []string
	var cells [][]string
	var parsedRows []int
	var parsedCells [][]int
	if cfg.SplitLines {
		rows = strings.Split(raw, "\n")

		if cfg.SplitFields {
			for _, row := range rows {
				cells = append(cells, strings.Fields(row))
			}
		} else if cfg.SplitWords {
			for _, row := range rows {
				cells = append(cells, strings.Split(row, cfg.SplitSep))
			}
		}
	} else if cfg.SplitFields {
		rows = strings.Fields(raw)
	} else if cfg.SplitWords {
		rows = strings.Split(raw, cfg.SplitSep)
	}

	if cfg.GetInts {
		for i, row := range rows {
			parsed, err := utils.GetInts(row)
			if err != nil {
				log.Fatalf("error parsing line %d:  %s -  %v", i+1, row, err)
			}
			switch len(parsed) {
			case 0:
				log.Fatalf("no number found in line %d:  %s", i+1, row)
			case 1:
				parsedRows = append(parsedRows, parsed[0])
			default:
				parsedCells = append(parsedCells, parsed)
			}
		}

		for i, row := range cells {
			var parsed []int

			for j, cell := range row {
				n, err := strconv.Atoi(cell)
				if err != nil {
					log.Fatalf("error parsing %d/%d:  %s -  %v", i+1, j+1, cell, err)
				}
				parsed = append(parsed, n)
			}
			parsedCells = append(parsedCells, parsed)
		}
	}

	return &Puzzle{
		Data:        raw,
		Rows:        rows,
		Cells:       cells,
		ParsedRows:  parsedRows,
		ParsedCells: parsedCells,
	}
}
