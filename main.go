package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Options struct {
	NumberLines    bool
	NumberNonEmpty bool
	Squeeze        bool
	ShowEnds       bool
	ShowTabs       bool
	ShowAll        bool
}

func main() {
	opts := &Options{}
	registerFlags(opts)

	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		files = []string{"-"}
	}

	for _, file := range files {
		file, err := openFile(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		processFile(file, opts)
		if file != os.Stdin {
			file.Close()
		}
	}
}

func registerFlags(opts *Options) {
	flag.BoolVar(&opts.NumberLines, "n", false, "Number all output lines")
	flag.BoolVar(&opts.NumberNonEmpty, "b", false, "Number only non-empty lines")
	flag.BoolVar(&opts.Squeeze, "s", false, "Suppress repeated empty lines")
	flag.BoolVar(&opts.ShowEnds, "E", false, "Display $ at the end of each line")
	flag.BoolVar(&opts.ShowTabs, "T", false, "Display tabs as ^I")
	flag.BoolVar(&opts.ShowAll, "A", false, "Shorthand for -E -T")
}

func openFile(file string) (*os.File, error) {
	if file == "-" {
		return os.Stdin, nil
	}

	return os.Open(file)
}

func processFile(file *os.File, opts *Options) {
	lineNum := 0
	lastEmptyLine := false

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lineNum++

		line := sc.Text()
		isEmptyLine := line == ""

		if opts.Squeeze && isEmptyLine && lastEmptyLine {
			continue
		}
		lastEmptyLine = isEmptyLine

		shouldNumberLine := opts.NumberLines || (opts.NumberNonEmpty && line != "")
		if shouldNumberLine {
			line = fmt.Sprintf("%6d\t%s", lineNum, line)
		}
		if opts.ShowEnds || opts.ShowAll {
			line = fmt.Sprintf("%s$", line)
		}
		if opts.ShowTabs || opts.ShowAll {
			line = strings.ReplaceAll(line, "\t", "^I")
		}

		fmt.Fprintln(os.Stdout, line)
	}
}
