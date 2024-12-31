package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var jenkinsURL = func() string {
	url := os.Getenv("JENKINS_URL")
	if url == "" {
		return "http://localhost:8080" // Default Jenkins URL
	}
	return url
}() // Replace with your Jenkins server URL

var jenkinsUser = "admin"
var jenkinsPassword = "11e9224b6f26f3c3896a00027c9a6d93fc"

// Helper function to add basic authentication to requests
func addBasicAuth(req *http.Request) {
	if jenkinsUser != "" && jenkinsPassword != "" {
		auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", jenkinsUser, jenkinsPassword)))
		req.Header.Set("Authorization", "Basic "+auth)
	}
}

func CreateJob(jobName, jenkinsfile string) error {
	jobConfig := fmt.Sprintf(`
<flow-definition plugin="workflow-job">
    <definition class="org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition" plugin="workflow-cps">
        <script>%s</script>
        <sandbox>true</sandbox>
    </definition>
    <triggers/>
    <disabled>false</disabled>
</flow-definition>`, jenkinsfile)

	url := fmt.Sprintf("%s/createItem?name=%s", jenkinsURL, jobName)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jobConfig)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/xml")
	addBasicAuth(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create job, status: %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}

func TriggerJob(jobName string) error {
	url := fmt.Sprintf("%s/job/%s/build", jenkinsURL, jobName)
	resp, err := http.Post(url, "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to trigger job, status: %d", resp.StatusCode)
	}

	return nil
}

func GetBuildStatus(jobName string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/job/%s/api/json", jenkinsURL, jobName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch status, status: %d", resp.StatusCode)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func GetBuildLogs(jobName, buildNumber string) (string, error) {
	url := fmt.Sprintf("%s/job/%s/%s/consoleText", jenkinsURL, jobName, buildNumber)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch logs, status: %d", resp.StatusCode)
	}

	logs, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(logs), nil
}

func DeleteJob(jobName string) error {
	url := fmt.Sprintf("%s/job/%s/doDelete", jenkinsURL, jobName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete job, status: %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}
