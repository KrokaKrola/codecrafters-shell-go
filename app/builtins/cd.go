package builtins

import (
	"bufio"
	"fmt"
	"os"
)

type CdBuiltIn struct {
	writer *bufio.Writer
}

func (c CdBuiltIn) Run(args []string) error {
	if len(args) == 1 {
		return nil
	}

	if err := os.Chdir(args[1]); err != nil {
		return fmt.Errorf("cd: %s: No such file or directory", args[1])
	}

	return nil
}
