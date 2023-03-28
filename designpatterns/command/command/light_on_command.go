package command

import "github.com/akarshippili/golang/designpatterns/command/reciver"

type LightOnCommand struct {
	Light reciver.Light
}

func (command *LightOnCommand) Execute() {
	command.Light.On()
}

func (command *LightOnCommand) Undo() {
	command.Light.Off()
}
