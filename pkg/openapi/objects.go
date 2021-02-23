package openapi

import "net/url"

// OpenAPI ...
type OpenAPI struct {
	Version      string                   `json:"openapi,omitempty" yaml:"openapi,omitempty"`
	Info         *Info                    `json:"info,omitempty" yaml:"info,omitempty"`
	Servers      []*Server                `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths        Paths                    `json:"paths,omitempty" yaml:"paths,omitempty"`
	Components   *Components              `json:"components,omitempty" yaml:"components,omitempty"`
	Security     []SecurityRequirement    `json:"security,omitempty" yaml:"security,omitempty"`
	Tags         []*Tag                   `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs []*ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// ExternalDocumentation ...
type ExternalDocumentation struct {
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
	URL         *url.URL `json:"url,omitempty" yaml:"url,omitempty"`
}

// Parameter ...
type Parameter struct {
	Name            string `json:"name,omitempty" yaml:"name,omitempty"`
	In              string `json:"in,omitempty" yaml:"in,omitempty"`
	Description     string `json:"description,omitempty" yaml:"description,omitempty"`
	Required        bool   `json:"required,omitempty" yaml:"required,omitempty"`
	Deprecated      bool   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue bool   `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`

	Style         string                   `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       bool                     `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool                     `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Schema        *SchemaOrRef             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example       Any                      `json:"example,omitempty" yaml:"example,omitempty"`
	Examples      map[string]*ExampleOrRef `json:"examples,omitempty" yaml:"examples,omitempty"`

	Content map[string]*MediaType `json:"content,omitempty" yaml:"content,omitempty"`
}

// ParameterOrRef ...
type ParameterOrRef struct {
	Parameter `yaml:",inline"`
	Reference `yaml:",inline"`
}

// IsRef ...
func (por *ParameterOrRef) IsRef() bool { return por.Reference.Ref != "" }

// RequestBody ...
type RequestBody struct {
	Description string                `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]*MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Required    bool                  `json:"required,omitempty" yaml:"required,omitempty"`
}

// RequestBodyOrRef ...
type RequestBodyOrRef struct {
	RequestBody `yaml:",inline"`
	Reference   `yaml:",inline"`
}

// IsRef ...
func (rbor *RequestBodyOrRef) IsRef() bool { return rbor.Reference.Ref != "" }

// MediaType ...
type MediaType struct {
	Schema   *SchemaOrRef             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example  Any                      `json:"example,omitempty" yaml:"example,omitempty"`
	Examples map[string]*ExampleOrRef `json:"examples,omitempty" yaml:"examples,omitempty"`
	Encoding map[string]*Encoding     `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

// Encoding ...
type Encoding struct {
	ContentType   string                  `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers       map[string]*HeaderOrRef `json:"headers,omitempty" yaml:"headers,omitempty"`
	Style         string                  `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       bool                    `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool                    `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
}

// Responses ...
type Responses map[string]*ResponseOrRef

// Response ...
type Response struct {
	Description string                  `json:"description,omitempty" yaml:"description,omitempty"`
	Headers     map[string]*HeaderOrRef `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     map[string]*MediaType   `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]*LinkOrRef   `json:"links,omitempty" yaml:"links,omitempty"`
}

// ResponseOrRef ...
type ResponseOrRef struct {
	Response  `yaml:",inline"`
	Reference `yaml:",inline"`
}

// IsRef ...
func (ror *ResponseOrRef) IsRef() bool { return ror.Reference.Ref != "" }

// Callback ...
type Callback map[string]*PathItem

// CallbackOrRef ...
type CallbackOrRef struct {
	Callback  `yaml:",inline"`
	Reference `yaml:",inline"`
}

// IsRef ...
func (cor *CallbackOrRef) IsRef() bool { return cor.Reference.Ref != "" }

// Example ...
type Example struct {
	Summary       string `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description   string `json:"description,omitempty" yaml:"description,omitempty"`
	Value         Any    `json:"value,omitempty" yaml:"value,omitempty"`
	ExternalValue string `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`
}

// ExampleOrRef ...
type ExampleOrRef struct {
	Example   `yaml:",inline"`
	Reference `yaml:",inline"`
}

// IsRef ...
func (eor *ExampleOrRef) IsRef() bool { return eor.Reference.Ref != "" }

// Link ...
type Link struct {
	OperationRef string                     `json:"operationRef,omitempty" yaml:"operationRef,omitempty"`
	OperationID  string                     `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   map[string]AnyOrExpression `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  AnyOrExpression            `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Description  string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Server       *Server                    `json:"server,omitempty" yaml:"server,omitempty"`
}

// LinkOrRef ...
type LinkOrRef struct {
	Link      `yaml:",inline"`
	Reference `yaml:",inline"`
}

// IsRef ...
func (lor *LinkOrRef) IsRef() bool { return lor.Reference.Ref != "" }

// Header ...
type Header struct {
	Description     string `json:"description,omitempty" yaml:"description,omitempty"`
	Required        bool   `json:"required,omitempty" yaml:"required,omitempty"`
	Deprecated      bool   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue bool   `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`

	Style         string                   `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       bool                     `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool                     `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Schema        *SchemaOrRef             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example       Any                      `json:"example,omitempty" yaml:"example,omitempty"`
	Examples      map[string]*ExampleOrRef `json:"examples,omitempty" yaml:"examples,omitempty"`

	Content map[string]*MediaType `json:"content,omitempty" yaml:"content,omitempty"`
}

// HeaderOrRef ...
type HeaderOrRef struct {
	Header    `yaml:",inline"`
	Reference `yaml:",inline"`
}

// IsRef ...
func (hor *HeaderOrRef) IsRef() bool { return hor.Reference.Ref != "" }

// Reference ...
type Reference struct {
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

// Schema ...
type Schema struct {
	Type       string                  `json:"type,omitempty" yaml:"type,omitempty"`
	Format     string                  `json:"format,omitempty" yaml:"format,omitempty"`
	Properties map[string]*SchemaOrRef `json:"properties,omitempty" yaml:"properties,omitempty"`
	Items      *SchemaOrRef            `json:"items,omitempty" yaml:"items,omitempty"`
}

// SchemaOrRef ...
type SchemaOrRef struct {
	Schema    `yaml:",inline"`
	Reference `yaml:",inline"`
}

// IsRef ...
func (sor *SchemaOrRef) IsRef() bool { return sor.Reference.Ref != "" }

// TODO extensions

// Any ...
type Any interface{}

// AnyOrExpression ...
type AnyOrExpression interface{}
