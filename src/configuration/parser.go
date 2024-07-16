package configuration

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/nikita-petko/translation-resource-generator/flags"
	"github.com/nikita-petko/translation-resource-generator/models"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

// This model is just a key-value pair for the fully qualified name of the template and the template itself.
// This is used to parse the template files.
type ConfigurationPair struct {
	// Fully qualified name of the template.
	TemplateFullyQualifiedPath string

	// The model of the template.
	Configuration *models.Configuration
}

func parseJSONFile(fileName string) (*models.Configuration, error) {
	var entity models.Configuration

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	jsonParser := json.NewDecoder(jsonFile)
	if err = jsonParser.Decode(&entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func parseYAMLFile(fileName string) (*models.Configuration, error) {
	var entity models.Configuration

	yamlFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer yamlFile.Close()

	yamlParser := yaml.NewDecoder(yamlFile)
	if err = yamlParser.Decode(&entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func parseTOMLFile(fileName string) (*models.Configuration, error) {
	var entity models.Configuration

	tomlFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer tomlFile.Close()

	tomlParser := toml.NewDecoder(tomlFile)
	if err = tomlParser.Decode(&entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func parseXMLFile(fileName string) (*models.Configuration, error) {
	var entity models.Configuration

	xmlFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer xmlFile.Close()

	xmlParser := xml.NewDecoder(xmlFile)
	if err = xmlParser.Decode(&entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func parseFileDependingOnExtension(fileName string) (*models.Configuration, error) {
	if fileName == "" {
		return nil, nil
	}

	fileExtension := path.Ext(fileName)

	switch fileExtension {
	case ".json":
		return parseJSONFile(fileName)
	case ".yaml":
		return parseYAMLFile(fileName)
	case ".yml":
		return parseYAMLFile(fileName)
	case ".toml":
		return parseTOMLFile(fileName)
	case ".xml":
		return parseXMLFile(fileName)
	default:
		return nil, nil
	}
}

var (
	ErrConfigurationDirectoryNotSpecified = errors.New("configuration directory not specified")
	ErrNilConfiguration                   = errors.New("the configuration is nil")
	ErrConfigurationNameNotSpecified      = errors.New("the translation namespace name is not specified")
	ErrResourcesEmpty                     = errors.New("at least one resource must be specified")
	ErrTranslationResourcesEmpty          = errors.New("at least one translation resource must be specified")
	ErrEnglishStringNotSpecified          = errors.New("the english string is not specified")
)

func validateConfiguration(configuration *models.Configuration) error {
	if configuration == nil {
		return ErrNilConfiguration
	}

	if configuration.Name == "" {
		return ErrConfigurationNameNotSpecified
	}

	if len(configuration.Resources) == 0 {
		return ErrResourcesEmpty
	}

	for _, resource := range configuration.Resources {
		if len(resource) == 0 {
			return ErrTranslationResourcesEmpty
		}

		for _, translationResource := range resource {
			if translationResource.EnglishString == "" {
				return fmt.Errorf("%w: %s", ErrEnglishStringNotSpecified, configuration.Name)
			}

			// Make sure all the keys in the translations are lowercase.
			for key := range translationResource.Translations {
				value := translationResource.Translations[key]
				delete(translationResource.Translations, key)
				translationResource.Translations[strings.ToLower(key)] = value
			}
		}

	}

	return nil
}

// Parse parses the configuration files.
// And produces the full configuration model.
func Parse() ([]*ConfigurationPair, error) {
	if *flags.ConfigurationDirectoryFlag == "" {
		return nil, ErrConfigurationDirectoryNotSpecified
	}

	var configurations []*ConfigurationPair

	err := filepath.Walk(*flags.ConfigurationDirectoryFlag, func(fileName string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fileInfo.IsDir() && !*flags.RecurseFlag {
			return filepath.SkipDir
		}

		configuration, err := parseFileDependingOnExtension(fileName)
		if err != nil {
			return err
		}

		if configuration != nil {
			configurations = append(configurations, &ConfigurationPair{
				TemplateFullyQualifiedPath: fileInfo.Name(),
				Configuration:              configuration,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	for _, entry := range configurations {
		if err := validateConfiguration(entry.Configuration); err != nil {
			return nil, err
		}
	}

	return configurations, nil
}
