/*
 * =============================================================================
 * IBM Confidential
 * © Copyright IBM Corp. 2020, 2021
 *
 * The source code for this program is not published or otherwise divested of
 * its trade secrets, irrespective of what has been deposited with the
 * U.S. Copyright Office.
 * =============================================================================
 */

package zhmcclient

// JobAPI defines an interface for issuing Job requests to ZHMC
//go:generate counterfeiter -o fakes/job.go --fake-name JobAPI . JobAPI
type JobAPI interface {
	QueryJob(jobID string) (*Job, error)
	DeleteJob(jobID string) error
}

type JobStatus string

const (
	Running       JobStatus = "running"
	CancelPending           = "cancel-pending"
	Canceled                = "canceled"
	Complete                = "complete"
)

type Job struct {
	URI           string
	Status        JobStatus
	JobStatusCode int
	JobReasonCode int
	//JobResults    object
}

type JobManager struct {
	client ClientAPI
}

func NewJobManager(client ClientAPI) *JobManager {
	return &JobManager{
		client: client,
	}
}

/**
* GET /api/jobs/{job-id}
* Return: 200 and job status
*     or: 400, 404
 */
func (m *JobManager) QueryJob(jobID string) (*Job, error) {
	return nil, nil
}

/**
* DELETE /api/jobs/{job-id}
* Return: 204
*     or: 400, 404, 409
 */
func (m *JobManager) DeleteJob(jobID string) error {
	return nil
}
