/*
Copyright © 2023 Jack Stupple <jack.stupple@protonmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

const (
	maxRed   int = 12
	maxGreen int = 13
	maxBlue  int = 14
)

// dayTwoCmd represents the dayTwo command
var dayTwoCmd = &cobra.Command{
	Use:   "day-two",
	Short: "https://adventofcode.com/2023/day/2",
	Run: func(cmd *cobra.Command, args []string) {
		if inputFile == "" {
			inputFile = "resources/day-two/input.txt"
		}

		doDayTwoPartOne()
		doDayTwoPartTwo()
	},
}

func doDayTwoPartOne() {
	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	lines := strings.Split(string(inputData), "\n")

	idReg := regexp.MustCompile(`Game (\d+)`)
	idSum := 0

	for _, line := range lines {
		matches := idReg.FindAllStringSubmatch(line, -1)
		gameID, err := strconv.Atoi(matches[0][1])
		if err != nil {
			slog.Error("failed to get game ID", slog.String("error", err.Error()))
			return
		}

		parts := strings.Split(line, ":")
		games := strings.Split(parts[1], ";")

		possible := true

		for _, game := range games {
			gamePlays := strings.Split(game, ",")
			// slog.Debug("handling game", slog.Int("id", gameID), slog.Any("gamePlays", gamePlays))
			for _, gamePlay := range gamePlays {
				gamePlay = strings.TrimSpace(gamePlay)
				if len(gamePlay) == 0 {
					continue
				}

				g := strings.Split(gamePlay, " ")
				count, err := strconv.Atoi(g[0])
				if err != nil {
					slog.Error("failed to get count", slog.String("error", err.Error()))
					return
				}

				colour := g[1]

				switch colour {
				case "red":
					if count > maxRed {
						possible = false
					}
				case "green":
					if count > maxGreen {
						possible = false
					}
				case "blue":
					if count > maxBlue {
						possible = false
					}
				default:
					panic(fmt.Errorf("what the world?!"))
				}

				if !possible {
					break
				}
			}

			if !possible {
				break
			}
		}

		if possible {
			slog.Debug("game possible", slog.Int("id", gameID), slog.String("games", parts[1]))
			idSum += gameID
		}
	}

	slog.Info("part 1", slog.Int("sum", idSum))
}

func doDayTwoPartTwo() {
	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	lines := strings.Split(string(inputData), "\n")

	// idReg := regexp.MustCompile(`Game (\d+)`)
	idSum := 0

	for _, line := range lines {
		smallestPossible := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		// matches := idReg.FindAllStringSubmatch(line, -1)
		// gameID, err := strconv.Atoi(matches[0][1])
		// if err != nil {
		// 	slog.Error("failed to get game ID", slog.String("error", err.Error()))
		// 	return
		// }

		parts := strings.Split(line, ":")
		games := strings.Split(parts[1], ";")

		for _, game := range games {
			gamePlays := strings.Split(game, ",")
			// slog.Debug("handling game", slog.Int("id", gameID), slog.Any("gamePlays", gamePlays))
			for _, gamePlay := range gamePlays {
				gamePlay = strings.TrimSpace(gamePlay)
				if len(gamePlay) == 0 {
					continue
				}

				g := strings.Split(gamePlay, " ")
				count, err := strconv.Atoi(g[0])
				if err != nil {
					slog.Error("failed to get count", slog.String("error", err.Error()))
					return
				}

				colour := g[1]
				if count >= smallestPossible[colour] {
					smallestPossible[colour] = count
				}
			}
		}

		pow := smallestPossible["red"] * smallestPossible["green"] * smallestPossible["blue"]
		slog.Debug("smallest possible", slog.Int("pow", pow), slog.Any("smallest", smallestPossible))
		idSum += pow
	}

	slog.Info("part 2", slog.Int("sum", idSum))
}

func init() {
	rootCmd.AddCommand(dayTwoCmd)
}
