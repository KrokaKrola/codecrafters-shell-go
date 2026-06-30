package executables

import (
	"fmt"
	"io"
	"os/exec"
)

func RunExecutable(writer io.Writer, errWriter io.Writer, args []string) error {
	_, err := exec.LookPath(args[0])
	if err != nil {
		fmt.Fprintf(writer, "%s: command not found\n", args[0])
		return nil
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdout = writer
	cmd.Stderr = errWriter

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
