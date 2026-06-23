package builtins

import (
	"bufio"
	"fmt"
	"strings"
)

type EchoBuiltIn struct {
	writer *bufio.Writer
}

func (e EchoBuiltIn) Run(input []string) error {
	if len(input) < 2 {
		fmt.Fprintln(e.writer)
		return nil
	}

	args := input[1:]

	fmt.Fprintln(e.writer, strings.Join(args, " "))
	return nil
}
