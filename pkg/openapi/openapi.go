package openapi

// LoadProject ...
func LoadProject(projectDir string) (*OpenAPI, error) {
	version, err := LoadOpenAPIVersion(projectDir)
	if err != nil {
		return nil, err
	}
	info, err := LoadInfo(projectDir)
	if err != nil {
		return nil, err
	}
	servers, err := LoadServers(projectDir)
	if err != nil {
		return nil, err
	}
	paths, err := LoadPaths(projectDir)
	if err != nil {
		return nil, err
	}
	components, err := LoadComponents(projectDir)
	if err != nil {
		return nil, err
	}
	security, err := LoadSecurity(projectDir)
	if err != nil {
		return nil, err
	}
	tags, err := LoadTags(projectDir)
	if err != nil {
		return nil, err
	}

	openapi := &OpenAPI{
		Version:    version,
		Info:       info,
		Servers:    servers,
		Paths:      paths,
		Components: components,
		Security:   security,
		Tags:       tags,
		// ExternalDocs: []*ExternalDocumentation{},
	}
	return openapi, nil
}

// DumpInOneFile ...
func DumpInOneFile(output string, openapi *OpenAPI) error {
	return dumpYAML(output, openapi)
}
