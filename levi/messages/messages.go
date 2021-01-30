/*
 *
 * Author: Miltiadis Alexis <alexmiltiadis@gmai.com>
 *
 */

package messages

import "github.com/bwmarrin/discordgo"

// guildsCache is being used to keep a map of available guilds
// providing O(1) lookup time.
var guildsCache = make(map[string]bool)

// Init register handlers for the bot
func Init(s *discordgo.Session) {
}
