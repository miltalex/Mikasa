/*
 *
 * Author: Miltiadis Alexis <alexmiltiadis@gmai.com>
 *
 */

package common

import (
	"github.com/bwmarrin/discordgo"
)

// SetThumbnail ...
func SetThumbnail(args ...string) *discordgo.MessageEmbedThumbnail {
	var URL string
	var proxyURL string

	if len(args) == 0 {
		return nil
	}

	if len(args) > 0 {
		URL = args[0]
	}

	if len(args) > 1 {
		proxyURL = args[1]
	}

	thumbnail := &discordgo.MessageEmbedThumbnail{
		URL:      URL,
		ProxyURL: proxyURL,
	}
	return thumbnail
}

// AddField [name] [value]
func AddField(name, value string) *discordgo.MessageEmbedField {
	if len(value) > 1024 {
		value = value[:1024]
	}

	if len(name) > 1024 {
		name = name[:1024]
	}

	return &discordgo.MessageEmbedField{
		Name:  name,
		Value: value,
	}

}
