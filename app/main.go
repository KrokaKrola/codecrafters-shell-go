package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	exitCommand string = "exit"
	echoCommand string = "echo"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")

		if !scanner.Scan() {
			break
		}

		line := scanner.Text()

		splits := strings.Fields(line)

		if len(splits) == 0 {
			fmt.Fprintln(os.Stderr, "invalid number of arguments")
			break
		}

		cmd := splits[0]

		switch cmd {
		case exitCommand:
			os.Exit(0)
		case echoCommand:
			if len(splits) < 2 {
				fmt.Println()
				break
			}

			args := splits[1:]

			fmt.Println(strings.Join(args, " "))
		default:
			fmt.Fprintf(writer, "%s: command not found\n", cmd)
		}

		writer.Flush()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
}
