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

	path := args[1]

	if args[1] == "~" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		path = homedir
	}

	if err := os.Chdir(path); err != nil {
		return fmt.Errorf("cd: %s: No such file or directory", args[1])
	}

	return nil
}
