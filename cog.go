package disgo

import "fmt"

//type CogBase interface {
//	GetName() string
//	Setup(Bot)
//	CogLoad() error
//	CogUnload() error
//	Commands() map[string]Command
//	GetCommand(string) Command
//	Command(string, CommandHandler) (Command, error)
//}

type Cog struct {
	Name        string
	Bot         Bot
	CommandList []Command
	commands    map[string]Command
}

func (cog Cog) GetName() string { return cog.Name }

func (cog Cog) Setup(bot Bot) {
	cog.Bot = bot
}

func (cog Cog) CogLoad() error { return nil }

func (cog Cog) CogUnload() error { return nil }

func (cog Cog) GetCommand(string) Command { panic("Not implemented") }

func (cog Cog) Commands() map[string]Command { return cog.commands }

func (cog Cog) Command(name string, callback CommandHandler) (*Command, error) {
	println("adding command", name, "to cog", cog.Name)
	if cog.commands == nil {
		cog.commands = make(map[string]Command)
	}

	_, hasKey := cog.commands[name]
	if hasKey {
		return NilCommand, fmt.Errorf("command %s already exists", name)
	}

	command := Command{
		Name:    name,
		Handler: callback,
	}

	cog.commands[name] = command

	cmd, hasKey := cog.commands[name]
	if hasKey {
		println(name, "successfully added to cog command map (", cmd.Name, ")")
	}

	return &command, nil
}
