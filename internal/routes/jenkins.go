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
		// jenkinsGroup.POST("/sequence", sequenceJobHandler)

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

//func sequenceJobHandler(c *gin.Context) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var req struct {
//			JobName     string `json:"jobName" binding:"required"`
//			Jenkinsfile string `json:"jenkinsfile" binding:"required"`
//		}
//
//		if err := c.ShouldBindJSON(&req); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		// Step 1: Check if the job exists
//		jobExists, err := utils.JobExists(cfg, req.JobName)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to check job existence: %v", err)})
//			return
//		}
//
//		if !jobExists {
//			// Step 2: Create the job if it doesn't exist
//			err := utils.CreateJob(cfg, req.JobName, req.Jenkinsfile)
//			if err != nil {
//				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create job: %v", err)})
//				return
//			}
//		}
//
//		// Step 3: Check if the job is running
//		isRunning, err := utils.IsJobRunning(cfg, req.JobName)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to check job running status: %v", err)})
//			return
//		}
//
//		if !isRunning {
//			// Step 4: Trigger the job if it's not already running
//			err = utils.TriggerJob(cfg, req.JobName)
//			if err != nil {
//				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to trigger job: %v", err)})
//				return
//			}
//		}
//
//		// Step 5: Get the build number
//		buildNumber, err := utils.GetLastBuildNumber(cfg, req.JobName)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get last build number: %v", err)})
//			return
//		}
//
//		// Step 6: Poll for logs asynchronously
//		resultChan := make(chan string)
//		errorChan := make(chan error)
//
//		go utils.WaitForBuildAndGetLogs(cfg, req.JobName, buildNumber, resultChan, errorChan)
//
//		select {
//		case logs := <-resultChan:
//			c.JSON(http.StatusOK, gin.H{
//				"message": "Job executed successfully",
//				"logs":    logs,
//			})
//		case err := <-errorChan:
//			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to execute job: %v", err)})
//		}
//	}
//}
