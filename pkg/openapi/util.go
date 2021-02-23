package openapi

import (
	"net/url"
	"os"

	"gopkg.in/yaml.v2"
)

func loadYAML(filename string, v interface{}) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	return yaml.NewDecoder(f).Decode(v)
}

func dumpYAML(filename string, v interface{}) (err error) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return
	}
	defer f.Close()

	return yaml.NewEncoder(f).Encode(v)
}

// URL ...
type URL struct {
	*url.URL
}

// MarshalYAML ...
func (u *URL) MarshalYAML() (interface{}, error) {
	return u.String(), nil
}

// UnmarshalYAML ...
func (u *URL) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var s string
	if err = unmarshal(&s); err != nil {
		return
	}
	u.URL, err = url.Parse(s)
	return
}

// MustParseURL ...
func MustParseURL(s string) *URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return &URL{u}
}
