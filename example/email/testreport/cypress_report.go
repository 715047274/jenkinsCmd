package testreport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CypressReport struct {
	URL      string
	Stats    map[string]interface{}
	Failures []map[string]string
}

func (cr *CypressReport) FetchAndProcess() error {
	resp, err := http.Get(cr.URL)
	if err != nil {
		return fmt.Errorf("failed to fetch report: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var reportData struct {
		Stats   map[string]interface{} `json:"stats"`
		Results []struct {
			Title  string `json:"title"`
			Suites []struct {
				Title string `json:"title"`
				Tests []struct {
					Title     string `json:"title"`
					FullTitle string `json:"fullTitle"`
					State     string `json:"state"`
					Context   string `json:"context"`
					Duration  int    `json:"duration"`
				} `json:"tests"`
			} `json:"suites"`
		} `json:"results"`
	}

	if err := json.Unmarshal(body, &reportData); err != nil {
		return fmt.Errorf("failed to parse report JSON: %w", err)
	}

	cr.Stats = reportData.Stats
	cr.Failures = []map[string]string{}
	for _, result := range reportData.Results {
		for _, suite := range result.Suites {
			for _, test := range suite.Tests {
				if test.State == "failed" {
					cr.Failures = append(cr.Failures, map[string]string{
						"Suite": suite.Title,
						"Test":  test.FullTitle,
						"Error": test.Context,
					})
				}
			}
		}
	}

	return nil
}

func (cr *CypressReport) FetchAndParse() error {
	resp, err := http.Get(cr.URL)
	if err != nil {
		return fmt.Errorf("failed to fetch report: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read report data: %w", err)
	}

	var reportData struct {
		Stats   map[string]interface{} `json:"stats"`
		Results []struct {
			Title  string `json:"title"`
			Suites []struct {
				Title string `json:"title"`
				Tests []struct {
					Title     string `json:"title"`
					FullTitle string `json:"fullTitle"`
					State     string `json:"state"`
					Context   string `json:"context"`
					Duration  int    `json:"duration"`
				} `json:"tests"`
			} `json:"suites"`
		} `json:"results"`
	}

	if err := json.Unmarshal(body, &reportData); err != nil {
		return fmt.Errorf("failed to parse report JSON: %w", err)
	}

	cr.Stats = reportData.Stats
	cr.Failures = []map[string]string{}
	for _, result := range reportData.Results {
		for _, suite := range result.Suites {
			for _, test := range suite.Tests {
				if test.State == "failed" {
					context := test.Context
					// Ignore problematic context fields
					if len(context) > 0 && context[0] == '[' {
						context = "Error context is too complex to display."
					}
					cr.Failures = append(cr.Failures, map[string]string{
						"Suite": suite.Title,
						"Test":  test.FullTitle,
						"Error": context,
					})
				}
			}
		}
	}

	return nil
}

func (cr *CypressReport) GenerateJSONData() (string, error) {
	data := map[string]interface{}{
		"Stats":    cr.Stats,
		"Failures": cr.Failures, // Keep Failures as structured data
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to generate JSON data: %w", err)
	}

	return string(jsonData), nil
}
