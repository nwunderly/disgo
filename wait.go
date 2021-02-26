package disgo

import "github.com/bwmarrin/discordgo"

func (bot *Bot) waitFor(check func(interface{}) bool) interface{} {
	result := make(chan interface{})
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, evt interface{}) {
			if check(evt) {
				// check if other channel is safe to write to
				_, ok := <-waiting
				if !ok {
					return
				}
				// send message to waiting goroutine
				result <- evt
				close(result)
			}
		})

	// block until a message is found and sent
	msg := <-result
	closeHandler()
	close(waiting)
	return msg
}

func (bot *Bot) WaitForChannelCreate(check func(*discordgo.ChannelCreate) bool) *discordgo.ChannelCreate {
	result := make(chan *discordgo.ChannelCreate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.ChannelCreate) {
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

func (bot *Bot) WaitForChannelDelete(check func(*discordgo.ChannelDelete) bool) *discordgo.ChannelDelete {
	result := make(chan *discordgo.ChannelDelete)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.ChannelDelete) {
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

func (bot *Bot) WaitForChannelPinsUpdate(check func(*discordgo.ChannelPinsUpdate) bool) *discordgo.ChannelPinsUpdate {
	result := make(chan *discordgo.ChannelPinsUpdate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.ChannelPinsUpdate) {
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

func (bot *Bot) WaitForChannelUpdate(check func(*discordgo.ChannelUpdate) bool) *discordgo.ChannelUpdate {
	result := make(chan *discordgo.ChannelUpdate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.ChannelUpdate) {
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

func (bot *Bot) WaitForConnect(check func(*discordgo.Connect) bool) *discordgo.Connect {
	result := make(chan *discordgo.Connect)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.Connect) {
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

func (bot *Bot) WaitForDisconnect(check func(*discordgo.Disconnect) bool) *discordgo.Disconnect {
	result := make(chan *discordgo.Disconnect)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.Disconnect) {
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

func (bot *Bot) WaitForEvent(check func(*discordgo.Event) bool) *discordgo.Event {
	result := make(chan *discordgo.Event)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.Event) {
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

func (bot *Bot) WaitForGuildBanAdd(check func(*discordgo.GuildBanAdd) bool) *discordgo.GuildBanAdd {
	result := make(chan *discordgo.GuildBanAdd)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.GuildBanAdd) {
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

func (bot *Bot) WaitForGuildBanRemove(check func(*discordgo.GuildBanRemove) bool) *discordgo.GuildBanRemove {
	result := make(chan *discordgo.GuildBanRemove)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.GuildBanRemove) {
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

func (bot *Bot) WaitForGuildCreate(check func(*discordgo.GuildCreate) bool) *discordgo.GuildCreate {
	result := make(chan *discordgo.GuildCreate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.GuildCreate) {
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

func (bot *Bot) WaitForGuildDelete(check func(*discordgo.GuildDelete) bool) *discordgo.GuildDelete {
	result := make(chan *discordgo.GuildDelete)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.GuildDelete) {
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

func (bot *Bot) WaitForGuildEmojisUpdate(check func(*discordgo.GuildEmojisUpdate) bool) *discordgo.GuildEmojisUpdate {
	result := make(chan *discordgo.GuildEmojisUpdate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.GuildEmojisUpdate) {
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

func (bot *Bot) WaitForGuildIntegrationsUpdate(check func(*discordgo.GuildIntegrationsUpdate) bool) *discordgo.GuildIntegrationsUpdate {
	result := make(chan *discordgo.GuildIntegrationsUpdate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.GuildIntegrationsUpdate) {
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

func (bot *Bot) WaitForGuildMemberAdd(check func(*discordgo.GuildMemberAdd) bool) *discordgo.GuildMemberAdd {
	result := make(chan *discordgo.GuildMemberAdd)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.GuildMemberAdd) {
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

func (bot *Bot) WaitForGuildMemberRemove(check func(*discordgo.GuildMemberRemove) bool) *discordgo.GuildMemberRemove {
	result := make(chan *discordgo.GuildMemberRemove)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.GuildMemberRemove) {
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

func (bot *Bot) WaitForGuildMemberUpdate(check func(*discordgo.GuildMemberUpdate) bool) *discordgo.GuildMemberUpdate {
	result := make(chan *discordgo.GuildMemberUpdate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.GuildMemberUpdate) {
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

func (bot *Bot) WaitForGuildMembersChunk(check func(*discordgo.GuildMembersChunk) bool) *discordgo.GuildMembersChunk {
	result := make(chan *discordgo.GuildMembersChunk)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.GuildMembersChunk) {
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

func (bot *Bot) WaitForGuildRoleCreate(check func(*discordgo.GuildRoleCreate) bool) *discordgo.GuildRoleCreate {
	result := make(chan *discordgo.GuildRoleCreate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.GuildRoleCreate) {
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

func (bot *Bot) WaitForGuildRoleDelete(check func(*discordgo.GuildRoleDelete) bool) *discordgo.GuildRoleDelete {
	result := make(chan *discordgo.GuildRoleDelete)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.GuildRoleDelete) {
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

func (bot *Bot) WaitForGuildRoleUpdate(check func(*discordgo.GuildRoleUpdate) bool) *discordgo.GuildRoleUpdate {
	result := make(chan *discordgo.GuildRoleUpdate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.GuildRoleUpdate) {
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

func (bot *Bot) WaitForGuildUpdate(check func(*discordgo.GuildUpdate) bool) *discordgo.GuildUpdate {
	result := make(chan *discordgo.GuildUpdate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.GuildUpdate) {
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

func (bot *Bot) WaitForMessageAck(check func(*discordgo.MessageAck) bool) *discordgo.MessageAck {
	result := make(chan *discordgo.MessageAck)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.MessageAck) {
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

func (bot *Bot) WaitForMessageCreate(check func(*discordgo.MessageCreate) bool) *discordgo.MessageCreate {
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

func (bot *Bot) WaitForMessageDelete(check func(*discordgo.MessageDelete) bool) *discordgo.MessageDelete {
	result := make(chan *discordgo.MessageDelete)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.MessageDelete) {
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

func (bot *Bot) WaitForMessageDeleteBulk(check func(*discordgo.MessageDeleteBulk) bool) *discordgo.MessageDeleteBulk {
	result := make(chan *discordgo.MessageDeleteBulk)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.MessageDeleteBulk) {
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

func (bot *Bot) WaitForMessageReactionAdd(check func(*discordgo.MessageReactionAdd) bool) *discordgo.MessageReactionAdd {
	result := make(chan *discordgo.MessageReactionAdd)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.MessageReactionAdd) {
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

func (bot *Bot) WaitForMessageReactionRemove(check func(*discordgo.MessageReactionRemove) bool) *discordgo.MessageReactionRemove {
	result := make(chan *discordgo.MessageReactionRemove)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.MessageReactionRemove) {
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

func (bot *Bot) WaitForMessageReactionRemoveAll(check func(*discordgo.MessageReactionRemoveAll) bool) *discordgo.MessageReactionRemoveAll {
	result := make(chan *discordgo.MessageReactionRemoveAll)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.MessageReactionRemoveAll) {
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

func (bot *Bot) WaitForMessageUpdate(check func(*discordgo.MessageUpdate) bool) *discordgo.MessageUpdate {
	result := make(chan *discordgo.MessageUpdate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.MessageUpdate) {
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

func (bot *Bot) WaitForPresenceUpdate(check func(*discordgo.PresenceUpdate) bool) *discordgo.PresenceUpdate {
	result := make(chan *discordgo.PresenceUpdate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.PresenceUpdate) {
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

func (bot *Bot) WaitForPresencesReplace(check func(*discordgo.PresencesReplace) bool) *discordgo.PresencesReplace {
	result := make(chan *discordgo.PresencesReplace)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.PresencesReplace) {
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

func (bot *Bot) WaitForRateLimit(check func(*discordgo.RateLimit) bool) *discordgo.RateLimit {
	result := make(chan *discordgo.RateLimit)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.RateLimit) {
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

func (bot *Bot) WaitForReady(check func(*discordgo.Ready) bool) *discordgo.Ready {
	result := make(chan *discordgo.Ready)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.Ready) {
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

func (bot *Bot) WaitForRelationshipAdd(check func(*discordgo.RelationshipAdd) bool) *discordgo.RelationshipAdd {
	result := make(chan *discordgo.RelationshipAdd)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.RelationshipAdd) {
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

func (bot *Bot) WaitForRelationshipRemove(check func(*discordgo.RelationshipRemove) bool) *discordgo.RelationshipRemove {
	result := make(chan *discordgo.RelationshipRemove)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.RelationshipRemove) {
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

func (bot *Bot) WaitForResumed(check func(*discordgo.Resumed) bool) *discordgo.Resumed {
	result := make(chan *discordgo.Resumed)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.Resumed) {
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

func (bot *Bot) WaitForTypingStart(check func(*discordgo.TypingStart) bool) *discordgo.TypingStart {
	result := make(chan *discordgo.TypingStart)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.TypingStart) {
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

func (bot *Bot) WaitForUserGuildSettingsUpdate(check func(*discordgo.UserGuildSettingsUpdate) bool) *discordgo.UserGuildSettingsUpdate {
	result := make(chan *discordgo.UserGuildSettingsUpdate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.UserGuildSettingsUpdate) {
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

func (bot *Bot) WaitForUserNoteUpdate(check func(*discordgo.UserNoteUpdate) bool) *discordgo.UserNoteUpdate {
	result := make(chan *discordgo.UserNoteUpdate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.UserNoteUpdate) {
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

func (bot *Bot) WaitForUserSettingsUpdate(check func(*discordgo.UserSettingsUpdate) bool) *discordgo.UserSettingsUpdate {
	result := make(chan *discordgo.UserSettingsUpdate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.UserSettingsUpdate) {
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

func (bot *Bot) WaitForUserUpdate(check func(*discordgo.UserUpdate) bool) *discordgo.UserUpdate {
	result := make(chan *discordgo.UserUpdate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.UserUpdate) {
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

func (bot *Bot) WaitForVoiceServerUpdate(check func(*discordgo.VoiceServerUpdate) bool) *discordgo.VoiceServerUpdate {
	result := make(chan *discordgo.VoiceServerUpdate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.VoiceServerUpdate) {
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

func (bot *Bot) WaitForVoiceStateUpdate(check func(*discordgo.VoiceStateUpdate) bool) *discordgo.VoiceStateUpdate {
	result := make(chan *discordgo.VoiceStateUpdate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.VoiceStateUpdate) {
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

func (bot *Bot) WaitForWebhooksUpdate(check func(*discordgo.WebhooksUpdate) bool) *discordgo.WebhooksUpdate {
	result := make(chan *discordgo.WebhooksUpdate)
	waiting := make(chan bool, 1)
	waiting <- true

	closeHandler := bot.Session.AddHandler(
		func(_ *discordgo.Session, msg *discordgo.WebhooksUpdate) {
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
