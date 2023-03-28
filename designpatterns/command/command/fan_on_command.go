package command

import "github.com/akarshippili/golang/designpatterns/command/reciver"

type FanOnCommand struct {
	fan reciver.Fan
}

func (command *FanOnCommand) Execute() {
	command.fan.On()
}

func (command *FanOnCommand) Undo() {
	command.fan.Off()
}
