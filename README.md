# Translation Resource Generator

Command line tool for parsing translations files and converting them into C# files used by [Roblox.TranslationResources](https://github.com/rbxinfra/translation-resources)

# Notice

## Usage of Roblox, or any of its assets.

# ***This project is not affiliated with Roblox Corporation.***

The usage of the name Roblox and any of its assets is purely for the purpose of providing a clear understanding of the project's purpose and functionality. This project is not endorsed by Roblox Corporation, and is not intended to be used for any commercial purposes.

Any code in this project was soley produced with or without the assistance of error traces and/or behaviour analysis of public facing APIs.

## Building

Ensure you have [Go 1.20.3+](https://go.dev/dl/)

1. Clone the repository via `git`:

    ```txt
    git clone git@github.com:rbxinfra/translation-resource-generator.git
    cd translation-resource-generator
    ```

2. Build via [make](https://www.gnu.org/software/make/)

    ```txt
    make build-debug
    ```

## Usage

`cd src && go run main.go --help` (use the build binary found in the bin directory if you downloaded a prebuilt or built it yourself)

```txt
Usage: translation-resource-generator
Build Mode: debug
Commit:  
        [-h|--help]
        [--config-to-resx] [--from-resx]
        [--configuration-directory <directory>] [--output-directory <directory>]
        [--recurse]
        [--namespace <namespace>]

  -config-to-resx
        Convert config files to ResX files
  -configuration-directory string
        The directory where the configuration files are located.
  -from-resx
        Is configuration coming from ResX files?
  -help
        Print usage.
  -namespace string
        The base namespace. (default "Roblox")
  -output-directory string
        The directory where the output files will be located. (default "./out")
  -recurse
        Recurse into subdirectories. (default true)
```

Example: 
translation-resource-generator --configuration-directory ./configurations --recurse
