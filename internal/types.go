package internal

// Data of one project
type NFACUConfigObject struct {
	Path     string            `json:"path"`
	Settings map[string]string `json:"settings"`
}

// NFACU config scheme
type NFACUConfig []NFACUConfigObject
