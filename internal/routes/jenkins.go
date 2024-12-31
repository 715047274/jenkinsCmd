package routes

import (
	"github.com/715047274/jenkinsCmd/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupJenkinsRoutes(r *gin.Engine) {
	jenkinsGroup := r.Group("/jenkins")
	{
		jenkinsGroup.POST("/create", createJobHandler)
		jenkinsGroup.POST("/trigger", triggerJobHandler)
		jenkinsGroup.GET("/status/:jobName", getStatusHandler)
		jenkinsGroup.GET("/logs/:jobName/:buildNumber", getLogsHandler)
		jenkinsGroup.DELETE("/delete/:jobName", deleteJobHandler)
	}
}

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

func createJobHandler(c *gin.Context) {
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

	err := utils.CreateJob(req.JobName, req.Jenkinsfile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job created successfully"})
}

func triggerJobHandler(c *gin.Context) {
	var req struct {
		JobName string `json:"jobName" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := utils.TriggerJob(req.JobName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job triggered successfully"})
}

func getStatusHandler(c *gin.Context) {
	jobName := c.Param("jobName")

	status, err := utils.GetBuildStatus(jobName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, status)
}

func getLogsHandler(c *gin.Context) {
	jobName := c.Param("jobName")
	buildNumber := c.Param("buildNumber")

	logs, err := utils.GetBuildLogs(jobName, buildNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"logs": logs})
}

func deleteJobHandler(c *gin.Context) {
	jobName := c.Param("jobName")

	err := utils.DeleteJob(jobName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job deleted successfully"})
}
