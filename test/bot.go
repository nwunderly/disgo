package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nwunderly/disgo"
	"os"
	"strings"
)

func main() {
	bot, err := disgo.NewAutoShardedBot(":=", TOKEN, []int{0, 1}, 2)

	if err != nil {
		fmt.Println("encountered error ", err)
		return
	}

	bot.Command("ping", "", ping)
	bot.Command("echo", "", echo)
	bot.Command("die", "", die)

	bot.AddHandler(
		func(_ *discordgo.Session, ready *discordgo.Ready) {
			fmt.Println("Logged in as", ready.User)
		})

	bot.CaseInsensitive = true
	//bot.RemoveCommand("help")

	bot.Run()
}

func die(ctx *disgo.Context) error {
	if ctx.Author.ID == "204414611578028034" {
		_, _ = ctx.Send("ok")
		_ = ctx.Bot.Session.Close()
		os.Exit(0)
	} else {
		_, _ = ctx.Send("no")
	}
	return nil
}

func echo(ctx *disgo.Context) error {
	_, err := ctx.Send(strings.Join(ctx.Args, " "))
	return err
}

func ping(ctx *disgo.Context) error {
	_, err := ctx.Send("Pong!")
	return err
}
