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
	defaultWriter := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)
	builtIns := builtins.NewBuiltIns()

	for {
		fmt.Print("$ ")

		if !scanner.Scan() {
			break
		}

		line := scanner.Text()

		args, redirects, err := tokenizer.Tokenize(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err.Error())
			defaultWriter.Flush()
			continue
		}

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		writer := defaultWriter
		errWriter := os.Stderr

		for _, redirect := range redirects {
			switch redirect.Type {
			case tokenizer.StdoutRedirect, tokenizer.StdoutAppend:
				flag := os.O_CREATE | os.O_WRONLY
				if redirect.Type == tokenizer.StdoutAppend {
					flag = os.O_CREATE | os.O_APPEND | os.O_WRONLY
				}

				file, err := os.OpenFile(redirect.Value, flag, 0644)

				if err != nil {
					fmt.Fprintln(errWriter, "error:", err.Error())
					continue
				}

				writer = bufio.NewWriter(file)
			case tokenizer.StderrRedirect:
				flag := os.O_CREATE | os.O_WRONLY
				if redirect.Type == tokenizer.StderrAppend {
					flag = os.O_CREATE | os.O_APPEND | os.O_WRONLY
				}

				file, err := os.OpenFile(redirect.Value, flag, 0644)

				if err != nil {
					fmt.Fprintln(errWriter, "error:", err.Error())
					continue
				}

				errWriter = file
			}

		}

		command, ok := builtIns.Get(cmd)
		if !ok {
			executables.RunExecutable(writer, errWriter, args)

			writer.Flush()
			continue
		}

		if err := command.Run(writer, args); err != nil {
			fmt.Fprintln(errWriter, err.Error())
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
