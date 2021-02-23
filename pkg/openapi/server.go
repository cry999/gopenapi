package openapi

import (
	"path/filepath"
)

const (
	fileServers = "servers.yml"
)

// Server ...
type Server struct {
	URL         *URL                       `json:"url,omitempty" yaml:"url,omitempty"`
	Description string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
}

// ServerVariable ...
type ServerVariable struct {
	Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default     string   `json:"default,omitempty" yaml:"default,omitempty"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
}

// LoadServers ...
func LoadServers(root string) (servers []*Server, err error) {
	filename := filepath.Join(root, fileServers)

	if err = loadYAML(filename, &servers); err != nil {
		return
	}
	return
}

// DumpServers ...
func DumpServers(root string, servers []*Server) (err error) {
	filename := filepath.Join(root, fileServers)

	return dumpYAML(filename, servers)
}
