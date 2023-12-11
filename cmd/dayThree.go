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
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

// dayThreeCmd represents the dayThree command
var dayThreeCmd = &cobra.Command{
	Use:   "day-three",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if inputFile == "" {
			inputFile = "resources/day-two/input.txt"
		}

		doDayThreePartOne()
		doDayThreePartTwo()
	},
}

func doDayThreePartOne() {
	inputData, err := getInputData()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	digitMap := map[int][][]int{}

	isNumber := regexp.MustCompile("[0-9]")
	for index, line := range inputData {
		digitMap[index] = [][]int{}

		lastNumber := ""
		lastNumberStart := 999
		for ci, c := range line {
			if isNumber.MatchString(string(c)) {
				if ci < lastNumberStart {
					lastNumberStart = ci
				}

				lastNumber += string(c)
			} else if lastNumber != "" {
				lastNumberInt, _ := strconv.Atoi(lastNumber)
				digitMap[index] = append(digitMap[index], []int{lastNumberStart, ci - 1, lastNumberInt})

				lastNumber = ""
				lastNumberStart = 999
			}
		}
	}

	found := map[string]int{}

	isSymbol := regexp.MustCompile(`[a-zA-Z\.0-9]`)
	for index, line := range inputData {
		line := strings.Trim(strings.TrimSpace(line), "\r")

		isFirstLine := index == 0
		isLastLine := index == (len(inputData) - 1)

		for cIndex, c := range line {
			character := string(c)
			if !isSymbol.Match([]byte(character)) {
				isFirstCharacter := cIndex == 0
				isLastCharacter := cIndex == (len(line) - 1)

				if !isFirstLine {
					if !isFirstCharacter {
						upL := inputData[index-1][cIndex-1]
						if string(upL) != "." {
							for _, m := range digitMap[index-1] {
								if cIndex-1 >= m[0] && cIndex-1 <= m[1] {
									idx := fmt.Sprintf("%d:%d:%d", index-1, m[0], m[1])
									if _, ok := found[idx]; !ok {
										found[idx] = m[2]
									} else {
										slog.Debug("FOUND A DUPE: ", slog.Any("m", m))
									}
								}
							}
						}
					}

					up := inputData[index-1][cIndex]
					if string(up) != "." {
						for _, m := range digitMap[index-1] {
							if cIndex >= m[0] && cIndex <= m[1] {
								idx := fmt.Sprintf("%d:%d:%d", index-1, m[0], m[1])
								if _, ok := found[idx]; !ok {
									found[idx] = m[2]
								} else {
									slog.Debug("FOUND A DUPE: ", slog.Any("m", m))
								}
							}
						}
					}

					if !isLastCharacter {
						upR := inputData[index-1][cIndex+1]
						if string(upR) != "." {
							for _, m := range digitMap[index-1] {
								if cIndex+1 >= m[0] && cIndex+1 <= m[1] {
									idx := fmt.Sprintf("%d:%d:%d", index-1, m[0], m[1])
									if _, ok := found[idx]; !ok {
										found[idx] = m[2]
									} else {
										slog.Debug("FOUND A DUPE: ", slog.Any("m", m))
									}
								}
							}
						}
					}
				}

				if !isFirstCharacter {
					l := line[cIndex-1]

					// do left check
					if string(l) != "." {
						for _, m := range digitMap[index] {
							if cIndex-1 >= m[0] && cIndex-1 <= m[1] {
								idx := fmt.Sprintf("%d:%d:%d", index, m[0], m[1])
								if _, ok := found[idx]; !ok {
									found[idx] = m[2]
								} else {
									slog.Debug("FOUND A DUPE: ", slog.Any("m", m))
								}
							}
						}
					}
				}

				if !isLastCharacter {
					r := line[cIndex+1]

					// do right check
					if string(r) != "." {
						for _, m := range digitMap[index] {
							if cIndex+1 >= m[0] && cIndex+1 <= m[1] {
								idx := fmt.Sprintf("%d:%d:%d", index, m[0], m[1])
								if _, ok := found[idx]; !ok {
									found[idx] = m[2]
								} else {
									slog.Debug("FOUND A DUPE: ", slog.Any("m", m))
								}
							}
						}
					}
				}

				if !isLastLine {
					if !isFirstCharacter {
						bl := inputData[index+1][cIndex-1]
						if string(bl) != "." {
							for _, m := range digitMap[index+1] {
								if cIndex-1 >= m[0] && cIndex-1 <= m[1] {
									idx := fmt.Sprintf("%d:%d:%d", index+1, m[0], m[1])
									if _, ok := found[idx]; !ok {
										found[idx] = m[2]
									} else {
										slog.Debug("FOUND A DUPE: ", slog.Any("m", m))
									}
								}
							}
						}
					}

					b := inputData[index+1][cIndex]
					if string(b) != "." {
						for _, m := range digitMap[index+1] {
							if cIndex >= m[0] && cIndex <= m[1] {
								idx := fmt.Sprintf("%d:%d:%d", index+1, m[0], m[1])
								if _, ok := found[idx]; !ok {
									found[idx] = m[2]
								} else {
									slog.Debug("FOUND A DUPE: ", slog.Any("m", m))
								}
							}
						}
					}

					if !isLastCharacter {
						br := inputData[index+1][cIndex+1]
						if string(br) != "." {
							for _, m := range digitMap[index+1] {
								if cIndex+1 >= m[0] && cIndex+1 <= m[1] {
									idx := fmt.Sprintf("%d:%d:%d", index+1, m[0], m[1])
									if _, ok := found[idx]; !ok {
										found[idx] = m[2]
									} else {
										slog.Debug("FOUND A DUPE: ", slog.Any("m", m))
									}
								}
							}
						}
					}
				}
			}
		}
	}

	sum := 0
	for _, i := range found {
		sum += i
	}

	slog.Info("part1", slog.Int("sum", sum))
}

