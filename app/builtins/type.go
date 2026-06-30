package builtins

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os/exec"
)

type TypeBuiltIn struct {
	builtIns *BuiltIns
}

func (t TypeBuiltIn) Run(writer io.Writer, input []string) error {
	if len(input) < 2 {
		return nil
	}

	for _, el := range input[1:] {
		if _, ok := t.builtIns.Get(el); ok {
			fmt.Fprintf(writer, "%s is a shell builtin\n", el)
			continue
		}

		path, err := exec.LookPath(el)
		if err != nil {
			if errors.Is(err, fs.ErrPermission) {
				continue
			}

			fmt.Fprintf(writer, "%s: not found\n", el)
			continue
		}

		fmt.Fprintf(writer, "%s is %s\n", el, path)
	}

	return nil
}
