package builtins

import (
	"io"
)

type BuiltIn interface {
	Run(io.Writer, []string) error
}

const (
	exitCommand string = "exit"
	echoCommand string = "echo"
	typeCommand string = "type"
	pwdCommand  string = "pwd"
	cdCommand   string = "cd"
)

type BuiltIns struct {
	commands map[string]BuiltIn
}

func (b *BuiltIns) add(key string, value BuiltIn) {
	b.commands[key] = value
}

func (b *BuiltIns) Get(key string) (BuiltIn, bool) {
	v, ok := b.commands[key]
	return v, ok
}

func NewBuiltIns() *BuiltIns {
	result := &BuiltIns{
		commands: make(map[string]BuiltIn),
	}

	result.add(echoCommand, EchoBuiltIn{})
	result.add(exitCommand, ExitBuiltIn{})
	result.add(pwdCommand, PwdBuiltIn{})
	result.add(cdCommand, CdBuiltIn{})
	result.add(typeCommand, TypeBuiltIn{builtIns: result})

	return result
}
