package main

import (
	"github.com/715047274/jenkinsCmd/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadEnv() {
	// Check the ENV variable or default to "development"
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	// Load the appropriate .env file
	envFile := env + ".env"
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}

	log.Printf("Loaded %s environment configuration", env)
}

func main() {
	// Load environment variables
	loadEnv()

	// Verify Jenkins URL is set
	jenkinsURL := os.Getenv("JENKINS_URL")
	if jenkinsURL == "" {
		log.Fatalf("JENKINS_URL not set in environment")
	}
	log.Printf("Using Jenkins URL: %s", jenkinsURL)

	// Initialize the Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupJenkinsRoutes(r)

	// Start the server
	r.Run(":8282")
}
