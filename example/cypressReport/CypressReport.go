package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CypressReport represents the structure of the Cypress report
type CypressReport struct {
	Stats   ReportStats     `json:"stats"`
	Results []ReportResults `json:"results"`
}

// ReportStats contains the test stats
type ReportStats struct {
	Suites         int     `json:"suites"`
	Tests          int     `json:"tests"`
	Passes         int     `json:"passes"`
	Pending        int     `json:"pending"`
	Failures       int     `json:"failures"`
	Skipped        int     `json:"skipped"`
	Duration       int     `json:"duration"`
	Start          string  `json:"start"`
	End            string  `json:"end"`
	PassPercent    float64 `json:"passPercent"`
	PendingPercent float64 `json:"pendingPercent"`
}

// ReportResults contains the individual test results
type ReportResults struct {
	Title  string  `json:"title"`
	Suites []Suite `json:"suites"`
}

// Suite represents a suite of tests
type Suite struct {
	Title string      `json:"title"`
	Tests []SuiteTest `json:"tests"`
}

// SuiteTest represents a single test in a suite
type SuiteTest struct {
	Title     string `json:"title"`
	FullTitle string `json:"fullTitle"`
	State     string `json:"state"`
	Context   string `json:"context"`
	Duration  int    `json:"duration"`
	Pass      bool   `json:"pass"`
	Fail      bool   `json:"fail"`
}

// FetchReport fetches the report from a network location and parses it into a CypressReport struct
func FetchReport(url string) (*CypressReport, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch report: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var report CypressReport
	if err := json.Unmarshal(body, &report); err != nil {
		return nil, fmt.Errorf("failed to parse report JSON: %w", err)
	}

	return &report, nil
}

// GenerateSummary generates a summary of the report
func (cr *CypressReport) GenerateSummary() string {
	summary := fmt.Sprintf(
		"Test Stats:\n- Suites: %d\n- Tests: %d\n- Passes: %d\n- Failures: %d\n- Skipped: %d\n- Duration: %dms\n\n",
		cr.Stats.Suites, cr.Stats.Tests, cr.Stats.Passes, cr.Stats.Failures, cr.Stats.Skipped, cr.Stats.Duration,
	)

	summary += "Failures:\n"
	for _, result := range cr.Results {
		for _, suite := range result.Suites {
			for _, test := range suite.Tests {
				if test.Fail {
					summary += fmt.Sprintf("- %s: %s\n", suite.Title, test.FullTitle)
				}
			}
		}
	}

	return summary
}
func main() {
	// Network location of the Cypress report
	reportURL := "http://nan4dfc1tst15.custadds.com:8080/job/Payroll_Intelligence_UI_Cypress_Test/95/execution/node/3/ws/cypress/reports/index.json"

	// Fetch and parse the report
	report, err := FetchReport(reportURL)
	if err != nil {
		log.Fatalf("Failed to fetch report: %v", err)
	}

	// Generate and print the summary
	summary := report.GenerateSummary()
	fmt.Println(summary)

	// Proceed with sending the summary via email if needed
	// ...
}
