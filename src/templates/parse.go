package templates

import (
	"fmt"
	"path"
	"regexp"
	"strings"
	"text/template"

	"github.com/nikita-petko/translation-resource-generator/models"
	masterresources "github.com/nikita-petko/translation-resource-generator/templates/master_resources"
	translationnamespacegroup "github.com/nikita-petko/translation-resource-generator/templates/translation_namespace_group"
	translationresource "github.com/nikita-petko/translation-resource-generator/templates/translation_resource"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func makeSafe(s string) string {
	// Can only be a-zA-Z0-9_ and must start with a-zA-Z or _, replace anything else with nothing.
	first := regexp.MustCompile("[^a-zA-Z0-9_]").ReplaceAllString(s, "")

	// If we start with a number, prefix with an underscore.
	if regexp.MustCompile("^[0-9]").MatchString(first) {
		first = "_" + first
	}

	return first
}

func escapeNewlines(s string) string {
	return strings.ReplaceAll(s, "\n", "\\n")
}

func escapeQuotes(s string) string {
	return strings.ReplaceAll(s, "\"", "\\\"")
}

func fixDescription(s string) string {
	// Remove trailing newlines.
	s = strings.TrimRight(s, "\n")

	return strings.ReplaceAll(s, "\n", "\n\t/// ")
}

func xmlEncode(s string) string {
	// Escape XML entities such as & and < and >.
	reg := regexp.MustCompile(`[&<>]`)
	return reg.ReplaceAllStringFunc(s, func(match string) string {
		// ampersand syntax
		if match == "&" {
			return "&amp;"
		}

		if match == "<" {
			return "&lt;"
		}

		if match == ">" {
			return "&gt;"
		}

		return match
	})
}

func toCamelCase(s string) string {
	// Remove all characters that are not alphanumeric or spaces or underscores
	s = regexp.MustCompile("[^a-zA-Z0-9_ ]+").ReplaceAllString(s, "")

	// Replace all underscores with spaces
	s = strings.ReplaceAll(s, "_", " ")

	// Title case s
	s = cases.Title(language.AmericanEnglish, cases.NoLower).String(s)

	// Remove all spaces
	s = strings.ReplaceAll(s, " ", "")

	// Lowercase the first letter
	if len(s) > 0 {
		s = strings.ToLower(s[:1]) + s[1:]
	}

	return s
}

func last(arr []string) string {
	return arr[len(arr)-1]
}

func applyCustomFunctions(tmpl *template.Template) {
	funcMap := template.FuncMap{
		"makeSafe":       makeSafe,
		"escapeNewlines": escapeNewlines,
		"fixDescription": fixDescription,
		"toCamelCase":    toCamelCase,
		"last":           last,
		"xmlEncode":      xmlEncode,
		"escapeQuotes":   escapeQuotes,
	}

	tmpl.Funcs(funcMap)
}

func executeTemplateForTranslationNamespaceGroupInterface(model *translationnamespacegroup.TranslationNamespaceGroup) (string, error) {
	tpl := template.New("translation_namespace_group_interface")

	applyCustomFunctions(tpl)

	var err error

	if tpl, err = tpl.Parse(translationnamespacegroup.InterfaceTemplate); err != nil {
		return "", err
	}

	var textWriter strings.Builder
	if err = tpl.Execute(&textWriter, model); err != nil {
		return "", err
	}

	return textWriter.String(), nil
}

func executeTemplateForTranslationNamespaceGroupImplementation(model *translationnamespacegroup.TranslationNamespaceGroup) (string, error) {
	tpl := template.New("translation_namespace_group_implementation")

	applyCustomFunctions(tpl)

	var err error

	if tpl, err = tpl.Parse(translationnamespacegroup.ClassTemplate); err != nil {
		return "", err
	}

	var textWriter strings.Builder
	if err = tpl.Execute(&textWriter, model); err != nil {
		return "", err
	}

	return textWriter.String(), nil
}

func executeTemplateForTranslationResourceInterface(model *translationresource.CommonTranslationResourceModel) (string, error) {
	tpl := template.New("translation_resource_interface")

	applyCustomFunctions(tpl)

	var err error

	if tpl, err = tpl.Parse(translationresource.InterfaceTemplate); err != nil {
		return "", err
	}

	var textWriter strings.Builder
	if err = tpl.Execute(&textWriter, model); err != nil {
		return "", err
	}

	return textWriter.String(), nil
}

func executeTemplateForTranslationResourceBaseClass(model *translationresource.CommonTranslationResourceModel) (string, error) {
	tpl := template.New("translation_resource_base_class")

	applyCustomFunctions(tpl)

	var err error

	if tpl, err = tpl.Parse(translationresource.BaseClassTemplate); err != nil {
		return "", err
	}

	var textWriter strings.Builder
	if err = tpl.Execute(&textWriter, model); err != nil {
		return "", err
	}

	return textWriter.String(), nil
}

func executeTemplateForTranslationResourceImplementation(model *translationresource.TranslatedResourceModel) (string, error) {
	tpl := template.New("translation_resource_implementation")

	applyCustomFunctions(tpl)

	var err error

	if tpl, err = tpl.Parse(translationresource.ClassTemplate); err != nil {
		return "", err
	}

	var textWriter strings.Builder
	if err = tpl.Execute(&textWriter, model); err != nil {
		return "", err
	}

	return textWriter.String(), nil
}

func executeTemplateForTranslationResourceFactory(model *translationresource.FactoryModel) (string, error) {
	tpl := template.New("translation_resource_factory")

	applyCustomFunctions(tpl)

	var err error

	if tpl, err = tpl.Parse(translationresource.FactoryTemplate); err != nil {
		return "", err
	}

	var textWriter strings.Builder
	if err = tpl.Execute(&textWriter, model); err != nil {
		return "", err
	}

	return textWriter.String(), nil
}

func executeTemplateForMasterResourcesInterface(model *masterresources.MasterResourcesModel) (string, error) {
	tpl := template.New("master_resources_interface")

	applyCustomFunctions(tpl)

	var err error

	if tpl, err = tpl.Parse(masterresources.InterfaceTemplate); err != nil {
		return "", err
	}

	var textWriter strings.Builder
	if err = tpl.Execute(&textWriter, model); err != nil {
		return "", err
	}

	return textWriter.String(), nil
}

func executeTemplateForMasterResourcesImplementation(model *masterresources.MasterResourcesModel) (string, error) {
	tpl := template.New("master_resources_implementation")

	applyCustomFunctions(tpl)

	var err error

	if tpl, err = tpl.Parse(masterresources.ClassTemplate); err != nil {
		return "", err
	}

	var textWriter strings.Builder
	if err = tpl.Execute(&textWriter, model); err != nil {
		return "", err
	}

	return textWriter.String(), nil
}

// ParseForMasterResources parses the templates for the master resources.
func ParseForMasterResources(toolName, version string, groups []string) (map[string]string, error) {
	files := make(map[string]string)

	fmt.Printf("Building master resources\n")

	model, err := masterresources.BuildModel(toolName, version, groups)
	if err != nil {
		return nil, err
	}

	masterResourcesInterface, err := executeTemplateForMasterResourcesInterface(model)
	if err != nil {
		return nil, err
	}

	masterResourcesInterfaceFile := path.Join("ResourceInterfaces", "IMasterResources.cs")
	files[masterResourcesInterfaceFile] = masterResourcesInterface

	masterResourcesImplementation, err := executeTemplateForMasterResourcesImplementation(model)
	if err != nil {
		return nil, err
	}

	masterResourcesImplementationFile := path.Join("ResourceImplementations", "MasterResources.cs")
	files[masterResourcesImplementationFile] = masterResourcesImplementation

	return files, nil
}

// ParseForConfiguration parses the templates for the configuration.
func ParseForConfiguration(toolName, version string, config *models.Configuration) (map[string]string, error) {
	files := make(map[string]string)

	model, err := translationnamespacegroup.BuildModel(toolName, version, config)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Building translation namespace group: %s\n", config.Name)

	translationNamespaceGroupInterface, err := executeTemplateForTranslationNamespaceGroupInterface(model)
	if err != nil {
		return nil, err
	}

	translationNamespaceGroupInterfaceFile := path.Join("ResourceInterfaces", fmt.Sprintf("I%sResources.cs", makeSafe(config.Name)))
	files[translationNamespaceGroupInterfaceFile] = translationNamespaceGroupInterface

	translationNamespaceGroupImplementation, err := executeTemplateForTranslationNamespaceGroupImplementation(model)
	if err != nil {
		return nil, err
	}

	translationNamespaceGroupImplementationFile := path.Join("ResourceImplementations", fmt.Sprintf("%sResources.cs", makeSafe(config.Name)))
	files[translationNamespaceGroupImplementationFile] = translationNamespaceGroupImplementation

	for name := range config.Resources {
		fmt.Printf("Building translation resource: %s.%s\n", config.Name, name)

		commonModel, err := translationresource.BuildCommonModel(toolName, version, name, config)
		if err != nil {
			return nil, err
		}

		translationResourceInterface, err := executeTemplateForTranslationResourceInterface(commonModel)
		if err != nil {
			return nil, err
		}

		translationResourceInterfaceFile := path.Join("ResourceInterfaces", config.Name, fmt.Sprintf("I%sResources.cs", makeSafe(name)))
		files[translationResourceInterfaceFile] = translationResourceInterface

		translationResourceBaseClass, err := executeTemplateForTranslationResourceBaseClass(commonModel)
		if err != nil {
			return nil, err
		}

		translationResourceBaseClassFile := path.Join("ResourceImplementations", config.Name, fmt.Sprintf("%sResources_en_us.cs", makeSafe(name)))
		files[translationResourceBaseClassFile] = translationResourceBaseClass

		for key := range config.Resources[name] {
			for locale := range config.Resources[name][key].Translations {
				fmt.Printf("Building translation entry, resource: %s.%s, key: %s, locale: %s\n", config.Name, name, key, locale)

				locale = strings.ToLower(locale)

				translatedModel, err := translationresource.BuildTranslatedModel(toolName, version, name, locale, config)
				if err != nil {
					return nil, err
				}

				translationResourceImplementation, err := executeTemplateForTranslationResourceImplementation(translatedModel)
				if err != nil {
					return nil, err
				}

				translationResourceImplementationFile := path.Join("ResourceImplementations", config.Name, fmt.Sprintf("%sResources_%s.cs", makeSafe(name), makeSafe(locale)))
				files[translationResourceImplementationFile] = translationResourceImplementation
			}
		}

		fmt.Printf("Building translation resource factory: %s.%s\n", config.Name, name)

		factoryModel, err := translationresource.BuildFactoryModel(toolName, version, name, config)
		if err != nil {
			return nil, err
		}

		translationResourceFactory, err := executeTemplateForTranslationResourceFactory(factoryModel)
		if err != nil {
			return nil, err
		}

		translationResourceFactoryFile := path.Join("ResourceFactories", config.Name, fmt.Sprintf("%sResourceFactory.cs", makeSafe(name)))
		files[translationResourceFactoryFile] = translationResourceFactory
	}

	return files, nil
}
