package env_helpers

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// TODO: Find an idiomatic way to determine whether it's temp debug build or executable
func GetRootPath() string {
	projGOPATH := strings.Split(os.Getenv("GOPATH"), ";")[0]
	if !strings.Contains(projGOPATH, "go-math-mapper") {
		exe, _ := os.Executable()
		exeDir := filepath.Dir(exe)
		return exeDir
	}
	return projGOPATH
}

func GetOsBitVersion() int {
	if strconv.IntSize == 32 {
		return 86
	}
	return 64
}