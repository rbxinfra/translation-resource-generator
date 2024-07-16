package masterresources

import "github.com/nikita-petko/translation-resource-generator/flags"

// MasterResourcesModel is the model for the master resources.
type MasterResourcesModel struct {
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

	// Groups is the list of groups.
	Groups []string
}

// BuildModel builds the model for the template.
func BuildModel(toolName, version string, groups []string) (*MasterResourcesModel, error) {
	model := MasterResourcesModel{
		Version:         version,
		ToolName:        toolName,
		GlobalNamespace: *flags.NamespaceFlag,
		Groups:          groups,
	}

	return &model, nil
}
