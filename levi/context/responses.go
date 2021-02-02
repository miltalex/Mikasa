/*
 *
 * Author: Miltiadis Alexis <alexmiltiadis@gmai.com>
 *
 */

package context

import (
	"github.com/bwmarrin/discordgo"
	"github.com/miltalex/Mikasa/common/logger"
	"github.com/sirupsen/logrus"
)

var log = logger.New(logrus.StandardLogger(), "responses")

// Ctx represents the context for a command event
type Ctx struct {
	Session *discordgo.Session
	Event   *discordgo.MessageCreate
}

// ExecutionHandler represents a handler for a context execution
type ExecutionHandler func(*Ctx)

// RespondText responds with the given text message
func (ctx *Ctx) RespondText(text string) error {
	_, err := ctx.Session.ChannelMessageSend(ctx.Event.ChannelID, text)
	return err
}

// RespondEmbed responds with the given embed message
func (ctx *Ctx) RespondEmbed(embed *discordgo.MessageEmbed) error {
	_, err := ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embed)
	return err
}

// RespondTextEmbed responds with the given text and embed message
func (ctx *Ctx) RespondTextEmbed(text string, embed *discordgo.MessageEmbed) error {
	_, err := ctx.Session.ChannelMessageSendComplex(ctx.Event.ChannelID, &discordgo.MessageSend{
		Content: text,
		Embed:   embed,
	})
	return err
}

// RespondPlayMusic responds with the given text and embed message
func (ctx *Ctx) RespondPlayMusic() error {

	c, err := ctx.Session.State.Channel(ctx.Event.ChannelID)
	if err != nil {
		log.Println("fail find channel")
		return err
	}

	g, err := ctx.Session.State.Guild(c.GuildID)
	if err != nil {
		log.Println("fail find guild")
		return err
	}

	for _, vs := range g.VoiceStates {
		if vs.UserID == ctx.Event.Author.ID {
			log.Println("trying to connect to channel")
			err = ctx.Session.ChannelVoiceJoinManual(c.GuildID, vs.ChannelID, false, false)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("channel voice join succeeded")
			}
		}
	}

	// query := ctx.Event.Content[8:]
	// node, err := lavalink.BestNode()
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	// tracks, err := node.LoadTracks(query)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	// if tracks.Type != gavalink.TrackLoaded {
	// 	log.Println("weird tracks type", tracks.Type)
	// }
	// track := tracks.Tracks[0].Data
	// err = player.Play(track)
	// if err != nil {
	// 	log.Println(err)
	// }

	// _, err = ctx.Session.ChannelMessageSendComplex(ctx.Event.ChannelID, &discordgo.MessageSend{
	// 	Content: text,
	// 	Embed:   embed,
	// })
	return err
}
