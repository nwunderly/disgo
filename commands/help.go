package commands

import (
	"sort"
	"strings"
)

type HelpCommand interface {
	//Invoke(Context) error
	SendHelp(ctx *Context) error
	SendBotHelp(ctx *Context) error
	SendCommandHelp(ctx *Context, commandName string) error
	Info() CommandInfo
}

//type helpCommandRunner struct {
//	bot  Bot
//	help HelpCommand
//}
//
//func (runner helpCommandRunner) Info() CommandInfo {
//	return runner.help.Info()
//}
//
//func (runner helpCommandRunner) Invoke(ctx Context) error {
//	return runner.help.SendBotHelp(ctx)
//}

func HelpCommandHandler(help HelpCommand) *Command {
	info := help.Info()

	command := Command{
		Name: info.Name,
		Desc: info.Desc,
		Handler: func(ctx *Context) error {
			return help.SendHelp(ctx)
		},
	}
	return &command
}

func NewDefaultHelpCommand() DefaultHelpCommand {
	var help DefaultHelpCommand
	help.Name = "help"
	help.Desc = "Shows this page."

	return help
}

type DefaultHelpCommand struct {
	Name string
	Desc string
	//Handler CommandHandler
}

func (help DefaultHelpCommand) Info() CommandInfo {
	return CommandInfo{Name: help.Name, Desc: help.Desc}
}

func (help DefaultHelpCommand) SendHelp(ctx *Context) error {
	return help.SendBotHelp(ctx)
}

func (help DefaultHelpCommand) SendBotHelp(ctx *Context) error {
	var commandNames []string
	for _, command := range ctx.Bot.Commands {
		info := command.Info()
		commandNames = append(commandNames, info.Name+": "+info.Desc)
	}
	sort.Strings(commandNames)
	_, err := ctx.Send("Commands:\n - " + strings.Join(commandNames, "\n - "))
	return err
}

func (help DefaultHelpCommand) SendCommandHelp(ctx *Context, commandName string) error {
	_, err := ctx.Send("Command Help: " + commandName)
	return err
}
