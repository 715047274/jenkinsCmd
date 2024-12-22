package test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"test-report/testreport"
	"testing"
)

// MockFileReader is a mock implementation of the FileReader interface.
type MockFileReader struct {
	Data  []byte
	Error error
}

func (m MockFileReader) ReadFile(filename string) ([]byte, error) {
	return m.Data, m.Error
}

// MockHTTPFetcher is a mock implementation of the HTTPFetcher interface.
type MockHTTPFetcher struct {
	Response *http.Response
	Error    error
}

func (m MockHTTPFetcher) Get(url string) (*http.Response, error) {
	return m.Response, m.Error
}

func TestLoadData_FromFile(t *testing.T) {
	mockData := `{"stats": {"suites": 1, "tests": 1, "passes": 1, "failures": 0}, "results": []}`
	cr := &testreport.CypressReport{
		FileReader: MockFileReader{
			Data:  []byte(mockData),
			Error: nil,
		},
	}

	err := cr.LoadData("dummy/path/to/file.json")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cr.Stats["suites"] != 1 {
		t.Errorf("expected 1 suite, got %v", cr.Stats["suites"])
	}
}

func TestLoadData_FromURL(t *testing.T) {
	mockData := `{"stats": {"suites": 1, "tests": 1, "passes": 1, "failures": 0}, "results": []}`
	cr := &testreport.CypressReport{
		HTTPFetcher: MockHTTPFetcher{
			Response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader(mockData)),
			},
			Error: nil,
		},
	}

	err := cr.LoadData("http://example.com/report.json")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cr.Stats["suites"] != 1 {
		t.Errorf("expected 1 suite, got %v", cr.Stats["suites"])
	}
}

func TestLoadData_InvalidJSON(t *testing.T) {
	cr := &testreport.CypressReport{}

	err := cr.LoadData(`{"invalid_json": }`)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestLoadData_FileNotFound(t *testing.T) {
	cr := &testreport.CypressReport{
		FileReader: MockFileReader{
			Data:  nil,
			Error: errors.New("file not found"),
		},
	}

	err := cr.LoadData("nonexistentfile.json")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}
