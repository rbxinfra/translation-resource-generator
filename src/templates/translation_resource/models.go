package translationresource

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nikita-petko/translation-resource-generator/flags"
	"github.com/nikita-petko/translation-resource-generator/models"
)

// StringModel is the string model.
type StringModel struct {
	// EnglishString is the english string.
	EnglishString string

	// Description is the description of the string.
	Description string

	// Params is a string of parameters (see comment in configuration.go)
	Parameters string

	// FormatString is the format string. Only used if there are parameters.
	FormatString string

	// FormatArguments is the format arguments. Only used if there are parameters.
	FormatArguments string
}

// CommonTranslationResourceModel is the common translation resource model.
type CommonTranslationResourceModel struct {
	//////////////////////////////////////////////
	// 		STANDARD TEMPLATE VARIABLES 		//
	//////////////////////////////////////////////

	// Version is the version of this tool.
	Version string

	// ToolName is the name of the tool.
	ToolName string

	// GlobalNamespace is the global namespace.
	GlobalNamespace string

	//////////////////////////////////////////////////
	// 		END STANDARD TEMPLATE VARIABLES 		//
	//////////////////////////////////////////////////

	// Namespace is the namespace of the translation.
	Namespace string

	// Name is the name of the translation.
	Name string

	// Strings is the map of strings.
	Strings map[string]StringModel
}

// TranslatedResourceModel is the model injected into the template.
type TranslatedResourceModel struct {
	//////////////////////////////////////////////
	// 		STANDARD TEMPLATE VARIABLES 		//
	//////////////////////////////////////////////

	// Version is the version of this tool.
	Version string

	// ToolName is the name of the tool.
	ToolName string

	// GlobalNamespace is the global namespace.
	GlobalNamespace string

	//////////////////////////////////////////////////
	// 		END STANDARD TEMPLATE VARIABLES 		//
	//////////////////////////////////////////////////

	// Namespace is the namespace of the translation.
	Namespace string

	// Name is the name of the translation.
	Name string

	// Locale is the locale of the translation.
	Locale string

	// Resource is the resource of the translation.
	Resource string

	// Strings is the map of strings.
	Strings map[string]TranslatedStringModel
}

// TranslatedStringModel is the model injected into the template.
type TranslatedStringModel struct {
	// EnglishString is the english string.
	EnglishString string

	// Description is the description of the string.
	Description string

	// Params is a string of parameters (see comment in configuration.go)
	Parameters string

	// FormatString is the format string. Only used if there are parameters.
	FormatString string

	// FormatArguments is the format arguments. Only used if there are parameters.
	FormatArguments string

	// Translation is the translation of the string.
	Translation string
}

// FactoryModel is the model injected into the template.
type FactoryModel struct {
	//////////////////////////////////////////////
	// 		STANDARD TEMPLATE VARIABLES 		//
	//////////////////////////////////////////////

	// Version is the version of this tool.
	Version string

	// ToolName is the name of the tool.
	ToolName string

	// GlobalNamespace is the global namespace.
	GlobalNamespace string

	//////////////////////////////////////////////////
	// 		END STANDARD TEMPLATE VARIABLES 		//
	//////////////////////////////////////////////////

	// Namespace is the namespace of the translation.
	Namespace string

	// Name is the name of the translation.
	Name string

	// Locales is the list of locales.
	Locales []string
}

func buildParameters(str string) string {
	// If the string is formatted as:
	// "abc {param1} def {param2} ghi", the params are "string param1, string param2".
	// If the string is formatted as:
	// "abc {{param1}} def {param2} ghi", the params are "string param2".
	// "abc {param1} def {param1} ghi {param2} jkl", the params are "string param1, string param2" (accounts for duplicates).

	paramsString := ""

	regex := `\{[a-zA-Z0-9]+\}`
	matches := regexp.MustCompile(regex).FindAllString(str, -1)

	if len(matches) == 0 {
		return paramsString
	}

	// Remove duplicates.
	uniqueMatches := make(map[string]bool)
	for _, match := range matches {
		if _, ok := uniqueMatches[match]; !ok {
			uniqueMatches[match] = true
		}
	}

	matches = make([]string, 0)
	for match := range uniqueMatches {
		matches = append(matches, match)
	}

	for i, match := range matches {
		if i != 0 {
			paramsString += ", "
		}

		paramsString += fmt.Sprintf("string %s", match[1:len(match)-1])
	}

	return paramsString
}

