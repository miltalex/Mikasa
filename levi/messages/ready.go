/*
 *
 * Author: Miltiadis Alexis <alexmiltiadis@gmai.com>
 *
 */

package messages

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ready(s *discordgo.Session, r *discordgo.Ready) {
	log.WithField("User", r.User.String()).Info("logged in.")
	// Set the playing status.
	s.UpdateGameStatus(0, fmt.Sprintf("!help | %d Servers!", len(s.State.Guilds)))

	// Cache the unavailable guilds.
	for _, guild := range r.Guilds {
		if guild.Unavailable {
			unavailableGuilds[guild.ID] = true
		}
	}

	// TODO miltalex: here we should add the lavalink node
}
