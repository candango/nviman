package main

import (
	"os"

	"github.com/candango/nviman/internal/cli"
	"github.com/jessevdk/go-flags"
)

// Root options
type Options struct {
	Verbose bool `short:"v" long:"verbose" description:"Enable verbose mode"`
}

func main() {
	var opts Options

	parser := flags.NewParser(&opts, flags.Default)

	parser.AddCommand(
		"current",
		"Display the active or installed Neovim version",
		"Show the version of Neovim currently in use or switch the active version to a specific installed build.",
		&cli.CurrentCommand{})
	parser.AddCommand(
		"install",
		"Install the latest or a specific Neovim version",
		"Download and install Neovim binaries directly from official releases. Supports 'latest', 'nightly', or specific version tags.",
		&cli.InstallCommand{})
	parser.AddCommand("list",
		"List Neovim installed versions",
		"List all Neovim versions currently installed and managed by nviman on this machine.",
		&cli.ListCommand{})

	_, err := parser.Parse()
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrUnknownCommand {
			parser.WriteHelp(os.Stderr)
			os.Exit(1)
		}
		parser.WriteHelp(os.Stderr)
		os.Exit(1)
	}
}
