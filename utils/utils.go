package utils

import (
	"path/filepath"
	"strings"
)

func IsYaml(fileName string) bool {
	extension := strings.ToLower(filepath.Ext(fileName))
	if extension == ".yaml" || extension == ".yml" {
		return true
	}
	return false
}
