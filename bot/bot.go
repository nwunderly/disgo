package main

import (
	"fmt"
	"github.com/nwunderly/disgo/commands"
	"os"
	"strings"
)

func main() {
	bot, err := commands.NewBot(":=", TOKEN)

	if err != nil {
		fmt.Println("encountered error ", err)
		return
	}

	bot.Command("ping", ping)
	bot.Command("echo", echo)
	bot.Command("die", die)

	bot.OnReady(func() {
		user, _ := bot.Me()
		fmt.Println("Logged in as", user)
	})

	bot.CaseInsensitive = true
	//bot.RemoveCommand("help")

	bot.Run()
}

func die(ctx *commands.Context) error {
	if ctx.Author.ID == "204414611578028034" {
		_, _ = ctx.Send("ok")
		_ = ctx.Bot.Session.Close()
		os.Exit(0)
	} else {
		_, _ = ctx.Send("no")
	}
	return nil
}

func echo(ctx *commands.Context) error {
	_, err := ctx.Send(strings.Join(ctx.Args, " "))
	return err
}

func ping(ctx *commands.Context) error {
	_, err := ctx.Send("Pong!")
	return err
}
