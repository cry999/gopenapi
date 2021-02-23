package openapi

import (
	"path/filepath"
)

const (
	fileInfo = "info.yml"
)

// Info ...
type Info struct {
	Title          string   `json:"title,omitempty" yaml:"title,omitempty"`
	Description    string   `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string   `json:"version,omitempty" yaml:"version,omitempty"`
}

// Contact ...
type Contact struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   *URL   `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

// License ...
type License struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	URL  *URL   `json:"url,omitempty" yaml:"url,omitempty"`
}

// LoadInfo ...
func LoadInfo(root string) (_ *Info, err error) {
	filename := filepath.Join(root, fileInfo)

	var info Info
	if err = loadYAML(filename, &info); err != nil {
		return
	}
	return &info, nil
}

// DumpInfo ...
func DumpInfo(root string, info *Info) (err error) {
	filename := filepath.Join(root, fileInfo)

	return dumpYAML(filename, info)
}
