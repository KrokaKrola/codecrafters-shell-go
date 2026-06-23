package builtins

import (
	"bufio"
	"fmt"
	"os"
)

type PwdBuiltIn struct {
	writer *bufio.Writer
}

func (p PwdBuiltIn) Run([]string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintf(p.writer, "%s\n", dir); err != nil {
		return err
	}

	return nil
}
