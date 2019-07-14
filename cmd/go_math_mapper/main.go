package main

import (
	"../../internal/pkg/job_processor"
	"../../internal/pkg/json_hndl"
	"fmt"
)

func main() {
	json_hndl.HandleJSON(func(jobs []json_hndl.Job, output string) {
		report := job_processor.ProcessJobs(&jobs)
		json_hndl.WriteJSON(report, output)
		fmt.Println("Done")
	})
}
