package models

// Key is a translation.
type TranslationResources struct {
	// EnglishString is the english string of the translation. i.e. "Submit"
	EnglishString string `json:"englishString" yaml:"english_string" toml:"english_string" xml:"englishString"`

	// Description is the description of the translation. i.e. "This is the submit button".
	// This is optional.
	Description string `json:"description,omitempty" yaml:"description,omitempty" toml:"description,omitempty" xml:"description,omitempty"`

	// Translations is a key-value map of the translations. i.e. {"ru_ru": "Отправить"}
	// There is a case for the value here formatted like this:
	// "ru_ru": "Отправить {someVariable}"
	// This will be made into a format string:
	// string.Format("Отправить {0}", someVariable)
	// For translations that wish to use curly braces, they must be doubled up
	// i.e. "ru_ru": "Отправить {{someVariable}}" -> "Отправить {someVariable}"
	// "ru_ru": "Отправить {someVariable} {{someOtherVariable}}" -> "Отправить рубли {someOtherVariable}"
	Translations StringMap `json:"translations" yaml:"translations" toml:"translations" xml:"translations"`
}
