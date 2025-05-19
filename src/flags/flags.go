package flags

import "flag"

var (
	// HelpFlag prints the usage.
	HelpFlag = flag.Bool("help", false, "Print usage.")

	// ConfigToResX converts config files to ResX files
	ConfigToResX = flag.Bool("config-to-resx", false, "Convert config files to ResX files")

	// FromResX determines if config is coming from ResX files
	FromResX = flag.Bool("from-resx", false, "Is configuration coming from ResX files?")

	// ConfigurationDirectoryFlag is the directory where the configuration files are located.
	ConfigurationDirectoryFlag = flag.String("configuration-directory", "", "The directory where the configuration files are located.")

	// OutputDirectoryFlag is the directory where the output files will be located.
	OutputDirectoryFlag = flag.String("output-directory", "./out", "The directory where the output files will be located.")

	// RecurseFlag indicates whether to recurse into subdirectories.
	RecurseFlag = flag.Bool("recurse", true, "Recurse into subdirectories.")

	// NamespaceFlag is the base namespace.
	NamespaceFlag = flag.String("namespace", "Roblox", "The base namespace.")
)

const FlagsUsageString string = `
	[-h|--help]
	[--config-to-resx]
	[--configuration-directory <directory>] [--output-directory <directory>]
	[--recurse]
	[--namespace <namespace>]`
