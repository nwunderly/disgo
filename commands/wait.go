package commands

import "github.com/bwmarrin/discordgo"

//func handlerForInterface(handler interface{}) EventHandler {
//	switch v := handler.(type) {
//	case func(*Session, interface{}):
//		return interfaceEventHandler(v)
//	case func(*Session, *ChannelCreate):
//		return channelCreateEventHandler(v)
//	case func(*Session, *ChannelDelete):
//		return channelDeleteEventHandler(v)
//	case func(*Session, *ChannelPinsUpdate):
//		return channelPinsUpdateEventHandler(v)
//	case func(*Session, *ChannelUpdate):
//		return channelUpdateEventHandler(v)
//	case func(*Session, *Connect):
//		return connectEventHandler(v)
//	case func(*Session, *Disconnect):
//		return disconnectEventHandler(v)
//	case func(*Session, *Event):
//		return eventEventHandler(v)
//	case func(*Session, *GuildBanAdd):
//		return guildBanAddEventHandler(v)
//	case func(*Session, *GuildBanRemove):
//		return guildBanRemoveEventHandler(v)
//	case func(*Session, *GuildCreate):
//		return guildCreateEventHandler(v)
//	case func(*Session, *GuildDelete):
//		return guildDeleteEventHandler(v)
//	case func(*Session, *GuildEmojisUpdate):
//		return guildEmojisUpdateEventHandler(v)
//	case func(*Session, *GuildIntegrationsUpdate):
//		return guildIntegrationsUpdateEventHandler(v)
//	case func(*Session, *GuildMemberAdd):
//		return guildMemberAddEventHandler(v)
//	case func(*Session, *GuildMemberRemove):
//		return guildMemberRemoveEventHandler(v)
//	case func(*Session, *GuildMemberUpdate):
//		return guildMemberUpdateEventHandler(v)
//	case func(*Session, *GuildMembersChunk):
//		return guildMembersChunkEventHandler(v)
//	case func(*Session, *GuildRoleCreate):
//		return guildRoleCreateEventHandler(v)
//	case func(*Session, *GuildRoleDelete):
//		return guildRoleDeleteEventHandler(v)
//	case func(*Session, *GuildRoleUpdate):
//		return guildRoleUpdateEventHandler(v)
//	case func(*Session, *GuildUpdate):
//		return guildUpdateEventHandler(v)
//	case func(*Session, *MessageAck):
//		return messageAckEventHandler(v)
//	case func(*Session, *MessageCreate):
//		return messageCreateEventHandler(v)
//	case func(*Session, *MessageDelete):
//		return messageDeleteEventHandler(v)
//	case func(*Session, *MessageDeleteBulk):
//		return messageDeleteBulkEventHandler(v)
//	case func(*Session, *MessageReactionAdd):
//		return messageReactionAddEventHandler(v)
//	case func(*Session, *MessageReactionRemove):
//		return messageReactionRemoveEventHandler(v)
//	case func(*Session, *MessageReactionRemoveAll):
//		return messageReactionRemoveAllEventHandler(v)
//	case func(*Session, *MessageUpdate):
//		return messageUpdateEventHandler(v)
//	case func(*Session, *PresenceUpdate):
//		return presenceUpdateEventHandler(v)
//	case func(*Session, *PresencesReplace):
//		return presencesReplaceEventHandler(v)
//	case func(*Session, *RateLimit):
//		return rateLimitEventHandler(v)
//	case func(*Session, *Ready):
//		return readyEventHandler(v)
//	case func(*Session, *RelationshipAdd):
//		return relationshipAddEventHandler(v)
//	case func(*Session, *RelationshipRemove):
//		return relationshipRemoveEventHandler(v)
//	case func(*Session, *Resumed):
//		return resumedEventHandler(v)
//	case func(*Session, *TypingStart):
//		return typingStartEventHandler(v)
//	case func(*Session, *UserGuildSettingsUpdate):
//		return userGuildSettingsUpdateEventHandler(v)
//	case func(*Session, *UserNoteUpdate):
//		return userNoteUpdateEventHandler(v)
//	case func(*Session, *UserSettingsUpdate):
//		return userSettingsUpdateEventHandler(v)
//	case func(*Session, *UserUpdate):
//		return userUpdateEventHandler(v)
//	case func(*Session, *VoiceServerUpdate):
//		return voiceServerUpdateEventHandler(v)
//	case func(*Session, *VoiceStateUpdate):
//		return voiceStateUpdateEventHandler(v)
//	case func(*Session, *WebhooksUpdate):
//		return webhooksUpdateEventHandler(v)
//	}
//
//	return nil
//}

type WaitForHandler func(interface{} bool)

func (bot *Bot) waitForHandler() {}

func (bot *Bot) WaitFor(check func(interface{}) bool) interface{} {
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