package reportParser

// ReportParser defines the interface for parsers.
type ReportParser interface {
	LoadData(input string) error
	GenerateJSONData() (string, error)
}
