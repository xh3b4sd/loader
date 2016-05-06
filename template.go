package main

// Template represents the template of the file being generated. It contains
// the asset map and methods to access them.
var Template = `package {{.Package}}

import (
	"fmt"
)

// {{.MapName}} contains all loaded files. Map keys are file names. Map values
// are file contents.
var {{.MapName}} = map[string]string{{"{"}}{{range $fileName, $content := .AssetMap}}
		"{{$fileName}}": {{printf "%q" $content}},{{end}}
}

// {{.LoaderFuncPrefix}}ReadFile returns the content of the loaded asset
// associated to assetName. If the requested asset does not exist, an error is
// returned.
func {{.LoaderFuncPrefix}}ReadFile(assetName string) ([]byte, error) {
	if s, ok := {{.MapName}}[assetName]; ok {
		return []byte(s), nil
	}

	return nil, fmt.Errorf("asset not found: %s", assetName)
}

// {{.LoaderFuncPrefix}}FileNames returns a list of file names of all loaded
// assets.
func {{.LoaderFuncPrefix}}FileNames() []string {
	var newFileNames []string

	for fileName := range {{.MapName}} {
		newFileNames = append(newFileNames, fileName)
	}

	return newFileNames
}`
