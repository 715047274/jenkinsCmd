package testreport

import (
	"fmt"
)

// TestReport represents the main structure for generating and sending test reports.
type TestReport struct {
	MailClient    *MailClient
	MJMLClient    *MJMLClient
	CypressReport *CypressReport
}

// NewTestReport initializes a TestReport with the necessary clients.
func NewTestReport(mailDomain, defaultTemplate string, defaultData map[string]string) *TestReport {
	return &TestReport{
		MailClient: &MailClient{Domain: mailDomain},
		MJMLClient: &MJMLClient{
			DefaultTemplate: defaultTemplate,
			DefaultData:     defaultData,
		},
		CypressReport: &CypressReport{},
	}
}

// GenerateAndSendReport processes the Cypress report, generates the HTML content, and sends the email.
func (tr *TestReport) GenerateAndSendReport(input, sender, recipient, subject, attachmentPath string) error {
	// Load and process the Cypress report
	if err := tr.CypressReport.LoadData(input); err != nil {
		return fmt.Errorf("failed to load and process Cypress report: %w", err)
	}

	// Generate JSON data summarizing the report
	jsonData, err := tr.CypressReport.GenerateJSONData()
	if err != nil {
		return fmt.Errorf("failed to generate JSON data: %w", err)
	}

	// Generate the HTML content
	htmlContent, err := tr.MJMLClient.CreateHTMLContent("", jsonData)
	if err != nil {
		return fmt.Errorf("failed to create HTML content: %w", err)
	}

	// Send the email
	return tr.MailClient.SendHTMLEmailWithAttachment(sender, recipient, subject, htmlContent, attachmentPath)
}
