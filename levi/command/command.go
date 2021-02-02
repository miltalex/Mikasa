/*
 *
 * Author: Miltiadis Alexis <alexmiltiadis@gmai.com>
 *
 */

package command

import (
	"sort"

	"github.com/miltalex/Mikasa/common"
	"github.com/miltalex/Mikasa/levi/context"
)

// Command represent any command that the user can request
type Command struct {
	Name        string
	Aliases     []string
	Description string
	Usage       string
	Example     string
	IgnoreCase  bool
	Subcommands []*Command
	// RateLimiter RateLimiter
	Handler context.ExecutionHandler
}

// GetSubCmd returns the sub command with the given name if it exists
func (command *Command) GetSubCmd(name string) *Command {

	// Loop through all commands to find the correct one
	for _, subCommand := range command.Subcommands {
		// Define the slice to check
		toCheck := make([]string, len(subCommand.Aliases)+1)
		toCheck = append(toCheck, subCommand.Name)
		toCheck = append(toCheck, subCommand.Aliases...)
		sort.Slice(toCheck, func(i, j int) bool {
			return len(toCheck[i]) > len(toCheck[j])
		})

		// Check the prefix of the string
		if common.StringArrayContains(toCheck, name, subCommand.IgnoreCase) {
			return subCommand
		}
	}
	return nil
}

// NotifyRateLimiter notifies the rate limiter about a new execution and returns false if the user is being rate limited
// func (command *Command) NotifyRateLimiter(ctx *Ctx) bool {
// 	if command.RateLimiter == nil {
// 		return true
// 	}
// 	return command.RateLimiter.NotifyExecution(ctx)
// }

// Trigger triggers the given command
func (command *Command) Trigger(ctx *context.Ctx) {
	command.Handler(ctx)
}
