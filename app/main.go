package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/builtins"
	"github.com/codecrafters-io/shell-starter-go/app/executables"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)
	builtIns := builtins.NewBuiltIns(writer)

	for {
		fmt.Print("$ ")

		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		lineFields := strings.Fields(line)

		if len(lineFields) == 0 {
			continue
		}

		cmd := lineFields[0]

		command, ok := builtIns[cmd]
		if !ok {
			if err := executables.RunExecutable(writer, lineFields); err != nil {
				fmt.Fprintln(os.Stderr, "error:", err.Error())
			}

			writer.Flush()
			continue
		}

		if err := command.Run(lineFields); err != nil {
			fmt.Fprintln(writer, err.Error())
			continue
		}

		writer.Flush()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
}
