package sshctl

// Host represents a host to be controlled
type Host struct {
	user    string // Username
	address string // Address of SSH server
	port    uint16 // Port on SSH server
}

// User returns the Host's username.
func (h *Host) User() string {
}

// SetUser sets the Host's username.
func (h *Host) SetUser(u string) {
}

// Adress returns the Host's address.
func (h *Host) Adress() string {
}

// SetAdress sets the Host's address.
func (h *Host) SetAdress(a string) {
}

// Port returns the Host's port.
func (h *Host) Port() uint16 {
}

// SetPort sets the Host's port.
func (h *Host) SetPort(p uint16) {
}
