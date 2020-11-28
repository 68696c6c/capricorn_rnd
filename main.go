package main

import (
	"os"

	"github.com/68696c6c/capricorn_rnd/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "capricorn",
		Short: "Root command for capricorn",
	}

	rootCmd.SetOutput(os.Stdout)
	rootCmd.AddCommand(
		cmd.GenerateGoat,
	)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
