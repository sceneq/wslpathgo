package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// "-u"
func WindowsToWSL(path string) string {
	path = strings.ReplaceAll(path, "\\", "/")
	if !strings.HasPrefix(path, ".") && len(path) > 1 && path[1] == ':' {
		path = "/mnt/" + strings.ToLower(path[:1]) + path[2:]
	}
	return path
}

// "-w"
func WSLToWindows(path string) string {
	if strings.Contains(path, "/mnt/") {
		driveLetter := strings.ToUpper(string(path[5])) + ":\\"
		return driveLetter + strings.ReplaceAll(path[7:], "/", "\\")
	} else {
		return strings.ReplaceAll(path, "/", "\\")
	}
}

func main() {
	var (
		win2WSL bool
		WSL2Win bool
	)
	flag.BoolVar(&win2WSL, "u", false, "translate from a Windows path to a WSL path (default)")
	flag.BoolVar(&WSL2Win, "w", false, "translate from a WSL path to a Windows path")
	flag.Parse()

	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	stdin := (fi.Mode()&os.ModeNamedPipe != 0) || (fi.Mode()&os.ModeCharDevice == 0)

	if !stdin {
		for _, input := range flag.Args() {
			var result string
			switch {
			case win2WSL:
				result = WindowsToWSL(input)
			case WSL2Win:
				result = WSLToWindows(input)
			default:
				result = WindowsToWSL(input)
			}
			fmt.Println(result)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		writer := bufio.NewWriter(os.Stdout)
		for scanner.Scan() {
			input := scanner.Text()

			var result string
			switch {
			case win2WSL:
				result = WindowsToWSL(input)
			case WSL2Win:
				result = WSLToWindows(input)
			default:
				result = WindowsToWSL(input)
			}

			fmt.Fprintln(writer, result)
		}
		writer.Flush()

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}

}
