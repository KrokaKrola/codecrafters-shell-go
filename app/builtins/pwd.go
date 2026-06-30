package builtins

import (
	"fmt"
	"io"
	"os"
)

type PwdBuiltIn struct {
}

func (p PwdBuiltIn) Run(writer io.Writer, args []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintf(writer, "%s\n", dir); err != nil {
		return err
	}

	return nil
}