func doDayThreePartTwo() {
	inputData, err := getInputData()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	digitMap := map[int][][]int{}

	isNumber := regexp.MustCompile("[0-9]")
	for index, line := range inputData {
		digitMap[index] = [][]int{}

		lastNumber := ""
		lastNumberStart := 999
		for ci, c := range line {
			if isNumber.MatchString(string(c)) {
				if ci < lastNumberStart {
					lastNumberStart = ci
				}

				lastNumber += string(c)
			} else if lastNumber != "" {
				lastNumberInt, _ := strconv.Atoi(lastNumber)
				digitMap[index] = append(digitMap[index], []int{lastNumberStart, ci - 1, lastNumberInt})

				lastNumber = ""
				lastNumberStart = 999
			}
		}
	}

	found := map[string]map[string]int{}

	for index, line := range inputData {
		line := strings.Trim(strings.TrimSpace(line), "\r")

		isFirstLine := index == 0
		isLastLine := index == (len(inputData) - 1)

		for cIndex, c := range line {
			character := string(c)
			if character != "*" {
				continue
			}

			if _, ok := found[fmt.Sprintf("%d:%d", index, cIndex)]; !ok {
				found[fmt.Sprintf("%d:%d", index, cIndex)] = map[string]int{}
			}

			isFirstCharacter := cIndex == 0
			isLastCharacter := cIndex == (len(line) - 1)

			if !isFirstLine {
				if !isFirstCharacter {
					upL := inputData[index-1][cIndex-1]
					if string(upL) != "." {
						for _, m := range digitMap[index-1] {
							if cIndex-1 >= m[0] && cIndex-1 <= m[1] {
								found[fmt.Sprintf("%d:%d", index, cIndex)][fmt.Sprintf("%d:%d:%d", index-1, m[0], m[1])] = m[2]
							}
						}
					}
				}

				up := inputData[index-1][cIndex]
				if string(up) != "." {
					for _, m := range digitMap[index-1] {
						if cIndex >= m[0] && cIndex <= m[1] {
							found[fmt.Sprintf("%d:%d", index, cIndex)][fmt.Sprintf("%d:%d:%d", index-1, m[0], m[1])] = m[2]
						}
					}
				}

				if !isLastCharacter {
					upR := inputData[index-1][cIndex+1]
					if string(upR) != "." {
						for _, m := range digitMap[index-1] {
							if cIndex+1 >= m[0] && cIndex+1 <= m[1] {
								found[fmt.Sprintf("%d:%d", index, cIndex)][fmt.Sprintf("%d:%d:%d", index-1, m[0], m[1])] = m[2]
							}
						}
					}
				}
			}

			if !isFirstCharacter {
				l := line[cIndex-1]

				// do left check
				if string(l) != "." {
					for _, m := range digitMap[index] {
						if cIndex-1 >= m[0] && cIndex-1 <= m[1] {
							found[fmt.Sprintf("%d:%d", index, cIndex)][fmt.Sprintf("%d:%d:%d", index, m[0], m[1])] = m[2]
						}
					}
				}
			}

			if !isLastCharacter {
				r := line[cIndex+1]

				// do right check
				if string(r) != "." {
					for _, m := range digitMap[index] {
						if cIndex+1 >= m[0] && cIndex+1 <= m[1] {
							found[fmt.Sprintf("%d:%d", index, cIndex)][fmt.Sprintf("%d:%d:%d", index, m[0], m[1])] = m[2]
						}
					}
				}
			}

			if !isLastLine {
				if !isFirstCharacter {
					bl := inputData[index+1][cIndex-1]
					if string(bl) != "." {
						for _, m := range digitMap[index+1] {
							if cIndex-1 >= m[0] && cIndex-1 <= m[1] {
								found[fmt.Sprintf("%d:%d", index, cIndex)][fmt.Sprintf("%d:%d:%d", index+1, m[0], m[1])] = m[2]
							}
						}
					}
				}

				b := inputData[index+1][cIndex]
				if string(b) != "." {
					for _, m := range digitMap[index+1] {
						if cIndex >= m[0] && cIndex <= m[1] {
							found[fmt.Sprintf("%d:%d", index, cIndex)][fmt.Sprintf("%d:%d:%d", index+1, m[0], m[1])] = m[2]
						}
					}
				}

				if !isLastCharacter {
					br := inputData[index+1][cIndex+1]
					if string(br) != "." {
						for _, m := range digitMap[index+1] {
							if cIndex+1 >= m[0] && cIndex+1 <= m[1] {
								found[fmt.Sprintf("%d:%d", index, cIndex)][fmt.Sprintf("%d:%d:%d", index+1, m[0], m[1])] = m[2]
							}
						}
					}
				}
			}
		}
	}

	sum := 0
	for _, f := range found {
		v := maps.Values(f)
		if len(v) != 2 {
			slog.Debug("not found", slog.Any("v", v))
			continue
		}

		sum += v[0] * v[1]
	}

	slog.Info("part2", slog.Int("sum", sum))
}

func init() {
	rootCmd.AddCommand(dayThreeCmd)
}
