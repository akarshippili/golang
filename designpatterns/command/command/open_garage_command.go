package command

import "github.com/akarshippili/golang/designpatterns/command/reciver"

type OpenGarageCommand struct {
	garageDoor reciver.GarageDoor
}

func (command *OpenGarageCommand) Execute() {
	command.garageDoor.Open()
}

func (command *OpenGarageCommand) Undo() {
	command.garageDoor.Close()
}
