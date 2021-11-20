package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

type Command interface {
	Description() string
	Usage() string
	Run(args []string) error
}

var (
	modeCommand = make(map[string]Command)
	modeFlagSet = make(map[string]*flag.FlagSet)

	Stderr = os.Stderr
)

func RegisterCommand(mode string, initCmd func(flagSet *flag.FlagSet) Command) {
	if _, found := modeCommand[mode]; found {
		panic("command already registered: " + mode)
	}
	flagSet := flag.NewFlagSet(mode, flag.ExitOnError)
	modeCommand[mode] = initCmd(flagSet)
	modeFlagSet[mode] = flagSet
	flagSet.SetOutput(Stderr)
	flagSet.Usage = func() {
		help(mode)
	}
}

func main() {
	var err error
	flag.CommandLine.SetOutput(Stderr)
	flag.Usage = func() {
		fmt.Fprintf(Stderr, "Management CLI tool for Gallery CSUI.\n")
		usage("")
	}

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
		os.Exit(0)
	}

	mode := args[0]
	if cmd, ok := modeCommand[mode]; ok {
		cmdFlag := modeFlagSet[mode]
		if err = cmdFlag.Parse(args[1:]); err != nil {
			os.Exit(1)
		}
		if err = cmd.Run(cmdFlag.Args()); err != nil {
			fmt.Fprintf(Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	} else {
		usage("unknown command: " + mode)
		os.Exit(1)
	}
}

func usage(msg string) {
	if msg != "" {
		fmt.Fprintf(Stderr, "Error: %s\n", msg)
	}

	fmt.Fprintf(Stderr, `
Usage: gallery-cli <mode> [commandopts] [commandargs]

Modes:

`,
	)
	modes := make([]string, 0, len(modeCommand))
	for m, cmd := range modeCommand {
		modes = append(modes, fmt.Sprintf("  %s: %s\n", m, cmd.Description()))
	}
	sort.Strings(modes)
	for _, m := range modes {
		fmt.Fprintf(Stderr, "%s", m)
	}

	fmt.Fprintf(Stderr, `
For help in each mode:

  gallery-cli <mode> --help

`)
}

func help(mode string) {
	cmd := modeCommand[mode]
	flags := modeFlagSet[mode]
	fmt.Fprintf(Stderr, "%s\n\n", cmd.Description())
	fmt.Fprintf(Stderr, "%s\n\n", cmd.Usage())
	fmt.Fprintf(Stderr, "Options:\n")
	flags.PrintDefaults()
}
