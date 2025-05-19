package configuration

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/nikita-petko/translation-resource-generator/flags"
	"github.com/nikita-petko/translation-resource-generator/models"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

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

var locales []string = []string{
	"ar_001",
	"de_de",
	"es_es",
	"fr_fr",
	"id_id",
	"it_it",
	"ja_jp",
	"ko_kr",
	"pl_pl",
	"pt_br",
	"ru_ru",
	"th_th",
	"tr_tr",
	"vi_vn",
	"zh_cjv",
	"zh_cn",
	"zh_tw",
}

var englishResxFileRegex *regexp.Regexp = regexp.MustCompile("^[a-zA-Z0-9_-]+\\.resx$")

func parseResXFiles(dirName, resourceNamespace string) (*models.Configuration, error) {
	resxEntries := make(map[string]map[string][]*models.ResxData)

	err := filepath.WalkDir(dirName, func(fileName string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if englishResxFileRegex.MatchString(path.Base(fileName)) {
			resourceName := strings.Split(path.Base(fileName), ".")[0]
			
			englishFile, err := os.Open(fileName)
			if err != nil {
				return err
			}

			var englishRoot models.ResxRoot
			xmlParser := xml.NewDecoder(englishFile)
			if err = xmlParser.Decode(&englishRoot); err != nil {
				return err
			}

			if resxEntries["en_us"] == nil { resxEntries["en_us"] = make(map[string][]*models.ResxData) }

			resxEntries["en_us"][resourceName] = englishRoot.Data

			englishFile.Close()

			for _, locale := range locales {
				localeResourceFileName := path.Join(dirName, fmt.Sprintf("%s.%s.resx", resourceName, locale))

				localeResourceFile, err := os.Open(localeResourceFileName)
				if errors.Is(err, os.ErrNotExist) {
					continue
				}
				if err != nil {
					return err
				}

				var localeRoot models.ResxRoot
				xmlParser := xml.NewDecoder(localeResourceFile)
				if err = xmlParser.Decode(&localeRoot); err != nil {
					return err
				}

				
				if resxEntries[locale] == nil { resxEntries[locale] = make(map[string][]*models.ResxData) }
				resxEntries[locale][resourceName] = localeRoot.Data

				localeResourceFile.Close()
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	configuration := &models.Configuration{
		Name: resourceNamespace,
		Resources: make(models.TranslationResourcesMap),
	}

	for locale, resourceMap := range resxEntries {
		for resourceName, entries := range resourceMap {
			stringMap, ok := configuration.Resources[resourceName]
			if !ok {
				stringMap = make(models.StringsMap)
			}

			for _, entry := range entries {
				resource, ok := stringMap[entry.NameAttr]
				if !ok {
					resource = &models.TranslationResources{
						Translations: make(models.StringMap),
					}
				}

				if locale == "en_us" {
					resource.EnglishString = entry.Value
				} else {
					resource.Translations[locale] = entry.Value
				}

				resource.Description = entry.Comment

				stringMap[entry.NameAttr] = resource
			}

			configuration.Resources[resourceName] = stringMap
		}
	}

	return configuration, nil
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
func Parse() ([]*models.Configuration, error) {
	if *flags.ConfigurationDirectoryFlag == "" {
		return nil, ErrConfigurationDirectoryNotSpecified
	}

	var configurations []*models.Configuration

	err := filepath.WalkDir(*flags.ConfigurationDirectoryFlag, func(fileName string, fileInfo fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if fileInfo.IsDir() && !*flags.RecurseFlag {
			return filepath.SkipDir
		}

		var configuration *models.Configuration

		if !*flags.FromResX {
			configuration, err = parseFileDependingOnExtension(fileName)
			if err != nil {
				return err
			}
		} else {
			if !strings.HasSuffix(fileName, ".resx") { return nil }
			if !englishResxFileRegex.MatchString(path.Base(fileName)) { return nil }

			dir := path.Dir(fileName)
			resourceNamespace := path.Base(dir)

			fmt.Printf("Parsing from ResX files %s, namespace = %s\n", fileName, resourceNamespace)

			configuration, err = parseResXFiles(dir, resourceNamespace)
			if err != nil {
				return err
			}
		}

		if configuration != nil {
			configurations = append(configurations, configuration)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	for _, entry := range configurations {
		if err := validateConfiguration(entry); err != nil {
			return nil, err
		}
	}

	return configurations, nil
}
