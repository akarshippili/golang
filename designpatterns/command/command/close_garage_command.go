package command

import "github.com/akarshippili/golang/designpatterns/command/reciver"

type CloseGarageCommand struct {
	garageDoor reciver.GarageDoor
}

func (command *CloseGarageCommand) Execute() {
	command.garageDoor.Close()
}

func (command *CloseGarageCommand) Undo() {
	command.garageDoor.Open()
}
