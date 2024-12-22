package testreport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type CypressReport struct {
	Stats    map[string]interface{}
	Failures []map[string]string
}

type Stats struct {
	Suites         int     `json:"suites"`
	Tests          int     `json:"tests"`
	Passes         int     `json:"passes"`
	Failures       int     `json:"failures"`
	Pending        int     `json:"pending"`
	Skipped        int     `json:"skipped"`
	PassPercent    float64 `json:"passPercent"`
	PendingPercent float64 `json:"pendingPercent"`
	Duration       int     `json:"duration"`
	Start          string  `json:"start"`
	End            string  `json:"end"`
}

type Test struct {
	Title     string `json:"title"`
	FullTitle string `json:"fullTitle"`
	State     string `json:"state"`
	Context   string `json:"context"`
	Err       Err
	Duration  int `json:"duration"`
}

type Err struct {
	Message string `json:"message"`
	Estack  string `json:"estack"`
}

type Suite struct {
	Title  string  `json:"title"`
	Tests  []Test  `json:"tests"`
	Suites []Suite `json:"suites"`
}

type Result struct {
	Title  string  `json:"title"`
	Suites []Suite `json:"suites"`
}

type ReportData struct {
	Stats   Stats    `json:"stats"`
	Results []Result `json:"results"`
}

func (cr *CypressReport) LoadData(input string) error {
	var data []byte
	var err error

	if strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://") {
		data, err = fetchFromURL(input)
		if err != nil {
			return err
		}
	} else {
		data, err = os.ReadFile(input)
		if err != nil {
			data = []byte(input)
		}
	}

	return cr.parseJSON(data)
}

func (cr *CypressReport) parseJSON(data []byte) error {
	var reportData ReportData
	if err := json.Unmarshal(data, &reportData); err != nil {
		return fmt.Errorf("failed to parse JSON data: %w", err)
	}

	// Populate Stats
	cr.Stats = map[string]interface{}{
		"suites":         reportData.Stats.Suites,
		"tests":          reportData.Stats.Tests,
		"passes":         reportData.Stats.Passes,
		"failures":       reportData.Stats.Failures,
		"pending":        reportData.Stats.Pending,
		"skipped":        reportData.Stats.Skipped,
		"passPercent":    reportData.Stats.PassPercent,
		"pendingPercent": reportData.Stats.PendingPercent,
		"duration":       reportData.Stats.Duration,
		"start":          reportData.Stats.Start,
		"end":            reportData.Stats.End,
	}

	// Initialize Failures
	cr.Failures = []map[string]string{}

	// Process Results Recursively
	for _, result := range reportData.Results {
		cr.processSuites(result.Suites, result.Title)
	}

	return nil
}

func (cr *CypressReport) processSuites(suites []Suite, parentTitle string) {
	for _, suite := range suites {
		// Process tests within the suite
		for _, test := range suite.Tests {
			if test.State == "failed" {
				context := cr.extractScreenshotPath(test.Context)
				if len(context) > 0 && context[0] == '[' {
					var parsedContext []map[string]interface{}
					if err := json.Unmarshal([]byte(context), &parsedContext); err != nil {
						context = "Error context is too complex to display."
					} else {
						var contextDetails strings.Builder
						for _, detail := range parsedContext {
							for key, value := range detail {
								contextDetails.WriteString(fmt.Sprintf("%s: %v\n", key, value))
							}
						}
						context = contextDetails.String()
					}
				}
				cr.Failures = append(cr.Failures, map[string]string{
					"Parent":     parentTitle,
					"Suite":      suite.Title,
					"Test":       test.FullTitle,
					"ScreenShot": context,
					"Error":      test.Err.Estack,
				})
			}
		}

		// Recursively process nested suites
		cr.processSuites(suite.Suites, suite.Title)
	}
}

func fetchFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

func (cr *CypressReport) GenerateJSONData() (string, error) {
	data := map[string]interface{}{
		"Stats":    cr.Stats,
		"Failures": cr.Failures,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to generate JSON data: %w", err)
	}

	return string(jsonData), nil
}

func (cr *CypressReport) extractScreenshotPath(context string) string {
	if len(context) > 0 && context[0] == '[' {
		var parsedContext []map[string]interface{}
		if err := json.Unmarshal([]byte(context), &parsedContext); err == nil {
			for _, detail := range parsedContext {
				if value, ok := detail["value"].(string); ok && strings.HasPrefix(value, "screenshots") && strings.HasSuffix(value, ".png") {
					return value
				}
			}
		}
	}
	return "No screenshot found"
}
