package resx

import _ "embed"

//go:embed resx_file.txt
var ResxFileTemplate string

//go:embed resx_file_header.txt
var ResxFileHeaderTemplate string
