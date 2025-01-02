package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config interface {
	GetJenkinsURL() string
	GetJenkinsUser() string
	GetJenkinsToken() string
}

type EnvConfig struct {
	jenkinsURL   string
	jenkinsUser  string
	jenkinsToken string
}

func NewEnvConfig(env string) Config {
	// Load the appropriate .env file
	envFile := env + ".env"
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}

	log.Printf("Loaded %s environment configuration", env)

	return &EnvConfig{
		jenkinsURL:   os.Getenv("JENKINS_URL"),
		jenkinsUser:  os.Getenv("JENKINS_USER"),
		jenkinsToken: os.Getenv("JENKINS_PASSWORD"),
	}
}

func (c *EnvConfig) GetJenkinsURL() string {
	return c.jenkinsURL
}

func (c *EnvConfig) GetJenkinsUser() string {
	return c.jenkinsUser
}

func (c *EnvConfig) GetJenkinsToken() string {
	return c.jenkinsToken
}
