package command

import "github.com/akarshippili/golang/designpatterns/command/reciver"

type StereoIncreaseVolumeCommand struct {
	Stereo reciver.Stereo
}

func (command *StereoIncreaseVolumeCommand) Execute() {
	command.Stereo.IncreaseVolume()
}

func (command *StereoIncreaseVolumeCommand) Undo() {
	command.Stereo.DecreaseVolume()
}
