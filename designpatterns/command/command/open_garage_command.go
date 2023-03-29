package command

import "github.com/akarshippili/golang/designpatterns/command/reciver"

type OpenGarageCommand struct {
	GarageDoor reciver.GarageDoor
}

func (command *OpenGarageCommand) Execute() {
	command.GarageDoor.Open()
}

func (command *OpenGarageCommand) Undo() {
	command.GarageDoor.Close()
}
