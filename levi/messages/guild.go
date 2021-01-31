/*
 *
 * Author: Miltiadis Alexis <alexmiltiadis@gmai.com>
 *
 */

package messages

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"github.com/miltalex/Mikasa/common"
	"github.com/miltalex/Mikasa/constants"
)

func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	log.WithField("guildID", event.ID).Info("guildCreated Event received.")
	if _, ok := unavailableGuilds[event.ID]; ok {
		// An unavailable guild became available, delete it from our cache and continue.
		delete(unavailableGuilds, event.ID)
		log.WithField("guildID", event.ID).Info("Remove guild from unavailable cache.")

		// Let us know when all guilds are loaded.
		if len(unavailableGuilds) == 0 {
			log.WithField("guilds", len(s.State.Guilds)).Info("Finish loading guilds")
		}

		return
	}
	// A new guild joined.

	owner, err := s.User(event.OwnerID)
	if err != nil {
		return
	}

	// So best we can show is the ID and the created time.
	created, _ := discordgo.SnowflakeTimestamp(event.ID)

	fields := make([]*discordgo.MessageEmbedField, 0)

	fields = append(fields, common.AddField("Owner", owner.String()))
	fields = append(fields, common.AddField("Member Count", strconv.Itoa(event.MemberCount)))
	fields = append(fields, common.AddField("Created At", fmt.Sprintf("%s (%s)", created.Format("2 January 2006"), humanize.Time(created))))

	// Set the new presence.
	s.UpdateGameStatus(0, fmt.Sprintf("!help | %d Servers!", len(s.State.Guilds)))

	var channelID string

	for _, channel := range event.Guild.Channels {
		if channel.Type == discordgo.ChannelTypeGuildText {
			channelID = channel.ID
			break
		}
	}
	_, err = s.ChannelMessageSendEmbed(channelID,
		&discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s has joined a new server!", constants.BotName),
			Thumbnail:   common.SetThumbnail(event.IconURL()),
			Description: event.ID,
			Color:       0xDFAC7C,
			Fields:      fields,
			Footer:      &discordgo.MessageEmbedFooter{Text: event.ID},
		})
	if err != nil {
		log.Error(err)
	}
}

func guildDelete(s *discordgo.Session, event *discordgo.GuildDelete) {
	// If this event got triggered because the guild became unavailable
	// Add it to our cache so guildCreate can be aware when it becomes
	// available and not treat it as a new guild.
	log.WithField("guildID", event.ID).Info("guildDelete Event received.")
	if event.Unavailable {
		unavailableGuilds[event.ID] = true
		log.WithField("guildID", event.ID).Info("Add guild to unavailable cache.")
		return
	}

	// Set the new presence.
	s.UpdateGameStatus(0, fmt.Sprintf("!help | %d Servers!", len(s.State.Guilds)))

	// And discordgo already cleared it from the cache by this point.
	// So best we can show is the ID and the created time.
	created, _ := discordgo.SnowflakeTimestamp(event.ID)
	fields := make([]*discordgo.MessageEmbedField, 0)

	fields = append(fields, common.AddField("Created At", fmt.Sprintf("%s (%s)", created.Format("2 January 2006"), humanize.Time(created))))

	var channelID string

	for _, channel := range event.Guild.Channels {
		if channel.Type == discordgo.ChannelTypeGuildText {
			channelID = channel.ID
			break
		}
	}
	_, err := s.ChannelMessageSendEmbed(channelID,
		&discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s has left a server.", constants.BotName),
			Description: event.ID,
			Color:       0xFF0000,
			Fields:      fields,
		})
	if err != nil {
		log.Error(err)
	}
}
