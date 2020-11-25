package main

import (
	"fmt"
	"github.com/nwunderly/disgo/commands"
)

func main() {
	bot, err := commands.NewBot("!", TOKEN)

	if err != nil {
		fmt.Println("encountered error ")
		return
	}

	bot.Command("ping",
		func(ctx commands.Context) error {
			_, err = ctx.Send("Pong!")
			return err
		})

	bot.Run()
}
