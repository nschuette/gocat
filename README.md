# gocat

A simple command-line tool that concatenates and prints files to stdout.

Built as a Go learning exercise, inspired by the Unix `cat` utility.

## Usage

```bash
gocat [flags] [file ...]
gocat -                  # read from stdin
```

## Flags

| Flag | Description |
|------|-------------|
| `-n` | Number all output lines |
| `-b` | Number only non-empty lines (overrides `-n`) |
| `-s` | Suppress repeated empty lines |
| `-E` | Display `$` at the end of each line |
| `-T` | Display tabs as `^I` |
| `-A` | Shorthand for `-E -T` |

## Examples

```bash
gocat file.txt
gocat -n file.txt
gocat file1.txt file2.txt > combined.txt
echo "hello" | gocat -E
```

## Build

```bash
go build -o gocat .
```