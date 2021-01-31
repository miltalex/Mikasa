/*
 *
 * Author: Miltiadis Alexis <alexmiltiadis@gmai.com>
 *
 */

package messages

import (
	"github.com/bwmarrin/discordgo"
	"github.com/miltalex/Mikasa/common/logger"
	"github.com/sirupsen/logrus"
)

var log = logger.New(logrus.StandardLogger(), "messages")

// guildsCache is being used to keep a map of
// unavailable guilds providing O(1) lookup time.
var unavailableGuilds = make(map[string]bool)

// Init register handlers for the bot
func Init(s *discordgo.Session) {
	s.AddHandler(ready)
	s.AddHandler(guildCreate)
	s.AddHandler(guildDelete)
}
