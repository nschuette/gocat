# gocat

> [!IMPORTANT]
> This was a **learning exercise** to practice Go fundamentals: command-line argument parsing, file I/O, and string manipulation.

Gocat is a simple command-line tool that concatenates and prints files to stdout.

It reimplements the core functionality of the Unix `cat` utility, including common flags for displaying line numbers, showing whitespace, and squeezing empty lines.

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

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.