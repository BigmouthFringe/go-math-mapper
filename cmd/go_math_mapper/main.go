package main

import (
	"../../internal/pkg/json_hndl"
	"../../internal/pkg/math_facade"
	"fmt"
)

func main() {
	json_hndl.HandleJSON(func(jobs []json_hndl.Job, output string) {
		for _, job := range jobs {
			copyJob := job
			go func() {
				div, error := math_facade.Div(copyJob.Arg1, copyJob.Arg2)
				fmt.Println(div, " ", error)
			}()
		}
	})
	// TODO: Add output extracting here
}
