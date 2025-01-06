package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/715047274/jenkinsCmd/internal/config"
	"io"
	"net/http"
)

type JenkinsAdapter struct {
	cfg    config.Config
	client *http.Client
}

func NewJenkinsAdapter(config config.Config) *JenkinsAdapter {
	return &JenkinsAdapter{
		cfg:    config,
		client: &http.Client{},
	}
}

func (a *JenkinsAdapter) addBasicAuth(req *http.Request) {
	req.SetBasicAuth(a.cfg.GetJenkinsUser(), a.cfg.GetJenkinsToken())
}

func (a *JenkinsAdapter) CheckJobExists(jobName string) (bool, error) {
	url := fmt.Sprintf("%s/job/%s/api/json", a.cfg.GetJenkinsURL(), jobName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}
	a.addBasicAuth(req)

	resp, err := a.client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}

	return false, fmt.Errorf("unexpected response code: %d", resp.StatusCode)
}

func (a *JenkinsAdapter) CreateJob(jobName, jenkinsfile string) error {
	jobConfig := fmt.Sprintf(`
<flow-definition plugin="workflow-job">
    <definition class="org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition" plugin="workflow-cps">
        <script>%s</script>
        <sandbox>true</sandbox>
    </definition>
    <triggers/>
    <disabled>false</disabled>
</flow-definition>`, jenkinsfile)

	url := fmt.Sprintf("%s/createItem?name=%s", a.cfg.GetJenkinsURL(), jobName)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jobConfig)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/xml")
	a.addBasicAuth(req)

	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create job: %s", body)
	}

	return nil
}

func (a *JenkinsAdapter) TriggerJob(jobName string) error {
	url := fmt.Sprintf("%s/job/%s/build", a.cfg.GetJenkinsURL(), jobName)
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

func (a *JenkinsAdapter) GetBuildStatus(jobName string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/job/%s/api/json", a.cfg.GetJenkinsURL(), jobName)
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

func (a *JenkinsAdapter) GetBuildLogs(jobName, buildNumber string) (string, error) {
	url := fmt.Sprintf("%s/job/%s/%d/consoleText", a.cfg.GetJenkinsURL(), jobName, buildNumber)
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

func (a *JenkinsAdapter) DeleteJob(jobName string) error {
	url := fmt.Sprintf("%s/job/%s/doDelete", a.cfg.GetJenkinsURL(), jobName)
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
