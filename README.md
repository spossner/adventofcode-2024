# Advent of code 2024

Juhhhuuuu - now in GO :-D

## PRE-CONFIGURATION
- update `YEAR` in `internal/config/config.go`

## 3rd party libs
- `github.com/joho/godotenv`: to parse the .env file

## New day
- Clone the `00` blueprint directory for the new day
- Note: you must use the day (e.g. 06 for 6th of december) in order to make the automated day detection and input download working

## Puzzle input
Puzzle input is automatically fetched when the folder is named like the day and a valid advent of code session cookie is specified in environment as `AOC_SESSION`.
You can specify the AOC_SESSION in a .env file. Therefore copy the `.env.sample` file into `.env` and
set your own SESSION cookie. See "Getting a valid AOC_SESSION" below.

Or if unit tests are invoked from IDE, the environment variable can also be set in the calling IDE.
In Goland it's in Settings > Go Modules > Environment
`AOC_SESSION=12345......fedcba98`

### Getting a valid AOC_SESSION
- log into adventofcode.com with your browser
- search for `session` cookie in network tab
- copy & paste the value into the settings with AOC_SESSION key
