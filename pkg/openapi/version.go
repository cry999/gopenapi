package openapi

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	fileOpenAPIVersion = "openapi_version"
)

// LoadOpenAPIVersion ...
func LoadOpenAPIVersion(root string) (ver string, err error) {
	path := filepath.Join(root, fileOpenAPIVersion)

	bver, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	ver = string(bver)
	ver = strings.Trim(ver, "\n")
	return
}

// DumpOpenAPIVersion ...
func DumpOpenAPIVersion(root string, ver string) (err error) {
	path := filepath.Join(root, fileOpenAPIVersion)

	return ioutil.WriteFile(path, []byte(ver), 0o644)
}
