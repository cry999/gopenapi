package openapi

import (
	"path/filepath"
)

const (
	fileSecurity = "security.yml"
)

// SecurityRequirement ...
type SecurityRequirement map[string][]string

// LoadSecurity ...
func LoadSecurity(root string) (security []SecurityRequirement, err error) {
	filename := filepath.Join(root, fileSecurity)

	if err = loadYAML(filename, &security); err != nil {
		return
	}
	return
}

// DumpSecurity ...
func DumpSecurity(root string, security []SecurityRequirement) (err error) {
	filename := filepath.Join(root, fileSecurity)

	return dumpYAML(filename, security)
}
