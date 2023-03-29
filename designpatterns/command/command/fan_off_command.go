package command

import "github.com/akarshippili/golang/designpatterns/command/reciver"

type FanOffCommand struct {
	Fan reciver.Fan
}

func (command *FanOffCommand) Execute() {
	command.Fan.Off()
}

func (command *FanOffCommand) Undo() {
	command.Fan.On()
}
