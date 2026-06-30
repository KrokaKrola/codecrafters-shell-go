package builtins

import (
	"fmt"
	"io"
	"strings"
)

type EchoBuiltIn struct {
}

func (e EchoBuiltIn) Run(writer io.Writer, input []string) error {
	if len(input) < 2 {
		fmt.Fprintln(writer)
		return nil
	}

	args := input[1:]

	fmt.Fprintln(writer, strings.Join(args, " "))
	return nil
}
