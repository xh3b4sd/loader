package main

import (
	"path/filepath"
	"strings"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func (l *loader) matchesInputPath(inputPath string) bool {
	if l.Flags.Depth == -1 {
		return true
	}

	// Note that strings.Split always returns a slice containing at least one
	// item, that is the input string itself if the input could not be split.
	// Thus we need to subtract 1 for a proper depth check.
	if len(strings.Split(inputPath, string(filepath.Separator)))-1 > l.Flags.Depth {
		return false
	}

	return true
}
