/*
Copyright © 2023 Purple Wifi Ltd
*/
package cmd

import (
	"bytes"
	_ "embed"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

/**
 * Generated by ChatGPT
 * Prompt: "Using go as my programming language I need a map of every number between 0 and 50 which maps the numeric
 * value to the word name of that number, for example "one""
 */
var names = []string{"zero", "one", "two", "three", "four", "five",
	"six", "seven", "eight", "nine"}

// dayOneCmd represents the dayOne command
var dayOneCmd = &cobra.Command{
	Use:   "day-one",
	Short: "https://adventofcode.com/2023/day/1",
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("doing part 1")
		getCalibrationSum()

		slog.Info("doing part 2")
		getCalibrationSumWithStringValues()
	},
}

func init() {
	rootCmd.AddCommand(dayOneCmd)
}

func getCalibrationSum() {
	inputData, err := os.ReadFile("resources/day-one/input.txt")
	if err != nil {
		slog.Error(err.Error())
		return
	}

	lines := strings.Split(string(inputData), "\n")

	reg := regexp.MustCompile(`\d`)

	calibrationSum := 0
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		digits := string(bytes.Join(reg.FindAll([]byte(l), -1), []byte("")))

		calibrationValueA := digits[:1]
		calibrationValueB := calibrationValueA

		slog.Debug("digits", slog.String("value", digits))

		if len(digits) > 1 {
			calibrationValueB = digits[len(digits)-1:]
		}

		slog.Debug(
			"calibration values",
			slog.String("input", l),
			slog.String("a", calibrationValueA),
			slog.String("b", calibrationValueB),
		)

		combined, err := strconv.Atoi(calibrationValueA + calibrationValueB)
		if err != nil {
			slog.Error(err.Error())
			return
		}

		calibrationSum += combined

		slog.Debug("new sum", slog.Int("sum", calibrationSum))
	}

	slog.Info("got the sum", slog.Int("sum", calibrationSum))
}

func getCalibrationSumWithStringValues() {
	inputData, err := os.ReadFile("resources/day-one/input.txt")
	if err != nil {
		slog.Error(err.Error())
		return
	}

	lines := strings.Split(string(inputData), "\n")

	calibrationSum := 0
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		calibrationValueA := ""
		calibrationValueB := ""

		aIndex := len(l)
		bIndex := 0

		for i := 0; i <= 9; i++ {
			firstIndex := strings.Index(l, strconv.Itoa(i))
			lastIndex := strings.LastIndex(l, strconv.Itoa(i))
			if firstIndex > -1 && firstIndex < aIndex {
				aIndex = firstIndex
				calibrationValueA = l[firstIndex : firstIndex+1]
			}

			if lastIndex >= bIndex {
				bIndex = lastIndex
				calibrationValueB = l[lastIndex : lastIndex+1]
			}
		}

		for k, n := range names {
			firstIndex := strings.Index(l, n)
			lastIndex := strings.LastIndex(l, n)
			if firstIndex > -1 && firstIndex < aIndex {
				aIndex = firstIndex
				calibrationValueA = strconv.Itoa(k)
			}

			if lastIndex >= bIndex {
				bIndex = lastIndex
				calibrationValueB = strconv.Itoa(k)
			}
		}

		slog.Debug(
			"calibration values",
			slog.String("input", l),
			slog.String("a", calibrationValueA),
			slog.String("b", calibrationValueB),
		)

		combined, err := strconv.Atoi(calibrationValueA + calibrationValueB)
		if err != nil {
			slog.Error(err.Error())
			return
		}

		calibrationSum += combined

		slog.Debug("new sum", slog.Int("sum", calibrationSum))
	}

	slog.Info("got the sum", slog.Int("sum", calibrationSum))
}