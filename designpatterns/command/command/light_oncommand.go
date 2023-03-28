package command

import "github.com/akarshippili/golang/designpatterns/command/reciver"

type LightOnCommand struct {
	light reciver.Light
}

func (command *LightOnCommand) Execute() {
	command.light.On()
}

func (command *LightOnCommand) Undo() {
	command.light.Off()
}
