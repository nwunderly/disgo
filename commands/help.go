package commands

import (
	"sort"
	"strings"
)

type HelpCommand interface {
	//Invoke(Context) error
	SendBotHelp(ctx Context) error
	SendCommandHelp(ctx Context, commandName string) error
}

//func NilHelpCommand() HelpCommand {
//	return HelpCommand(nil)
//}

type helpCommandRunner struct {
	bot Bot
}

func (runner helpCommandRunner) Invoke(ctx Context) error {
	return runner.bot.HelpCommand.SendBotHelp(ctx)
}

func NewDefaultHelpCommand() DefaultHelpCommand {
	var help DefaultHelpCommand
	//help.Name = "help"
	//help.InvokeCallback = func(ctx Context) error {return help.Invoke(ctx)}

	return help
}

type DefaultHelpCommand struct {
	//Name string
	//InvokeCallback commandInvokeCallback
}

func (help DefaultHelpCommand) SendBotHelp(ctx Context) error {
	var commandNames []string
	for name, _ := range ctx.Bot.Commands {
		commandNames = append(commandNames, name)
	}
	sort.Strings(commandNames)
	_, err := ctx.Send("Commands:\n - " + strings.Join(commandNames, "\n - "))
	return err
}

func (help DefaultHelpCommand) SendCommandHelp(ctx Context, commandName string) error {
	_, err := ctx.Send("Command Help: " + commandName)
	return err
}
