package builtins

import (
	"bufio"
	"fmt"
)

type TypeBuiltIn struct {
	writer   *bufio.Writer
	builtIns BuiltInsMap
}

func (t TypeBuiltIn) Run(input []string) error {
	if len(input) < 2 {
		return nil
	}

	for _, el := range input[1:] {
		if _, ok := t.builtIns[el]; !ok {
			fmt.Fprintf(t.writer, "%s: not found\n", el)
		} else {
			fmt.Fprintf(t.writer, "%s is a shell builtin\n", el)
		}
	}

	return nil
}
