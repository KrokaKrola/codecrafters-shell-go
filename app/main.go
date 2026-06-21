package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	exitCommand string = "exit"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")

		if !scanner.Scan() {
			break
		}

		cmd := scanner.Text()

		switch cmd {
		case exitCommand:
			os.Exit(0)
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
