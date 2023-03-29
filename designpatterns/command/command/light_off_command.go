package command

import "github.com/akarshippili/golang/designpatterns/command/reciver"

type LightOffCommand struct {
	Light reciver.Light
}

func (command *LightOffCommand) Execute() {
	command.Light.Off()
}

func (command *LightOffCommand) Undo() {
	command.Light.On()
}
