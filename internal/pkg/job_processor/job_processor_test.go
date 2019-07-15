package job_processor

import "testing"
import "../json_hndl"

func TestProcessJobs_JobsLenEqualsReportsLen(t *testing.T) {
	jobs := make([]json_hndl.Job, 42)
	reports := ProcessJobs(&jobs)
	if len(jobs) != len(*reports) {
		t.Fail()
	}
}
