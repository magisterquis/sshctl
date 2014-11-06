package sshctl

import (
	"regexp"
)

// Group represents a group of hosts, defined in the file or with
// Set.NewGroup().
type Group struct {
	re   *regexp.Regexp // The regular expression that defines the group
	name string         // The name of the group

}

// Add a host to the group and returns a pointer to the (changed) group
func (g *Group) add(host Host) *Group {
}
