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

	onlivingRoomLight := command.LightOnCommand{
		Light: livingRoomLight,
	}

	offlivingRoomLight := command.LightOffCommand{
		Light: livingRoomLight,
	}

	remote := invocker.GetRemoteController(10)
	remote.SetCommand(&onlivingRoomLight, 0)
	remote.SetCommand(&offlivingRoomLight, 1)
	remote.InvockCommand(0)
	remote.InvockCommand(1)
	remote.Undo()
}
