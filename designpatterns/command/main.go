package main

import (
	"github.com/akarshippili/golang/designpatterns/command/command"
	"github.com/akarshippili/golang/designpatterns/command/invocker"
	"github.com/akarshippili/golang/designpatterns/command/reciver"
)

func main() {

	livingRoomLight := reciver.Light{
		Brightness: 10,
	}

	livingRoomFan := reciver.Fan{
		Speed: 3,
	}

	gerege := reciver.GarageDoor{}

	stereo := reciver.Stereo{
		Volume: 3,
	}

	onlivingRoomLight := command.LightOnCommand{
		Light: livingRoomLight,
	}

	offlivingRoomLight := command.LightOffCommand{
		Light: livingRoomLight,
	}

	onlivingRoomFan := command.FanOnCommand{
		Fan: livingRoomFan,
	}

	offlivingRoomFan := command.FanOffCommand{
		Fan: livingRoomFan,
	}

	openGarageCommand := command.OpenGarageCommand{
		GarageDoor: gerege,
	}

	closeGarageCommand := command.CloseGarageCommand{
		GarageDoor: gerege,
	}

	stereoOnCommand := command.StereoOnCommand{
		Stereo: stereo,
	}

	stereoOffCommand := command.StereoOffCommand{
		Stereo: stereo,
	}

	stereoIncreaseVolumeCommand := command.StereoIncreaseVolumeCommand{
		Stereo: stereo,
	}

	stereoDecreaseVolumeCommand := command.StereoDecreaseVolumeCommand{
		Stereo: stereo,
	}

	partyCommand := command.PartyCommand{
		Commands: make([]command.Command, 0, 10),
	}
	partyCommand.AddCommand(&onlivingRoomLight)
	partyCommand.AddCommand(&onlivingRoomFan)
	partyCommand.AddCommand(&stereoOnCommand)

	remote := invocker.GetRemoteController(11)

	remote.SetCommand(&onlivingRoomLight, 0)
	remote.SetCommand(&offlivingRoomLight, 1)

	remote.SetCommand(&onlivingRoomFan, 2)
	remote.SetCommand(&offlivingRoomFan, 3)

	remote.SetCommand(&openGarageCommand, 4)
	remote.SetCommand(&closeGarageCommand, 5)

	remote.SetCommand(&stereoOnCommand, 6)
	remote.SetCommand(&stereoOffCommand, 7)
	remote.SetCommand(&stereoIncreaseVolumeCommand, 8)
	remote.SetCommand(&stereoDecreaseVolumeCommand, 9)

	remote.SetCommand(&partyCommand, 10)

	for index := 0; index < 11; index++ {
		remote.InvockCommand(index)
		remote.Undo()
	}

}
