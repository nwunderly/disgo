package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

type Bot struct {
	Prefix          string
	CaseInsensitive bool
	Session         *discordgo.Session
	Commands        []*Command
	Cogs            []*Cog
	HelpCommand     *HelpCommand
}

func NewBot(prefix, token string) (*Bot, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	commands := make([]*Command, 0, 1)
	cogs := make([]*Cog, 0)

	bot := Bot{
		Prefix:   prefix,
		Session:  session,
		Commands: commands,
		Cogs:     cogs,
	}

	session.AddHandler(bot.CommandMessageCreateHandler())

	err = bot.SetHelpCommand(NewDefaultHelpCommand())
	if err != nil {
		return nil, err
	}

	return &bot, nil
}

func (bot *Bot) CommandMessageCreateHandler() func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(session *discordgo.Session, event *discordgo.MessageCreate) {
		ctx, valid := bot.GetContext(event)
		if valid {
			defer RecoverAndLog()
			_ = ctx.Invoke()
		}
	}
}

func (bot *Bot) Command(name string, desc string, callback CommandHandler) (*Command, error) {

	command := &Command{
		Name:    name,
		Desc:    desc,
		Handler: callback,
	}
	err := bot.AddCommand(command)

	return command, err
}

func (bot *Bot) AddCommand(command *Command) error {
	_, commandExists := bot.GetCommand(command.Name)
	if commandExists {
		return fmt.Errorf("command %s already exists", command.Name)
	}
	bot.Commands = append(bot.Commands, command)
	return nil
}

func (bot *Bot) GetCommand(name string) (*Command, bool) {
	for _, command := range bot.Commands {
		if command == nil {
			continue
		}
		if command.Name == name {
			return command, true
		}
	}
	return NilCommand, false
}

func (bot *Bot) RemoveCommand(name string) {
	for i, command := range bot.Commands {
		if command.Name == name {
			bot.Commands = deleteCommand(bot.Commands, i)
		}
	}
}

func (bot *Bot) SetHelpCommand(help HelpCommand) error {
	return bot.AddCommand(HelpCommandHandler(help))
}

func (bot *Bot) GetContext(event *discordgo.MessageCreate) (*Context, bool) {
	message := event.Message
	session := bot.Session
	content := message.Content

	if message.Author.Bot || message.Author.ID == session.State.User.ID {
		return NilContext, false
	}

	// check for command prefix in message. if present, trim it/redefine content
	contentHasPrefix, content := hasPrefix(content, bot.Prefix, false)
	if !contentHasPrefix {
		return NilContext, false
	}

	if content == "" {
		return NilContext, false
	}

	split := GetArgs(content)

	command, validCommand := bot.GetCommand(split[0])
	args := split[1:]

	if !validCommand {
		return NilContext, false
	}

	guild, errGuild := bot.Session.Guild(message.GuildID)
	channel, errChannel := bot.Session.Channel(message.ChannelID)

	if errGuild != nil || errChannel != nil {
		return NilContext, false
	}

	return NewContext(bot, command, message.Author, message.Member, channel, guild, args), true

}

//func (bot Bot) LoadCog(cog CogBase) {
//	defer func() {
//		err := recover()
//		if err != nil {
//			fmt.Printf("Error loading cog %s: %s\n", cog.GetName(), err)
//		}
//	}()
//
//	println("Setting up cog", cog.GetName())
//	println(bot.Cogs)
//	cog.Setup(bot)
//
//	println("Cog set up")
//	bot.Cogs[cog.GetName()] = cog
//
//	println("adding commands")
//	for name, cmd := range cog.Commands() {
//		println("adding command", name)
//		err := bot.AddCommand(cmd)
//		if err != nil {
//			panic(err)
//		}
//	}
//
//	go ExecuteSafely(cog.CogLoad)
//}

func (bot *Bot) Run() {
	session := bot.Session

	err := session.Open()
	if err != nil {
		fmt.Println("error opening connection:", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	_ = session.Close()
}

func (bot *Bot) Me() (*discordgo.User, error) {
	return bot.Session.State.User, nil
	//return bot.Session.User("@me")
}

func (bot *Bot) WaitForMessage(check func(*discordgo.MessageCreate) bool) *discordgo.MessageCreate {
	result := make(chan *discordgo.MessageCreate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.MessageCreate) {
			if check(msg) {
				// check if other channel is safe to write to
				_, ok := <-waiting
				if !ok {
					return
				}
				// send message to waiting goroutine
				result <- msg
				close(result)
			}
		})

	// block until a message is found and sent
	msg := <-result
	closeHandler()
	close(waiting)
	return msg
}
