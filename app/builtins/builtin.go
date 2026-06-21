package builtins

import (
	"bufio"
)

type BuiltIn interface {
	Run([]string) error
}

type BuiltInCommand string

const (
	exitCommand string = "exit"
	echoCommand string = "echo"
	typeCommand string = "type"
)

type BuiltInsMap map[string]BuiltIn

func NewBuiltIns(writer *bufio.Writer) BuiltInsMap {
	result := make(map[string]BuiltIn)

	result[exitCommand] = ExitBuiltIn{}
	result[echoCommand] = EchoBuiltIn{writer: writer}

	// always last
	result[typeCommand] = TypeBuiltIn{writer: writer, builtIns: result}

	return result
}
