/*
 *
 * Author: Miltiadis Alexis <alexmiltiadis@gmai.com>
 *
 */

package common

import (
	"strings"

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

// StringArrayContains checks whether or not the given string array contains the given string
func StringArrayContains(array []string, str string, ignoreCase bool) bool {
	if ignoreCase {
		str = strings.ToLower(str)
	}
	for _, value := range array {
		if ignoreCase {
			value = strings.ToLower(value)
		}
		if value == str {
			return true
		}
	}
	return false
}

// StringHasPrefix checks whether or not the string contains one of the given prefixes and returns the string without the prefix
func StringHasPrefix(str string, prefixes []string, ignoreCase bool) (bool, string) {
	for _, prefix := range prefixes {
		stringToCheck := str
		if ignoreCase {
			stringToCheck = strings.ToLower(stringToCheck)
			prefix = strings.ToLower(prefix)
		}
		if strings.HasPrefix(stringToCheck, prefix) {
			return true, string(str[len(prefix):])
		}
	}
	return false, str
}

// StringTrimPreSuffix returns the string without the defined pre- and suffix
func StringTrimPreSuffix(str string, preSuffix string) string {
	if !(strings.HasPrefix(str, preSuffix) && strings.HasSuffix(str, preSuffix)) {
		return str
	}
	return strings.TrimPrefix(strings.TrimSuffix(str, preSuffix), preSuffix)
}

// Equals provides a simple method to check whether or not 2 strings are equal
func Equals(str1, str2 string, ignoreCase bool) bool {
	if !ignoreCase {
		return str1 == str2
	}
	return strings.ToLower(str1) == strings.ToLower(str2)
}
