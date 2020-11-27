package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Bot struct {
	Prefix          string
	CaseInsensitive bool
	Session         *discordgo.Session
	Commands        map[string]CommandBase
	Cogs            map[string]CogBase
	//HelpCommand HelpCommand
}

func NewBot(prefix, token string) (*Bot, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	bot := Bot{
		Prefix:   prefix,
		Session:  session,
		Commands: make(map[string]CommandBase),
		Cogs:     make(map[string]CogBase),
	}

	session.AddHandler(bot.MessageCreateHandler())
	bot.SetHelpCommand(NewDefaultHelpCommand())

	return &bot, nil
}

func (bot Bot) MessageCreateHandler() func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(session *discordgo.Session, event *discordgo.MessageCreate) {
		ctx, valid := bot.GetContext(event)
		if valid {
			defer RecoverAndLog()
			_ = ctx.Invoke()
		}
	}
}

func (bot Bot) Command(name string, callback CommandCallback) (Command, error) {
	_, hasKey := bot.Commands[name]
	if hasKey {
		return Command{}, fmt.Errorf("command %s already exists", name)
	}

	command := Command{
		Name:           name,
		InvokeCallback: callback,
	}

	bot.Commands[name] = command

	return command, nil
}

func (bot Bot) AddCommand(command Command) error {
	_, commandExists := bot.GetCommand(command.Name)
	if commandExists {
		return fmt.Errorf("command %s already exists", command.Name)
	}
	bot.Commands[command.Name] = command
	return nil
}

func (bot Bot) GetCommand(name string) (CommandBase, bool) {

	command, hasKey := bot.Commands[name]

	if hasKey {
		return command, true
	}
	return Command{}, false
}

func (bot Bot) RemoveCommand(name string) {
	delete(bot.Commands, name)
}

func (bot Bot) SetHelpCommand(help HelpCommand) {
	//bot.HelpCommand = help
	info := help.Info()
	bot.Commands[info.Name] = helpCommandRunner{bot, help}
}

func (bot *Bot) GetContext(event *discordgo.MessageCreate) (Context, bool) {
	message := event.Message
	session := bot.Session
	content := message.Content

	if message.Author.Bot || message.Author.ID == session.State.User.ID {
		return Context{}, false
	}

	// check for command prefix in message. if present, trim it/redefine content
	contentHasPrefix, content := hasPrefix(content, bot.Prefix, false)
	if !contentHasPrefix {
		return Context{}, false
	}

	if content == "" {
		return Context{}, false
	}

	split := strings.Split(content, " ")

	command, validCommand := bot.GetCommand(split[0])

	if !validCommand {
		return Context{}, false
	}

	guild, errGuild := bot.Session.Guild(message.GuildID)
	channel, errChannel := bot.Session.Channel(message.ChannelID)

	if errGuild != nil || errChannel != nil {
		return Context{}, false
	}

	return NewContext(bot, command, message.Author, message.Member, channel, guild), true

}

func (bot Bot) LoadCog(cog CogBase) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("Error loading cog %s: %s\n", cog.GetName(), err)
		}
	}()

	println("Setting up cog", cog.GetName())
	println(bot.Cogs)
	cog.Setup(bot)

	println("Cog set up")
	bot.Cogs[cog.GetName()] = cog

	println("adding commands")
	for name, cmd := range cog.Commands() {
		println("adding command", name)
		err := bot.AddCommand(cmd)
		if err != nil {
			panic(err)
		}
	}

	go ExecuteSafely(cog.CogLoad)
}

func (bot *Bot) Run() {
	session := bot.Session

	err := session.Open()
	if err != nil {
		fmt.Println("error opening connection:", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	_ = session.Close()
}
