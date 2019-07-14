package math_facade

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"syscall"
)

var math syscall.Handle

func init() {
	dllFile := fmt.Sprintf("internal/pkg/math_facade/math_x%d", getOsBitVersion())
	absPath, _ := filepath.Abs(dllFile)
	lib, err := syscall.LoadLibrary(absPath)
	if err != nil {
		log.Fatal("load: ", err)
	}
	math = lib
}

func Div(arg1, arg2 int) int {
	proc, _ := syscall.GetProcAddress(math, "Div")
	divPtr, _, _ := syscall.Syscall(uintptr(proc), 0, uintptr(arg1), uintptr(arg2), 0)
	return int(divPtr)
}

// TODO: Move to global helpers
func getOsBitVersion() int {
	if strconv.IntSize == 32 {
		return 86
	}
	return 64
}