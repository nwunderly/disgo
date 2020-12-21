package testcog

import (
	"github.com/nwunderly/disgo/commands"
)

type TestCog struct {
	commands.Cog
}
//
//func (cog TestCog) Setup(bot commands.Bot) {
//	cog.Bot = bot
//	cog.Name = "TestCog"
//	_, err := cog.Command("testcogmethodcommand", cog.TestCogMethodCommand)
//	if err != nil {
//		println("Error adding testcogmethodcommand")
//		panic(err)
//	}
//	_, err = cog.Command("testcogfunctioncommand", TestCogFunctionCommand)
//	if err != nil {
//		println("Error adding testcogfunctioncommand")
//		panic(err)
//	}
//	println(len(cog.Commands()), "commands detected in cog by TestCog.Setup")
//}
//
//func (cog TestCog) CogLoad() error {
//	fmt.Println("Loaded cog TestCog")
//	return nil
//}
//
//func (cog TestCog) TestCogMethodCommand(ctx commands.Context) error {
//	_, err := ctx.Send("Works")
//	return err
//}
//
//func TestCogFunctionCommand(ctx commands.Context) error {
//	_, err := ctx.Send("Works")
//	return err
//}
