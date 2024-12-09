package utils

import "strings"

func GetParentPaths(path string) []string {
	path = strings.TrimSuffix(path, "/")

	var paths []string
	components := strings.Split(path, "/")

	for i := 1; i <= len(components); i++ {
		paths = append(paths, strings.Join(components[:i], "/"))
	}

	return paths
}
