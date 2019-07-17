package math_facade

import (
	"../env_helpers"
	"fmt"
	"log"
	"syscall"
)

var math syscall.Handle

func init() {
	dllFile := fmt.Sprintf("%s/pkg/math/math_x%d.dll", env_helpers.GetRootPath(), env_helpers.GetOsBitVersion())
	lib, err := syscall.LoadLibrary(dllFile)
	if err != nil {
		context := fmt.Sprintf("load %s: ", dllFile)
		log.Fatal(context, err)
	}
	math = lib
}

func Div(arg1, arg2 int) int {
	proc, _ := syscall.GetProcAddress(math, "Div")
	divPtr, _, _ := syscall.Syscall(uintptr(proc), 0, uintptr(arg1), uintptr(arg2), 0)
	return int(divPtr)
}
