package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/xh3b4sd/loader/spec"
)

var (
	// version is the project version. It is given via buildflags that inject the
	// commit hash.
	version string
)

// Config represents the configuration used to create a new command line
// object.
type Config struct {
	// Settings.
	Cmd     *cobra.Command
	Flags   Flags
	Version string
}

// DefaultConfig provides a default configuration to create a new command line
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		Version: version,
	}

	return newConfig
}

// NewLoader creates a new configured command line object.
func NewLoader(config Config) (spec.Loader, error) {
	// loader
	newLoader := &loader{
		Config: config,
	}

	// command
	newLoader.Cmd = &cobra.Command{
		Use:   "loader",
		Short: "Asset management and code generation. For more information see https://github.com/xh3b4sd/loader",
		Long:  "Asset management and code generation. For more information see https://github.com/xh3b4sd/loader",

		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if newLoader.Flags.Depth < -1 {
				log.Fatalf("%#v\n", maskAnyf(invalidConfigError, "depth must be greater than -1"))
			}
			if newLoader.Flags.InputPath == "" {
				log.Fatalf("%#v\n", maskAnyf(invalidConfigError, "input path must not be empty"))
			}
			if newLoader.Flags.LoaderFuncPrefix == "" {
				log.Fatalf("%#v\n", maskAnyf(invalidConfigError, "loader function name must not be empty"))
			}
			if newLoader.Flags.OutputFileName == "" {
				log.Fatalf("%#v\n", maskAnyf(invalidConfigError, "output file name must not be empty"))
			}
			if newLoader.Flags.Package == "" {
				log.Fatalf("%#v\n", maskAnyf(invalidConfigError, "package must not be empty"))
			}

			// Calculate input depth.
			newLoader.InputDepth = newLoader.depthOf(newLoader.Flags.InputPath)
			fileInfo, err := os.Stat(newLoader.Flags.InputPath)
			if err != nil {
				log.Fatalf("%#v\n", maskAny(err))
			}
			if fileInfo.IsDir() && newLoader.Flags.InputPath[len(newLoader.Flags.InputPath)-1] != filepath.Separator {
				// In case the given input path represents a directory, but the given
				// input path does not contain a slash, we need to fix the input depth
				// explicitely by increasing the input depth by 1.
				newLoader.InputDepth++
			}
		},
	}

	// flags
	newLoader.Cmd.PersistentFlags().IntVarP(&newLoader.Flags.Depth, "depth", "d", 0, "depth of traversed directories (default 0)")
	newLoader.Cmd.PersistentFlags().StringVarP(&newLoader.Flags.InputPath, "input-path", "i", ".", "input path to load from")
	newLoader.Cmd.PersistentFlags().StringVarP(&newLoader.Flags.LoaderFuncPrefix, "loader-func-prefix", "l", "Loader", "prefix of the generated loader functions")
	newLoader.Cmd.PersistentFlags().StringVarP(&newLoader.Flags.OutputFileName, "output-file-name", "o", "loader.go", "name of the generated output file")
	newLoader.Cmd.PersistentFlags().StringVarP(&newLoader.Flags.Package, "package", "p", "main", "name of the generated output file")

	return newLoader, nil
}

func (l *loader) Boot() {
	// init
	l.Cmd.AddCommand(l.InitGenerateCmd())
	l.Cmd.AddCommand(l.InitVersionCmd())

	// execute
	l.Cmd.Execute()
}

type loader struct {
	Config

	// InputDepth describes the initial depth of the given input path. See the
	// following examples.
	//
	//     bar.ext          0
	//     foo/bar          1
	//     foo/bar.ext      1
	//     foo/bar/baz.ext  2
	//
	InputDepth int
}

func main() {
	newLoader, err := NewLoader(DefaultConfig())
	if err != nil {
		log.Fatalf("%#v\n", maskAny(err))
	}

	newLoader.Boot()
}
