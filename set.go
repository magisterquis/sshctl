package sshctl

import (
	"regexp"
)

// Set represents all the hosts read from a particular file.
type Set struct {
	groups map[string]*Group
	file   string // File from which set is read
}

// New reurns a new Group as read from the specified file.  If file is "",
// an empty group will be returned.  If file has no hosts, an empty set is
// returned.  Any changes to the set will be written to the file.
func New(file string) Set {
}

// Group returns a pre-defined group of hosts.
func (s Set) Group(name string) *Group {
}

// AddGroup makes a new group of hosts based on a regular expression and
// returns a pointer to that group.  The group will be saved for later
// retrieval with Group()
func (s *Set) NewGroup(re regexp.Regexp) *Group {
}

// 
