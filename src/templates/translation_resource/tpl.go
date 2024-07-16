package translationresource

import _ "embed"

//go:embed class.txt
var ClassTemplate string

//go:embed interface.txt
var InterfaceTemplate string

//go:embed base_class.txt
var BaseClassTemplate string

//go:embed factory.txt
var FactoryTemplate string
