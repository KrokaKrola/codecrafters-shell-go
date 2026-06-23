package builtins

import (
	"bufio"
)

type BuiltIn interface {
	Run([]string) error
}

const (
	exitCommand string = "exit"
	echoCommand string = "echo"
	typeCommand string = "type"
	pwdCommand  string = "pwd"
)

type BuiltInsMap map[string]BuiltIn

func NewBuiltIns(writer *bufio.Writer) BuiltInsMap {
	result := make(BuiltInsMap)

	result[exitCommand] = ExitBuiltIn{}
	result[echoCommand] = EchoBuiltIn{writer: writer}
	result[typeCommand] = TypeBuiltIn{writer: writer, builtIns: result}
	result[pwdCommand] = PwdBuiltIn{writer: writer}

	return result
}
