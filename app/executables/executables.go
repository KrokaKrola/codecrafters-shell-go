package executables

import (
	"bufio"
	"fmt"
	"os/exec"
)

func RunExecutable(writer *bufio.Writer, args []string) error {
	_, err := exec.LookPath(args[0])
	if err != nil {
		fmt.Fprintf(writer, "%s: command not found\n", args[0])
		writer.Flush()
		return nil
	}

	cmd := exec.Command(args[0], args[1:]...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Fprint(writer, string(output))
	return nil
}
