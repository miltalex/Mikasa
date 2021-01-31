/*
 *
 * Author: Miltiadis Alexis <alexmiltiadis@gmai.com>
 *
 */

package router

// Router is  struct that will retrieve the correct endpoint
// for the given command
type Router struct {
	Prefixes   []string
	IgnoreCase bool
	Commands   []*Command
}
