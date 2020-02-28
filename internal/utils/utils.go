package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func CaseInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}

func GetDefaultConfigDir() string {
	return filepath.Join(Homedir(), ".gdown")
}

func Homedir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("APPDATA")
	}
	return os.Getenv("HOME")
}
