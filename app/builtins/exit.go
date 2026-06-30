package builtins

import (
	"io"
	"os"
	"strconv"
)

type ExitBuiltIn struct {
}

func (t ExitBuiltIn) Run(writer io.Writer, input []string) error {
	if len(input) > 1 {
		statusCodeArg := input[1]

		value, err := strconv.ParseInt(statusCodeArg, 10, 32)
		if err != nil {
			return err
		}

		os.Exit(int(value))
		return nil
	}

	os.Exit(0)
	return nil
}
