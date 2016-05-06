package main

// Flags represents the flags of the command line object.
type Flags struct {
	Depth            int
	InputPath        string
	LoaderFuncPrefix string
	OutputFileName   string
	Package          string
}
