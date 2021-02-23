package openapi

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	dirPaths = "paths"

	filePathIndex = "index.yml"
)

// Paths ...
type Paths map[string]*PathItem

// PathItem ...
type PathItem struct {
	Ref         string            `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string            `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string            `json:"description,omitempty" yaml:"description,omitempty"`
	Get         *Operation        `json:"get,omitempty" yaml:"get,omitempty"`
	Put         *Operation        `json:"put,omitempty" yaml:"put,omitempty"`
	Post        *Operation        `json:"post,omitempty" yaml:"post,omitempty"`
	Delete      *Operation        `json:"delete,omitempty" yaml:"delete,omitempty"`
	Options     *Operation        `json:"options,omitempty" yaml:"options,omitempty"`
	Head        *Operation        `json:"head,omitempty" yaml:"head,omitempty"`
	Patch       *Operation        `json:"patch,omitempty" yaml:"patch,omitempty"`
	Trace       *Operation        `json:"trace,omitempty" yaml:"trace,omitempty"`
	Servers     []*Server         `json:"servers,omitempty" yaml:"servers,omitempty"`
	Parameters  []*ParameterOrRef `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation ...
type Operation struct {
	Tags         []string                  `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary      string                    `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string                    `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocumentation    `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	OperationID  string                    `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   []*ParameterOrRef         `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  *RequestBodyOrRef         `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses    *Responses                `json:"responses,omitempty" yaml:"responses,omitempty"`
	Callbacks    map[string]*CallbackOrRef `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	Deprecated   bool                      `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security     []*SecurityRequirement    `json:"security,omitempty" yaml:"security,omitempty"`
	Servers      []*Server                 `json:"servers,omitempty" yaml:"servers,omitempty"`
}

// LoadPaths ...
func LoadPaths(root string) (Paths, error) {
	paths := Paths{}
	pathsRoot := filepath.Join(root, dirPaths)

	if err := loadPathItem(pathsRoot, pathsRoot, paths); err != nil {
		return nil, err
	}

	return paths, nil
}

func loadPathItem(cwd, root string, paths Paths) error {
	index := filepath.Join(cwd, filePathIndex)
	if _, err := os.Stat(index); err == nil || os.IsExist(err) {
		var pathitem PathItem
		if err := loadYAML(index, &pathitem); err != nil {
			return err
		}

		if err := loadPathItemOperations(cwd, &pathitem); err != nil {
			return err
		}

		servers, err := LoadServers(cwd)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		} else {
			pathitem.Servers = servers
		}

		path := strings.TrimPrefix(cwd, root)
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}

		paths[path] = &pathitem
	}

	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		return err
	}

	for _, fileinfo := range files {
		if !fileinfo.IsDir() {
			continue
		}

		nwd := filepath.Join(cwd, fileinfo.Name())
		if err := loadPathItem(nwd, root, paths); err != nil {
			return err
		}
	}

	return nil
}

func loadPathItemOperations(cwd string, item *PathItem) error {
	for filename, op := range map[string]*Operation{
		"get.yml":     item.Get,
		"put.yml":     item.Put,
		"post.yml":    item.Post,
		"delete.yml":  item.Delete,
		"options.yml": item.Options,
		"head.yml":    item.Head,
		"patch.yml":   item.Patch,
		"trace.yml":   item.Trace,
	} {
		filename = filepath.Join(cwd, filename)
		if _, err := os.Stat(filename); err != nil && !os.IsExist(err) {
			continue
		}

		if err := loadYAML(filename, op); err != nil {
			return err
		}
	}
	return nil
}
