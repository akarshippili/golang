package invocker

import "github.com/akarshippili/golang/designpatterns/command/command"

type RemoteController struct {
	commands []command.Command
	prev     command.Command // last executed command
}

func GetRemoteController(size int) *RemoteController {
	remote := RemoteController{
		commands: make([]command.Command, size),
	}

	return &remote
}

func (remote *RemoteController) SetCommand(command command.Command, index int) {
	if index < len(remote.commands) {
		(*remote).commands[index] = command
	}
}

func (remote *RemoteController) InvockCommand(index int) {
	if index < len(remote.commands) && remote.commands[index] != nil {
		remote.commands[index].Execute()
		remote.prev = remote.commands[index]
	}
}

func (remote *RemoteController) Undo() {
	if remote.prev != nil {
		remote.prev.Undo()
		remote.prev = nil
	}
}
