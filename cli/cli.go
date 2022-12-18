package cli

import (
	"flag"
)

// Options contain the command line options passed to the program.
type Options struct {
	ASCII bool
}

// ParseOptions parses the command line options.
func ParseOptions() *Options {
	var opt Options

	flag.BoolVar(&opt.ASCII, "ascii", false, "Only use ASCII characters.")
	flag.Parse()

	return &opt
}
