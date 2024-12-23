package testreport

import (
	"encoding/json"
	"testing"
)

func TestLoadData(t *testing.T) {
	// Sample JSON input
	jsonInput := `{
		"stats": {
			"suites": 5,
			"tests": 10,
			"passes": 8,
			"failures": 2,
			"pending": 0,
			"skipped": 1,
			"passPercent": 80.0,
			"pendingPercent": 0.0,
			"duration": 123,
			"start": "2024-12-21T08:53:17Z",
			"end": "2024-12-21T08:55:20Z"
		},
		"results": [
			{
				"title": "Suite 1",
				"suites": [
					{
						"title": "Sub Suite 1",
						"tests": [
							{
								"title": "Test 1",
								"fullTitle": "Suite 1 Sub Suite 1 Test 1",
								"state": "failed",
								"context": "Error context",
								"duration": 30
							}
						]
					}
				]
			}
		]
	}`

	cr := &CypressParser{}
	err := cr.LoadData(jsonInput)
	if err != nil {
		t.Fatalf("LoadData failed: %v", err)
	}

	// Verify Stats
	if cr.Stats["suites"] != 5 {
		t.Errorf("Expected 5 suites, got %v", cr.Stats["suites"])
	}
	if cr.Stats["passes"] != 8 {
		t.Errorf("Expected 8 passes, got %v", cr.Stats["passes"])
	}

	// Verify Failures
	if len(cr.Failures) != 1 {
		t.Errorf("Expected 1 failure, got %d", len(cr.Failures))
	} else {
		failure := cr.Failures[0]
		if failure["Suite"] != "Sub Suite 1" {
			t.Errorf("Expected Suite 'Sub Suite 1', got '%s'", failure["Suite"])
		}
		if failure["Test"] != "Suite 1 Sub Suite 1 Test 1" {
			t.Errorf("Expected Test 'Suite 1 Sub Suite 1 Test 1', got '%s'", failure["Test"])
		}
		if failure["Error"] != "Error context" {
			t.Errorf("Expected Error 'Error context', got '%s'", failure["Error"])
		}
	}
}

func TestGenerateJSONData(t *testing.T) {
	cr := &CypressParser{
		Stats: map[string]interface{}{
			"suites":         5,
			"tests":          10,
			"passes":         8,
			"failures":       2,
			"pending":        0,
			"skipped":        1,
			"passPercent":    80.0,
			"pendingPercent": 0.0,
			"duration":       123,
			"start":          "2024-12-21T08:53:17Z",
			"end":            "2024-12-21T08:55:20Z",
		},
		Failures: []map[string]string{
			{
				"Suite": "Sub Suite 1",
				"Test":  "Suite 1 Sub Suite 1 Test 1",
				"Error": "Error context",
			},
		},
	}

	jsonData, err := cr.GenerateJSONData()
	if err != nil {
		t.Fatalf("GenerateJSONData failed: %v", err)
	}

	// Parse the generated JSON to verify its content
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		t.Fatalf("Failed to parse generated JSON data: %v", err)
	}

	// Verify Stats
	stats := data["Stats"].(map[string]interface{})
	if stats["suites"].(float64) != 5 {
		t.Errorf("Expected 5 suites, got %v", stats["suites"])
	}
	if stats["passes"].(float64) != 8 {
		t.Errorf("Expected 8 passes, got %v", stats["passes"])
	}

	// Verify Failures
	failures := data["Failures"].([]interface{})
	if len(failures) != 1 {
		t.Errorf("Expected 1 failure, got %d", len(failures))
	} else {
		failure := failures[0].(map[string]interface{})
		if failure["Suite"] != "Sub Suite 1" {
			t.Errorf("Expected Suite 'Sub Suite 1', got '%s'", failure["Suite"])
		}
		if failure["Test"] != "Suite 1 Sub Suite 1 Test 1" {
			t.Errorf("Expected Test 'Suite 1 Sub Suite 1 Test 1', got '%s'", failure["Test"])
		}
		if failure["Error"] != "Error context" {
			t.Errorf("Expected Error 'Error context', got '%s'", failure["Error"])
		}
	}
}
