package utils

import (
	"path/filepath"
	"strings"
)

func CleanPath(paths ...string) string {
	for i, p := range paths {
		if p != "" && !strings.HasSuffix(p, "/") && filepath.Ext(p) == "" {
			paths[i] = p + "/"
		}
	}

	joinedPath := filepath.Join(paths...)

	return joinedPath
}
