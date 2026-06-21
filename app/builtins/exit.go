package builtins

import "os"

type ExitBuiltIn struct {
}

func (t ExitBuiltIn) Run(input []string) error {
	os.Exit(0)
	return nil
}
