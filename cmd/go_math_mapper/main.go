package main

import (
	"../../internal/pkg/json_hndl"
	"fmt"
)

func main() {
	var outFile string
	json_hndl.HandleJSON(func(jobs []json_hndl.Job, output string) {
		outFile = output
		fmt.Println(jobs, outFile)
	})
	// TODO: Add input mapping here
	// TODO: Add output extracting here
}