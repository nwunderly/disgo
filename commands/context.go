package commands

import "github.com/bwmarrin/discordgo"

type Context struct {
	Bot     *Bot
	Command *Command

	Session *discordgo.Session
	Author  *discordgo.User
	Member  *discordgo.Member
	Channel *discordgo.Channel
	Guild   *discordgo.Guild
}

var NilContext = &Context{
	Command: NilCommand,
}

func NewContext(bot *Bot, command *Command, author *discordgo.User, member *discordgo.Member,
	channel *discordgo.Channel, guild *discordgo.Guild) *Context {
	return &Context{
		Bot:     bot,
		Command: command,
		Session: bot.Session,
		Author:  author,
		Member:  member,
		Channel: channel,
		Guild:   guild,
	}
}

func (ctx *Context) Invoke() error {
	return ctx.Command.Invoke(ctx)
}

func (ctx *Context) Send(content string) (*discordgo.Message, error) {
	return ctx.Session.ChannelMessageSend(ctx.Channel.ID, content)
}
