package builtins

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type CdBuiltIn struct {
}

func (c CdBuiltIn) Run(writer io.Writer, args []string) error {
	var path string

	if len(args) == 1 || args[1] == "~" || strings.HasPrefix(args[1], "~/") {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		path = filepath.Join(homedir, strings.TrimPrefix(args[1], "~"))
	} else {
		path = args[1]
	}

	if err := os.Chdir(path); err != nil {
		return fmt.Errorf("cd: %s: No such file or directory", path)
	}

	return nil
}
