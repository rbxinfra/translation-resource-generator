package models

// TranslationsCdnContent is the content of the translations api.
// These are per-locale.
type TranslationsCdnContent struct {
	// Namespace is the namespace of the translations.
	Namespace string `json:"namespace"`

	// Key is the key of the translation.
	Key string `json:"key"`

	// Description is the description of the translation.
	Description string `json:"description,omitempty"`

	// English is the english string of the translation.
	English string `json:"english"`

	// Translation is the translation of the translation.
	Translation string `json:"translation"`
}

// TranslationsCdnResponse is the response from the translations cdn.
type TranslationsCdnResponse struct {
	// Contents is the contents of the translations.
	Contents []TranslationsCdnContent `json:"contents"`

	// Locale is the locale of the translations.
	Locale string `json:"locale"`
}
