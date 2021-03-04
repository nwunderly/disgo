package disgo

import "github.com/bwmarrin/discordgo"

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

	bot.ErrorHandler = bot.defaultErrorHandler
	bot.PanicHandler = bot.defaultPanicHandler

	err = bot.SetHelpCommand(NewDefaultHelpCommand())
	if err != nil {
		return nil, err
	}

	return &bot, nil
}

func NewAutoShardedBot(prefix, token string, shardIDs []int, shardCount int) (*AutoShardedBot, error) {
	var err error
	var session *discordgo.Session

	if len(shardIDs) == 0 {
		if shardCount == 0 {
			shardCount, err = GetShardCount(token)
			if err != nil {
				return nil, err
			}
		}
		shardIDs = intRange(0, shardCount)
	} else if shardCount == 0 {
		shardCount, err = GetShardCount(token)
		if err != nil {
			return nil, err
		}
	}

	shards := make(map[int]*Shard)
	for _, id := range shardIDs {
		session, err = discordgo.New("Bot " + token)
		if err != nil {
			return nil, err
		}
		shard := &Shard{
			ID:      id,
			Session: session,
		}
		shards[id] = shard
	}

	commands := make([]*Command, 0, 1)
	cogs := make([]*Cog, 0)

	bot := AutoShardedBot{
		Bot: &Bot{
			Prefix:   prefix,
			Commands: commands,
			Cogs:     cogs,
		},
		Shards: shards,
	}

	bot.AddHandler(bot.CommandMessageCreateHandler())

	bot.ErrorHandler = bot.defaultErrorHandler
	bot.PanicHandler = bot.defaultPanicHandler

	err = bot.SetHelpCommand(NewDefaultHelpCommand())
	if err != nil {
		return nil, err
	}

	return &bot, nil
}
