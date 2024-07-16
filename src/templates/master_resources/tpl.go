package masterresources

import _ "embed"

//go:embed interface.txt
var InterfaceTemplate string

//go:embed class.txt
var ClassTemplate string
