/*
Copyright © 2023 Purple Wifi Ltd
*/
package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var debug bool = false

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc-2023",
	Short: "Advent of code 2023",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		level := new(slog.LevelVar)
		level.Set(slog.LevelInfo)

		if debug {
			level.Set(slog.LevelDebug)
		}

		sl := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		}))

		slog.SetDefault(sl)

		slog.Info("logger initialised", slog.String("level", level.String()))
	})
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "extra verbose output")
}
