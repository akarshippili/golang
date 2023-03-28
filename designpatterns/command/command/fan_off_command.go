package command

import "github.com/akarshippili/golang/designpatterns/command/reciver"

type FanOffCommand struct {
	fan reciver.Fan
}

func (command *FanOffCommand) Execute() {
	command.fan.Off()
}

func (command *FanOffCommand) Undo() {
	command.fan.On()
}
