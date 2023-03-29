package command

import "github.com/akarshippili/golang/designpatterns/command/reciver"

type StereoOffCommand struct {
	Stereo reciver.Stereo
}

func (command *StereoOffCommand) Execute() {
	command.Stereo.Off()
}

func (command *StereoOffCommand) Undo() {
	command.Stereo.On()
	command.Stereo.Volume = 3
}
