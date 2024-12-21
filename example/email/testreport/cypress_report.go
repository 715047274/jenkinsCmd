package testreport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type CypressReport struct {
	URL      string
	Stats    map[string]interface{}
	Failures []map[string]string
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
					if len(context) > 0 && context[0] == '[' {
						parsedContext := []map[string]interface{}{}
						if err := json.Unmarshal([]byte(context), &parsedContext); err != nil {
							context = "Error context is too complex to display."
						} else {
							contextDetails := ""
							for _, detail := range parsedContext {
								for key, value := range detail {
									contextDetails += fmt.Sprintf("%s: %v\n", key, value)
								}
							}
							context = contextDetails
						}
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
	// Convert failures into a structured JSON-like table without rendering
	data := map[string]interface{}{
		"Stats":    cr.Stats,
		"Failures": cr.Failures, // Pass failures as raw data
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to generate JSON data: %w", err)
	}

	return string(jsonData), nil
}

func ReplaceTemplatePlaceholders(template string, parsedData map[string]interface{}) string {
	for key, value := range parsedData {
		switch v := value.(type) {
		case string:
			template = strings.ReplaceAll(template, fmt.Sprintf("{{.%s}}", key), v)
		case float64, int:
			template = strings.ReplaceAll(template, fmt.Sprintf("{{.%s}}", key), fmt.Sprintf("%v", v))
		case map[string]interface{}:
			for subKey, subValue := range v {
				template = strings.ReplaceAll(template, fmt.Sprintf("{{.%s.%s}}", key, subKey), fmt.Sprintf("%v", subValue))
			}
		}
	}
	return template
}
