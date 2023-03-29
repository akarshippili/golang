package command

import "github.com/akarshippili/golang/designpatterns/command/reciver"

type StereoOnCommand struct {
	Stereo reciver.Stereo
}

func (command *StereoOnCommand) Execute() {
	command.Stereo.On()
	command.Stereo.Volume = 3
}

func (command *StereoOnCommand) Undo() {
	command.Stereo.Off()
}
