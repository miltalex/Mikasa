/*
 *
 * Author: Miltiadis Alexis <alexmiltiadis@gmai.com>
 *
 */

package command

// Command represent any command that the user can request
type Command struct {
	Name        string
	Aliases     []string
	Description string
	Usage       string
	Example     string
	IgnoreCase  bool
	Subcommands []*Command
}
