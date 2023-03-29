package command

import "github.com/akarshippili/golang/designpatterns/command/reciver"

type StereoDecreaseVolumeCommand struct {
	Stereo reciver.Stereo
}

func (command *StereoDecreaseVolumeCommand) Execute() {
	command.Stereo.DecreaseVolume()
}

func (command *StereoDecreaseVolumeCommand) Undo() {
	command.Stereo.IncreaseVolume()
}
