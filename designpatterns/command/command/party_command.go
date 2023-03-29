package command

type PartyCommand struct {
	Commands []Command
}

func (partyCommand *PartyCommand) Execute() {
	for _, command := range partyCommand.Commands {
		command.Execute()
	}
}

func (partyCommand *PartyCommand) Undo() {
	for _, command := range partyCommand.Commands {
		command.Undo()
	}
}

func (partyCommand *PartyCommand) AddCommand(command Command) {
	partyCommand.Commands = append(partyCommand.Commands, command)
}
