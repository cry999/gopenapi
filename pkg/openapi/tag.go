package openapi

import "path/filepath"

const (
	fileTags = "tags.yml"
)

// Tag ...
type Tag struct {
	Name         string                 `json:"name,omitempty" yaml:"name,omitempty"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// LoadTags ...
func LoadTags(root string) (tags []*Tag, err error) {
	filename := filepath.Join(root, fileTags)

	if err = loadYAML(filename, &tags); err != nil {
		return
	}
	return
}

// DumpTags ...
func DumpTags(root string, tags []*Tag) (err error) {
	filename := filepath.Join(root, fileTags)

	return dumpYAML(filename, tags)
}
