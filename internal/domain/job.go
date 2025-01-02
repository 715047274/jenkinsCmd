package domain

type Job struct {
	Name        string
	Jenkinsfile string
}

type JobService interface {
	CreateJob(job Job) error
	TriggerJob(jobName string) error
	GetJobLogs(jobName string, buildNumber int) (string, error)
	CheckJobExists(jobName string) (bool, error)
	CheckJobRunning(jobName string) (bool, error)
}
