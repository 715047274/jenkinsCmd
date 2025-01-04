package handlers

import (
	"github.com/715047274/jenkinsCmd/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JobHandler struct {
	service domain.JobService
}

func NewJobHandler(service domain.JobService) *JobHandler {
	return &JobHandler{service: service}
}

func (h *JobHandler) RegisterRoutes(r *gin.Engine) {
	jenkinsGroup := r.Group("/jenkins")
	// jenkinsGroup.POST("/sequence", h.sequenceJob)
	jenkinsGroup.POST("/create", h.createJob)
	//jenkinsGroup.POST("/")
	//jenkinsGroup.POST("/")

}

//func (h *JobHandler) sequenceJob(c *gin.Context) {
//	// Retrieve the config from the Gin context
//	cfg, exists := c.MustGet("config").(config.Config)
//	if !exists {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Config not found"})
//		return
//	}
//
//	var req struct {
//		JobName     string `json:"jobName" binding:"required"`
//		Jenkinsfile string `json:"jenkinsfile" binding:"required"`
//	}
//
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	// Use the config (example: log the Jenkins URL)
//	jenkinsURL := cfg.GetJenkinsURL()
//	c.JSON(http.StatusOK, gin.H{"jenkinsURL": jenkinsURL})
//}

// Default Jenkinsfile content
const defaultJenkinsfile = `
pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo 'Building the project...'
            }
        }
        stage('Test') {
            steps {
                echo 'Running tests...'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying to production...'
            }
        }
    }
}`

func (h *JobHandler) createJob(c *gin.Context) {
	var req struct {
		JobName     string `json:"jobName"`
		Jenkinsfile string `json:"jenkinsfile"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Use default Jenkinsfile if none is provided
	if req.Jenkinsfile == "" {
		req.Jenkinsfile = defaultJenkinsfile
	}

	job := domain.Job{Name: req.JobName, Jenkinsfile: req.Jenkinsfile}
	if err := h.service.CreateJob(job); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job created successfully"})

}
func (h *JobHandler) triggerJob(c *gin.Context) {}
func (h *JobHandler) getStatus(c *gin.Context)  {}
func (h *JobHandler) getLogs(c *gin.Context)    {}
func (h *JobHandler) deleteJob(c *gin.Context)  {}
