package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type options struct {
	numberLines    bool
	numberNonEmpty bool
	squeeze        bool
	showEnds       bool
	showTabs       bool
	showAll        bool
}

func main() {
	opts := &options{}

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

func registerFlags(opts *options) {
	flag.BoolVar(&opts.numberLines, "n", false, "Number all output lines")
	flag.BoolVar(&opts.numberNonEmpty, "b", false, "Number only non-empty lines")
	flag.BoolVar(&opts.squeeze, "s", false, "Suppress repeated empty lines")
	flag.BoolVar(&opts.showEnds, "E", false, "Display $ at the end of each line")
	flag.BoolVar(&opts.showTabs, "T", false, "Display tabs as ^I")
	flag.BoolVar(&opts.showAll, "A", false, "Shorthand for -E -T")
}

func openFile(file string) (*os.File, error) {
	if file == "-" {
		return os.Stdin, nil
	}

	return os.Open(file)
}

func processFile(file *os.File, opts *options) {
	lineNum := 0
	lastEmptyLine := false

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lineNum++

		line := sc.Text()
		isEmptyLine := line == ""

		if opts.squeeze && isEmptyLine && lastEmptyLine {
			continue
		}
		lastEmptyLine = isEmptyLine

		shouldNumberLine := opts.numberLines || (opts.numberNonEmpty && line != "")
		if shouldNumberLine {
			line = fmt.Sprintf("%6d\t%s", lineNum, line)
		}
		if opts.showEnds || opts.showAll {
			line = fmt.Sprintf("%s$", line)
		}
		if opts.showTabs || opts.showAll {
			line = strings.ReplaceAll(line, "\t", "^I")
		}

		fmt.Fprintln(os.Stdout, line)
	}
}
