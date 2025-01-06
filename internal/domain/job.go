package domain

type Job struct {
	Name        string `json:"jobName"`
	Jenkinsfile string `json:"jenkinsfile"`
}

type JobService interface {
	CreateJob(job Job) error
	TriggerJob(jobName string) error
	GetJobLogs(jobName string, buildNumber string) (string, error)
	CheckJobExists(jobName string) (bool, error)
	CheckJobRunning(jobName string) (bool, error)
}
