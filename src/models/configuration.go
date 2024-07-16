package models

import "encoding/xml"

// Configuration is a configuration.
type Configuration struct {
	XMLName xml.Name `json:"-" yaml:"-" toml:"-" xml:"configuration"`

	// Name is the name of the configuration. (namespace)
	Name string `json:"name" yaml:"name" toml:"name" xml:"name,attr"`

	// Resources is a key-value map of the strings.
	Resources TranslationResourcesMap `json:"resources" yaml:"resources" toml:"resources" xml:"resources"`
}
