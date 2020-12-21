package main

import (
	"fmt"
	"github.com/nwunderly/disgo/commands"
	"strings"
)

func main() {
	bot, err := commands.NewBot(":=", TOKEN)

	if err != nil {
		fmt.Println("encountered error ", err)
		return
	}

	_, _ = bot.Command("ping",
		func(ctx *commands.Context) error {
			_, err = ctx.Send("Pong!")
			return err
		})

	_, _ = bot.Command("echo",
		func(ctx *commands.Context) error {
			_, err = ctx.Send(strings.Join(ctx.Args, " "))
			return err
		})

	if err != nil {
		fmt.Println("encountered error", err)
		return
	}

	bot.CaseInsensitive = true
	//bot.RemoveCommand("help")

	user, _ := bot.Me()
	fmt.Println("Logged in as", user)
	bot.Run()
}
