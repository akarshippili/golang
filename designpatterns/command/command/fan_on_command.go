package command

import "github.com/akarshippili/golang/designpatterns/command/reciver"

type FanOnCommand struct {
	Fan reciver.Fan
}

func (command *FanOnCommand) Execute() {
	command.Fan.On()
}

func (command *FanOnCommand) Undo() {
	command.Fan.Off()
}
