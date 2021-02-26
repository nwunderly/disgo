package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nwunderly/disgo"
	"os"
	"strings"
)

func main() {
	bot, err := disgo.NewBot(":=", TOKEN)

	if err != nil {
		fmt.Println("encountered error ", err)
		return
	}

	bot.Command("ping", "", ping)
	bot.Command("echo", "", echo)
	bot.Command("die", "", die)
	bot.Command("testwaitfor", "", testWaitFor)
	bot.Command("getargs", "", getArgs)

	cmd, _ := bot.Command("testsubcommand", "", testSubcommandMain)
	_, err = cmd.Subcommand("subcommand", "", testSubcommandSub)
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd, _ = bot.Command("testchecks", "", testChecks)
	cmd.Check(func(ctx *disgo.Context) bool {
		check := ctx.Author.ID == "204414611578028034"
		ctx.Send(fmt.Sprintf("check result: %t", check))
		return check
	})

	cmd, _ = bot.Command("testerror", "", testError)
	cmd.ErrorHandler = func(ctx *disgo.Context, err error) {
		fmt.Println("TESTERROR WORKS")
	}

	cmd, _ = bot.Command("testpanic", "", testPanic)
	cmd.PanicHandler = func(ctx *disgo.Context, i interface{}) {
		fmt.Println("TESTPANIC WORKS")
	}

	bot.Session.AddHandler(
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

func testWaitFor(ctx *disgo.Context) error {
	_, _ = ctx.Send("Send a message.")
	msg := ctx.Bot.WaitForMessageCreate(
		func(m *discordgo.MessageCreate) bool {
			return m.Author.ID == ctx.Author.ID && m.Content == ctx.Args[0]
		})
	_, _ = ctx.Send(msg.ContentWithMentionsReplaced())
	return nil
}

func getArgs(ctx *disgo.Context) error {
	_, err := ctx.Send(strings.Join(ctx.Args, ", "))
	return err
}

func testSubcommandMain(ctx *disgo.Context) error {
	_, err := ctx.Send("No subcommand given.")
	return err
}

func testSubcommandSub(ctx *disgo.Context) error {
	_, err := ctx.Send("Subcommand worked, args:" + strings.Join(ctx.Args, ", "))
	return err
}

func testChecks(ctx *disgo.Context) error {
	_, err := ctx.Send("works")
	return err
}

func testError(ctx *disgo.Context) error {
	return fmt.Errorf("test")
}

func testPanic(ctx *disgo.Context) error {
	panic("test")
}
