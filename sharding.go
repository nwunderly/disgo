package disgo

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func GetShardCount(token string) (int, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return 0, err
	}
	r, err := session.GatewayBot()
	if err != nil {
		return 0, err
	}
	return r.Shards, nil
}

type Shard struct {
	ID      int
	Session *discordgo.Session
}

type AutoShardedBot struct {
	*Bot
	Shards map[int]*Shard
}

func (bot *AutoShardedBot) AddHandler(handler interface{}) func() {
	removeHandlers := make([]func(), 0, len(bot.Shards))
	for _, shard := range bot.Shards {
		removeHandlers = append(removeHandlers, shard.Session.AddHandler(handler))
	}
	return func() {
		for _, f := range removeHandlers {
			f()
		}
	}
}

func (bot *AutoShardedBot) Run() {
	for _, shard := range bot.Shards {
		fmt.Printf("Starting shard %d\n", shard.ID)
		shard.Session.ShardID = shard.ID
		shard.Session.ShardCount = len(bot.Shards)
		err := shard.Session.Open()
		if err != nil {
			panic(err)
		}
	}
	<-make(chan bool)
}

func (bot *AutoShardedBot) GetShard(guildID string) (*Shard, error) {
	id, err := strconv.Atoi(guildID)
	if err != nil {
		return nil, err
	}
	return bot.Shards[(id>>22)%len(bot.Shards)], nil
}

func (bot *AutoShardedBot) CommandMessageCreateHandler() func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(session *discordgo.Session, event *discordgo.MessageCreate) {
		ctx, valid := bot.GetContext(event)

		if valid {
			recovered, returned := bot.ExecuteSafely(ctx.Invoke)
			if recovered != nil {
				go bot.PanicHandler(ctx, recovered)
				go ctx.Command.PanicHandler(ctx, recovered)
			}
			if returned != nil {
				go bot.ErrorHandler(ctx, returned)
				go ctx.Command.ErrorHandler(ctx, returned)
			}
		}
	}
}

func (bot *AutoShardedBot) GetContext(event *discordgo.MessageCreate) (*Context, bool) {
	shard, err := bot.GetShard(event.GuildID)
	if err != nil {
		// this shouldn't happen
		panic(err)
	}

	ctx, valid := bot.Bot.GetContext(event)
	if !valid {
		return ctx, valid
	}

	ctx.Shard = shard
	ctx.Session = shard.Session

	return ctx, true
}
