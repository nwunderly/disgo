package main

import (
	"fmt"
	"github.com/nwunderly/disgo/bot/testcog"
	"github.com/nwunderly/disgo/commands"
)

func main() {
	bot, err := commands.NewBot(":=", TOKEN)

	if err != nil {
		fmt.Println("encountered error ", err)
		return
	}

	_, err = bot.Command("ping",
		func(ctx commands.Context) error {
			_, err = ctx.Send("Pong!")
			return err
		})

	bot.CaseInsensitive = true
	//bot.RemoveCommand("help")

	bot.LoadCog(testcog.TestCog{})

	bot.Run()
}
