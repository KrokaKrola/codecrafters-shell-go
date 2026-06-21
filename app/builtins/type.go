package builtins

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os/exec"
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
		if _, ok := t.builtIns[el]; ok {
			fmt.Fprintf(t.writer, "%s is a shell builtin\n", el)
			continue
		}

		path, err := exec.LookPath(el)
		if err != nil {
			if errors.Is(err, fs.ErrPermission) {
				continue
			}

			fmt.Fprintf(t.writer, "%s: not found\n", el)
			continue
		}

		fmt.Fprintf(t.writer, "%s is %s\n", el, path)
	}

	return nil
}
