package translationnamespacegroup

import (
	"github.com/nikita-petko/translation-resource-generator/flags"
	"github.com/nikita-petko/translation-resource-generator/models"
)

// TranslationNamespaceGroup is the model injected into the template.
type TranslationNamespaceGroup struct {
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

	// Resources is the map of resources.
	Resources []string
}

// BuildModel builds the model for the template.
func BuildModel(toolName, version string, config *models.Configuration) (*TranslationNamespaceGroup, error) {
	model := TranslationNamespaceGroup{
		Version:         version,
		ToolName:        toolName,
		GlobalNamespace: *flags.NamespaceFlag,
		Namespace:       config.Name,
		Resources:       make([]string, 0),
	}

	for key := range config.Resources {
		model.Resources = append(model.Resources, key)
	}

	return &model, nil
}
