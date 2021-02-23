package openapi

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	dirComponents     = "components"
	dirSchema         = "schemas"
	dirResponse       = "responses"
	dirParameter      = "parameters"
	dirExample        = "examples"
	dirRequestBody    = "requestBodies"
	dirHeader         = "headers"
	dirSecuritySchema = "securitySchemes"
	dirLink           = "links"
	dirCallback       = "callbacks"
)

// Components ...
type Components struct {
	Schemas         map[string]*SchemaOrRef         `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Responses       map[string]*ResponseOrRef       `json:"responses,omitempty" yaml:"responses,omitempty"`
	Parameters      map[string]*ParameterOrRef      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Examples        map[string]*ExampleOrRef        `json:"examples,omitempty" yaml:"examples,omitempty"`
	RequestBodies   map[string]*RequestBodyOrRef    `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Headers         map[string]*HeaderOrRef         `json:"headers,omitempty" yaml:"headers,omitempty"`
	SecuritySchemes map[string]*SecuritySchemeOrRef `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Links           map[string]*LinkOrRef           `json:"links,omitempty" yaml:"links,omitempty"`
	Callbacks       map[string]*CallbackOrRef       `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
}

// SecurityScheme ...
type SecurityScheme struct {
	Type             string      `json:"type,omitempty" yaml:"type,omitempty"`
	Description      string      `json:"description,omitempty" yaml:"description,omitempty"`
	Name             string      `json:"name,omitempty" yaml:"name,omitempty"`
	In               string      `json:"in,omitempty" yaml:"in,omitempty"`
	Scheme           string      `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	BearerFormat     string      `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
	Flows            *OAuthFlows `json:"flows,omitempty" yaml:"flows,omitempty"`
	OpenIDConnectURL string      `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"`
}

// OAuthFlows ...
type OAuthFlows struct {
	Implicit          *OAuthFlow `json:"implicit,omitempty" yaml:"implicit,omitempty"`
	Password          *OAuthFlow `json:"password,omitempty" yaml:"password,omitempty"`
	ClientCredentials *OAuthFlow `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
}

// OAuthFlow ...
type OAuthFlow struct {
	AuthorizationURL *URL              `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`
	TokenURL         *URL              `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
	RefreshURL       *URL              `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes,omitempty" yaml:"scopes,omitempty"`
}

// SecuritySchemeOrRef ...
type SecuritySchemeOrRef struct {
	SecurityScheme `yaml:",inline"`
	Reference      `yaml:",inline"`
}

// IsRef ...
func (sor *SecuritySchemeOrRef) IsRef() bool { return sor.Reference.Ref != "" }

// LoadComponents ...
func LoadComponents(root string) (_ *Components, err error) {
	newRoot := filepath.Join(root, dirComponents)

	schemas, err := loadSchemas(newRoot)
	if err != nil {
		return
	}

	responses, err := loadResponses(newRoot)
	if err != nil {
		return
	}

	requestBodies, err := loadRequestBodies(newRoot)
	if err != nil {
		return
	}

	securitySchemas, err := loadSecuritySchemas(newRoot)
	if err != nil {
		return
	}

	return &Components{
		Schemas:         schemas,
		Responses:       responses,
		RequestBodies:   requestBodies,
		SecuritySchemes: securitySchemas,
	}, nil
}

func filenameWithoutExt(filename string) string {
	ext := filepath.Ext(filename)
	base := filepath.Base(filename)
	return strings.TrimRight(base, ext)
}

func loadSchemas(root string) (_ map[string]*SchemaOrRef, err error) {
	schemas := map[string]*SchemaOrRef{}

	dirname := filepath.Join(root, dirSchema)
	if err := walk(dirname, func(f *os.File) (err error) {
		schema, err := loadSchema(f)
		if err != nil {
			return fmt.Errorf("%s: %v", f.Name(), err)
		}

		schemaName := filenameWithoutExt(f.Name())
		schemas[schemaName] = schema

		return
	}); err != nil {
		return nil, err
	}
	return schemas, nil
}

func loadSchema(r io.Reader) (_ *SchemaOrRef, err error) {
	var sor SchemaOrRef
	if err = yaml.NewDecoder(r).Decode(&sor); err != nil {
		return
	}
	return &sor, nil
}

func loadResponses(root string) (_ map[string]*ResponseOrRef, err error) {
	responses := map[string]*ResponseOrRef{}

	dirname := filepath.Join(root, dirResponse)
	if err := walk(dirname, func(f *os.File) (err error) {
		res, err := loadResponse(f)
		if err != nil {
			return fmt.Errorf("%s: %v", f.Name(), err)
		}

		responseName := filenameWithoutExt(f.Name())
		responses[responseName] = res

		return
	}); err != nil {
		return nil, err
	}
	return responses, nil
}

func loadResponse(r io.Reader) (_ *ResponseOrRef, err error) {
	var ror ResponseOrRef
	if err = yaml.NewDecoder(r).Decode(&ror); err != nil {
		return
	}
	return &ror, nil
}

func loadRequestBodies(root string) (_ map[string]*RequestBodyOrRef, err error) {
	bodies := map[string]*RequestBodyOrRef{}

	dirname := filepath.Join(root, dirRequestBody)
	if err := walk(dirname, func(f *os.File) error {
		body, err := loadRequestBody(f)
		if err != nil {
			return fmt.Errorf("%s: %v", f.Name(), err)
		}

		bodyName := filenameWithoutExt(f.Name())
		bodies[bodyName] = body

		return nil
	}); err != nil {
		return nil, err
	}
	return bodies, nil
}

func loadRequestBody(r io.Reader) (*RequestBodyOrRef, error) {
	var rbor RequestBodyOrRef
	if err := yaml.NewDecoder(r).Decode(&rbor); err != nil {
		return nil, err
	}
	return &rbor, nil
}

func loadSecuritySchemas(root string) (_ map[string]*SecuritySchemeOrRef, err error) {
	ss := map[string]*SecuritySchemeOrRef{}

	dirname := filepath.Join(root, dirSecuritySchema)
	if err := walk(dirname, func(f *os.File) error {
		ssor, err := loadSecuritySchema(f)
		if err != nil {
			return err
		}

		name := filenameWithoutExt(f.Name())
		ss[name] = ssor

		return nil
	}); err != nil {
		return nil, err
	}

	return ss, nil
}

func loadSecuritySchema(r io.Reader) (*SecuritySchemeOrRef, error) {
	var ssor SecuritySchemeOrRef
	if err := yaml.NewDecoder(r).Decode(&ssor); err != nil {
		return nil, err
	}
	return &ssor, nil
}

func walk(dirname string, callback func(f *os.File) error) (err error) {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		return
	}

	for _, fileinfo := range fileInfos {
		if fileinfo.IsDir() {
			// TODO recursive directory is not supported
			continue
		}

		filename := filepath.Join(dirname, fileinfo.Name())
		ext := filepath.Ext(filename)
		if ext != ".yml" && ext != ".yaml" && ext != ".json" {
			continue
		}

		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		if err := callback(f); err != nil {
			return err
		}
	}
	return
}
