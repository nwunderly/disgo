package disgo

import "fmt"

type CommandHandler func(*Context) error
type CommandCheck func(*Context) bool

type CommandInfo struct {
	Name string
	Desc string
}

//type CommandBase interface {
//	Invoke(Context) error
//	Info() CommandInfo
//}

type Command struct {
	Name    string
	Desc    string
	Handler CommandHandler
	Checks []CommandCheck
	Subcommands []*Command
}

var NilCommand = &Command{
	Handler: func(ctx *Context) error { return nil },
}

func (cmd *Command) Invoke(ctx *Context) error {
	for _, check := range cmd.Checks {
		if !check(ctx) {
			return fmt.Errorf("check for command %s failed", cmd.Name)
		}
	}
	return cmd.Handler(ctx)
}

func (cmd *Command) Info() CommandInfo {
	return CommandInfo{Name: cmd.Name, Desc: cmd.Desc}
}

func(cmd *Command) Check(checks ...CommandCheck) {
	for _, check := range checks {
		cmd.Checks = append(cmd.Checks, check)
	}
}

func (cmd *Command) AddSubcommand(command *Command) error {
	_, commandExists := cmd.GetSubcommand(command.Name)
	if commandExists {
		return fmt.Errorf("subcommand %s already exists", command.Name)
	}
	cmd.Subcommands = append(cmd.Subcommands, command)
	return nil
}

func (cmd *Command) GetSubcommand(name string) (*Command, bool) {
	for _, command := range cmd.Subcommands {
		if command == nil {
			continue
		}
		if command.Name == name {
			return command, true
		}
	}
	return NilCommand, false
}

func (bot *Bot) RemoveSubcommand(name string) {
	for i, command := range bot.Commands {
		if command.Name == name {
			bot.Commands = deleteCommand(bot.Commands, i)
		}
	}
}