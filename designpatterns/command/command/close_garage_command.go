package command

import "github.com/akarshippili/golang/designpatterns/command/reciver"

type CloseGarageCommand struct {
	GarageDoor reciver.GarageDoor
}

func (command *CloseGarageCommand) Execute() {
	command.GarageDoor.Close()
}

func (command *CloseGarageCommand) Undo() {
	command.GarageDoor.Open()
}
