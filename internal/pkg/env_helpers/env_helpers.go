package env_helpers

import (
	"os"
	"strconv"
	"strings"
)

func GetGOPATH() string {
	return strings.Split(os.Getenv("GOPATH"), ";")[0]
}

func GetOsBitVersion() int {
	if strconv.IntSize == 32 {
		return 86
	}
	return 64
}