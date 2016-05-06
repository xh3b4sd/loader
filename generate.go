package main

import (
	"bytes"
	"compress/gzip"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

func (l *loader) InitGenerateCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate code containing the specified assets.",
		Long:  "Generate code containing the specified assets.",
		Run:   l.ExecGenerateCmd,
	}

	return newCmd
}

type tmplCtx struct {
	AssetMap         map[string][]byte
	LoaderFuncPrefix string
	MapName          string
	Package          string
}

func (l *loader) ExecGenerateCmd(cmd *cobra.Command, args []string) {
	newTmplCtx := tmplCtx{
		AssetMap:         map[string][]byte{},
		LoaderFuncPrefix: l.Flags.LoaderFuncPrefix,
		MapName:          "loaderGeneratedAssetMapVarName",
		Package:          l.Flags.Package,
	}

	err := filepath.Walk(l.Flags.InputPath, func(path string, info os.FileInfo, err error) error {
		if !l.matchesInputPath(path) {
			return nil
		}
		if info.IsDir() {
			return nil
		}

		raw, err := ioutil.ReadFile(path)
		if err != nil {
			return maskAny(err)
		}
		var b bytes.Buffer
		w := gzip.NewWriter(&b)
		_, err = w.Write(raw)
		w.Close()
		if err != nil {
			return maskAny(err)
		}
		newTmplCtx.AssetMap[path] = b.Bytes()

		return nil
	})

	if err != nil {
		log.Fatalf("%#v\n", maskAny(err))
	}

	// tmpl
	tmpl, err := template.New(l.Flags.OutputFileName).Parse(Template)
	if err != nil {
		log.Fatalf("%#v\n", maskAny(err))
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, newTmplCtx)
	if err != nil {
		log.Fatalf("%#v\n", maskAny(err))
	}

	// format
	raw, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatalf("%#v\n", maskAny(err))
	}

	err = ioutil.WriteFile(l.Flags.OutputFileName, raw, os.FileMode(0644))
	if err != nil {
		log.Fatalf("%#v\n", maskAny(err))
	}
}
