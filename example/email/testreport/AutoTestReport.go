package testreport

import (
	"fmt"
)

// AutoTestReport represents the main structure for generating and sending test reports.
type AutoTestReport struct {
	MailClient    *MailClient
	MjmlHandler   *MjmlHandler
	CypressParser *CypressParser
}

// NewAutoTestReport initializes a AutoTestReport with the necessary clients.
func NewAutoTestReport(mailDomain, defaultTemplate string, defaultData map[string]string) *AutoTestReport {
	return &AutoTestReport{
		MailClient: &MailClient{Domain: mailDomain},
		MjmlHandler: &MjmlHandler{
			DefaultTemplate: defaultTemplate,
			DefaultData:     defaultData,
		},
		CypressParser: &CypressParser{},
	}
}

// GenerateAndSendReport processes the Cypress report, generates the HTML content, and sends the email.
func (tr *AutoTestReport) GenerateAndSendReport(input, sender, recipient, subject, attachmentPath string) error {
	// Load and process the Cypress report
	if err := tr.CypressParser.LoadData(input); err != nil {
		return fmt.Errorf("failed to load and process Cypress report: %w", err)
	}

	// Generate JSON data summarizing the report
	jsonData, err := tr.CypressParser.GenerateJSONData()
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
