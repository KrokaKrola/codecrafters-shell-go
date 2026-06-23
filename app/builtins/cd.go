package builtins

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

type CdBuiltIn struct {
	writer *bufio.Writer
}

func (c CdBuiltIn) Run(args []string) error {
	if len(args) == 1 {
		return nil
	}

	// abs path handling
	if path.IsAbs(args[1]) {
		if err := os.Chdir(args[1]); err != nil {
			return fmt.Errorf("cd: %s: No such file or directory", args[1])
		}

		return nil
	}

	panic("unimplemented")
}