func buildFormatString(str string) string {
	// If the string is formatted as:
	// "abc {param1} def {param2} ghi", the format string is "abc {0} def {1} ghi".
	// If the string is formatted as:
	// "abc {{param1}} def {param2} ghi", the format string is "abc {0} def {1} ghi".
	// "abc {param1} def {param1} ghi {param2} jkl", the format string is "abc {0} def {0} ghi {1} jkl".
	// "abc {param1} def {param2} ghi {param1} jkl", the format string is "abc {0} def {1} ghi {0} jkl".

	formatString := str

	regex := `\{[a-zA-Z0-9]+\}`
	matches := regexp.MustCompile(regex).FindAllString(str, -1)

	if len(matches) == 0 {
		return formatString
	}

	// Remove duplicates.
	uniqueMatches := make(map[string]bool)
	for _, match := range matches {
		if _, ok := uniqueMatches[match]; !ok {
			uniqueMatches[match] = true
		}
	}

	matches = make([]string, 0)
	for match := range uniqueMatches {
		matches = append(matches, match)
	}

	for i, match := range matches {
		formatString = strings.ReplaceAll(formatString, match, fmt.Sprintf("{%d}", i))
	}

	return formatString
}

func buildFormatArguments(str string) string {
	// If the string is formatted as:
	// "abc {param1} def {param2} ghi", the format arguments are "param1, param2".
	// If the string is formatted as:
	// "abc {{param1}} def {param2} ghi", the format arguments are "param2".
	// "abc {param1} def {param1} ghi {param2} jkl", the format arguments are "param1, param2" (accounts for duplicates).

	formatArguments := ""

	regex := `\{[a-zA-Z0-9]+\}`
	matches := regexp.MustCompile(regex).FindAllString(str, -1)

	if len(matches) == 0 {
		return formatArguments
	}

	// Remove duplicates.
	uniqueMatches := make(map[string]bool)
	for _, match := range matches {
		if _, ok := uniqueMatches[match]; !ok {
			uniqueMatches[match] = true
		}
	}

	matches = make([]string, 0)
	for match := range uniqueMatches {
		matches = append(matches, match)
	}

	for i, match := range matches {
		if i != 0 {
			formatArguments += ", "
		}

		formatArguments += match[1 : len(match)-1]
	}

	return formatArguments
}

// BuildCommonModel builds the common model for the template.
func BuildCommonModel(toolName, version, name string, config *models.Configuration) (*CommonTranslationResourceModel, error) {
	model := CommonTranslationResourceModel{
		Version:         version,
		ToolName:        toolName,
		GlobalNamespace: *flags.NamespaceFlag,
		Namespace:       config.Name,
		Name:            name,
		Strings:         make(map[string]StringModel),
	}

	for key, value := range config.Resources[name] {
		model.Strings[key] = StringModel{
			EnglishString:   value.EnglishString,
			Description:     value.Description,
			Parameters:      buildParameters(value.EnglishString),
			FormatString:    buildFormatString(value.EnglishString),
			FormatArguments: buildFormatArguments(value.EnglishString),
		}
	}

	return &model, nil
}

// BuildTranslatedModel builds the translated model for the template.
func BuildTranslatedModel(toolName, version, name, locale string, config *models.Configuration) (*TranslatedResourceModel, error) {
	model := TranslatedResourceModel{
		Version:         version,
		ToolName:        toolName,
		GlobalNamespace: *flags.NamespaceFlag,
		Namespace:       config.Name,
		Name:            name,
		Locale:          locale,
		Resource:        name,
		Strings:         make(map[string]TranslatedStringModel),
	}

	for key, value := range config.Resources[name] {
		model.Strings[key] = TranslatedStringModel{
			EnglishString:   value.EnglishString,
			Description:     value.Description,
			Parameters:      buildParameters(value.EnglishString),
			FormatString:    buildFormatString(value.Translations[locale]),
			FormatArguments: buildFormatArguments(value.EnglishString),
			Translation:     value.Translations[locale],
		}
	}

	return &model, nil
}

func determineLocales(resource map[string]*models.TranslationResources) []string {
	// Go through each string, compile together a list of their translations.
	// Then, go through each translation, and if it's not in the list, add it.
	// This will give us a list of all the locales.

	locales := make([]string, 0)

	for _, value := range resource {
		for locale := range value.Translations {
			locales = append(locales, locale)
		}
	}

	// Remove duplicates.
	localeMap := make(map[string]bool)
	for _, locale := range locales {
		localeMap[locale] = true
	}

	locales = make([]string, 0)
	for locale := range localeMap {
		locale = strings.ToLower(locale)
		locales = append(locales, locale)
	}

	return locales
}

// BuildFactoryModel builds the factory model for the template.
func BuildFactoryModel(toolName, version, name string, config *models.Configuration) (*FactoryModel, error) {
	model := FactoryModel{
		Version:         version,
		ToolName:        toolName,
		GlobalNamespace: *flags.NamespaceFlag,
		Namespace:       config.Name,
		Name:            name,
		Locales:         make([]string, 0),
	}

	model.Locales = append(model.Locales, determineLocales(config.Resources[name])...)

	return &model, nil
}
