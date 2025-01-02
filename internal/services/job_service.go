package services

import (
	"github.com/715047274/jenkinsCmd/internal/adapters"
	"github.com/715047274/jenkinsCmd/internal/domain"
)

type jobService struct {
	adapter *adapters.JenkinsAdapter
}

func NewJobService(adapter *adapters.JenkinsAdapter) domain.JobService {
	return &jobService{adapter: adapter}
}

func (s *jobService) TriggerJob(jobName string) error {

	return s.adapter.TriggerJob(jobName)
}

func (s *jobService) GetJobLogs(jobName string, buildNumber int) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s *jobService) CheckJobExists(jobName string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *jobService) CheckJobRunning(jobName string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *jobService) CreateJob(job domain.Job) error {
	jobExists, err := s.adapter.CheckJobExists(job.Name)
	if err != nil {
		return err
	}
	if jobExists {
		return nil
	}
	return s.adapter.CreateJob(job.Name, job.Jenkinsfile)
}

//func (s *jobService) TriggerJob(jobName string) error {
//	return s.adapter.TriggerJob(jobName)
//}
