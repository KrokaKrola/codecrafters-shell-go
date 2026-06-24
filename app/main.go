package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/builtins"
	"github.com/codecrafters-io/shell-starter-go/app/executables"
	"github.com/codecrafters-io/shell-starter-go/app/tokenizer"
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

		args, err := tokenizer.Tokenize(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err.Error())
			writer.Flush()
			continue
		}

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		command, ok := builtIns[cmd]
		if !ok {
			if err := executables.RunExecutable(writer, args); err != nil {
				fmt.Fprintln(os.Stderr, "error:", err.Error())
			}

			writer.Flush()
			continue
		}

		if err := command.Run(args); err != nil {
			fmt.Fprintln(writer, err.Error())
			writer.Flush()
			continue
		}

		writer.Flush()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
}
