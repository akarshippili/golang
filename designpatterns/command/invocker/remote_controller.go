package invocker

import "github.com/akarshippili/golang/designpatterns/command/command"

type RemoteController struct {
	commands []command.Command
}

func GetRemoteController(size int) *RemoteController {
	remote := RemoteController{
		commands: make([]command.Command, size),
	}

	return &remote
}
