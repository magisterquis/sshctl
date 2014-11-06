package sshctl

import (
	"regexp"
)

// Set represents all the hosts read from a particular file.
type Set struct {
}

// Group returns a pre-defined group of hosts.
func (s Set) Group(name string) *Group {
}

// AddGroup makes a new group of hosts based on a regular expression and
// returns a pointer to that group.  The group will be saved for later
// retrieval with Group()
func (s *Set) NewGroup(re regexp.Regexp) *Group {
}
