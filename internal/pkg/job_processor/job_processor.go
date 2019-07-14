package job_processor

import (
	"../json_hndl"
	"../math_facade"
	"fmt"
	"sync"
)

// This function calculates div for every job concurrently and aggregates
// them in one report which is basically an array of processed job summaries.
func ProcessJobs(jobs *[]json_hndl.Job) *[]json_hndl.JobReport {
	wg := &sync.WaitGroup{}
	wg.Add(len(*jobs))

	reports := make([]json_hndl.JobReport, len(*jobs))
	for i, job := range *jobs {
		copyIndex, copyJob := i, job
		go func() {
			div := math_facade.Div(copyJob.Arg1, copyJob.Arg2)
			name := fmt.Sprintf("Job #%d", copyIndex+1)
			reports[copyIndex] = json_hndl.JobReport{Title: name, Div: div}
			wg.Done()
		}()
	}
	wg.Wait()
	return &reports
}
