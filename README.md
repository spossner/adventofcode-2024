# Advent of code 2024

Juhhhuuuu - now in GO :-D

## PRE-CONFIGURATION
- update `YEAR` in `internal/config/config.go`

## 3rd party libs
- none at the moment - aoc helpers included this year

## New day
- Clone the `00` blueprint directory for the new day
- update the `DAY` in `<new-day>/main.go`

## Puzzle input

Puzzle input is automatically fetched when the valid advent of code session cookie is specified in environment as `AOC_SESSION`.
Due to the unit test setup and IDE invocation, the environment variable should be set in the calling IDE.
In Goland it's in Settings > Go Modules > Environment
`AOC_SESSION=12345......fedcba98`

### Getting a valid AOC_SESSION
- log into adventofcode.com with your browser
- search for `session` cookie in network tab
- copy & paste the value into the settings with AOC_SESSION key
