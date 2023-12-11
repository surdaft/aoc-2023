/*
Copyright © 2023 Purple Wifi Ltd
*/
package cmd

import (
	"log"
	"log/slog"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// dayFourCmd represents the dayFour command
var dayFourCmd = &cobra.Command{
	Use:   "day-four",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if inputFile == "" {
			inputFile = "resources/day-four/input.txt"
		}

		doDayFourPartOne()
	},
}

func doDayFourPartOne() {
	inputData, err := getInputData()
	if err != nil {
		log.Panic(err)
	}

	sumPoints := 0
	for _, line := range inputData {
		if len(line) == 0 {
			continue
		}

		cardNumber := getCardNumber(line)
		sides := parseSideData(line)

		lineLogger := slog.With(slog.Int("card", cardNumber))
		lineLogger.Debug("parsed data", slog.Any("sides", sides))

		points := 0
		for _, m := range sides["mine"] {
			winningIndex := slices.Index(sides["winning"], m)
			if winningIndex < 0 {
				continue
			}

			if points == 0 {
				points += 1
			} else {
				points *= 2
			}

			lineLogger.Debug("new points total", slog.Int("points", points))
		}

		sumPoints += points
	}

	slog.Info("part 1", slog.Int("sum", sumPoints))
}

var reg *regexp.Regexp

func getCardNumber(line string) int {
	if reg == nil {
		reg = regexp.MustCompile(`Card\s+(\d+)`)
	}

	matches := reg.FindAllStringSubmatch(line, 1)
	slog.Debug("matches", slog.Any("matches", matches))
	conv, err := strconv.Atoi(matches[0][1])
	if err != nil {
		log.Panic(err)
	}

	return conv
}

func parseSideData(line string) map[string][]int {
	sides := strings.Split(line[(strings.Index(line, ":")+1):], " | ")
	leftStrings := strings.Split(sides[0], " ")
	rightStrings := strings.Split(sides[1], " ")

	left := make([]int, 0)
	right := make([]int, 0)

	for _, l := range leftStrings {
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			continue
		}

		in, err := strconv.Atoi(l)
		if err != nil {
			log.Panic(err)
		}

		left = append(left, in)
	}

	for _, l := range rightStrings {
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			continue
		}

		in, err := strconv.Atoi(l)
		if err != nil {
			log.Panic(err)
		}

		right = append(right, in)
	}

	return map[string][]int{
		"mine":    left[:],
		"winning": right[:],
	}
}

func init() {
	rootCmd.AddCommand(dayFourCmd)
}
