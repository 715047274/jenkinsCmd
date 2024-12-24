package testreport

import (
	"fmt"
	"strings"
	"test-report/testreport/mailer"
	"test-report/testreport/mjmlHandler"
	"test-report/testreport/reportParser"
)

// AutoTestReport represents the main structure for generating and sending test reports.
type AutoTestReport struct {
	MailClient  *mailer.MailClient
	MjmlHandler *mjmlHandler.MjmlHandler
}

// NewAutoTestReport initializes a AutoTestReport with the necessary clients.
func NewAutoTestReport(mailDomain, defaultTemplate string, defaultData map[string]string) *AutoTestReport {
	return &AutoTestReport{
		MailClient: &mailer.MailClient{Domain: mailDomain},
		MjmlHandler: &mjmlHandler.MjmlHandler{
			DefaultTemplate: defaultTemplate,
			DefaultData:     defaultData,
		},
	}
}

// reportParserManager determines the appropriate parser based on the input.
func (tr *AutoTestReport) reportParserManager(input string) reportParser.ReportParser {
	projectNames := []string{
		"payroll_intelligence_ui_cypress_test",
		"Payroll_Intelligence_UI_Cypress_Test_Release_Branch",
		"sanity-test-payroll-ui",
		"Qa-Sanity-Test-Payroll-UI",
	}

	switch {
	case Some(projectNames, func(keyword string) bool {
		return strings.Contains(strings.ToLower(input), strings.ToLower(keyword))
	}):
		return &reportParser.CypressParser{}
	default:
		fmt.Printf("No suitable parser found for input: %s\n", input)
		return nil
	}
}

// GenerateAndSendReport processes the Cypress report, generates the HTML content, and sends the email.
func (tr *AutoTestReport) GenerateAndSendReport(input, sender, recipient, subject, attachmentPath string) error {
	// Load and process the Cypress report
	// Get the appropriate parser
	parser := tr.reportParserManager(input)

	if parser == nil {
		return fmt.Errorf("no suitable parser found for input: %s", input)
	}

	if err := parser.LoadData(input); err != nil {
		return fmt.Errorf("failed to load and process Cypress report: %w", err)
	}

	// Generate JSON data summarizing the report
	jsonData, err := parser.GenerateJSONData()
	if err != nil {
		return fmt.Errorf("failed to generate JSON data: %w", err)
	}

	// Generate the HTML content
	htmlContent, err := tr.MjmlHandler.CreateHTMLContent("", jsonData)
	if err != nil {
		return fmt.Errorf("failed to create HTML content: %w", err)
	}

	// Send the email
	return tr.MailClient.SendHTMLEmailWithAttachment(sender, recipient, subject, htmlContent, attachmentPath)
}

func Some(slice []string, predicate func(string) bool) bool {
	for _, item := range slice {
		if predicate(item) {
			return true
		}
	}
	return false
}
