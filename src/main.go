package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/nikita-petko/translation-resource-generator/configuration"
	"github.com/nikita-petko/translation-resource-generator/flags"
	"github.com/nikita-petko/translation-resource-generator/templates"
)

var applicationName string
var buildMode string
var commitSha string

// Pre-setup, runs before main.
func init() {
	flags.SetupFlags(applicationName, buildMode, commitSha)
}

// Main entrypoint.
func main() {
	if len(os.Args) == 1 {
		flag.Usage()

		return
	}

	if *flags.HelpFlag {
		flag.Usage()

		return
	}

	config, err := configuration.Parse()
	if err != nil {
		panic(err)
	}

	groups := make([]string, 0)

	for _, pair := range config {
		outputDirectory := *flags.OutputDirectoryFlag

		if err := os.MkdirAll(outputDirectory, os.ModePerm); err != nil {
			panic(err)
		}

		files, err := templates.ParseForConfiguration(applicationName, commitSha, pair.Configuration)
		if err != nil {
			panic(err)
		}

		for fileName, fileContents := range files {
			fileName = path.Join(outputDirectory, fileName)
			pathName := path.Dir(fileName)

			if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
				panic(err)
			}

			file, err := os.Create(fileName)
			if err != nil {
				panic(err)
			}

			defer file.Close()

			_, err = file.WriteString(fileContents)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Wrote file: %s\n", fileName)
		}

		groups = append(groups, pair.Configuration.Name)
	}

	files, err := templates.ParseForMasterResources(applicationName, commitSha, groups)
	if err != nil {
		panic(err)
	}

	for fileName, fileContents := range files {
		fileName = path.Join(*flags.OutputDirectoryFlag, fileName)
		pathName := path.Dir(fileName)

		if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
			panic(err)
		}

		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		_, err = file.WriteString(fileContents)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Wrote file: %s\n", fileName)
	}
}
