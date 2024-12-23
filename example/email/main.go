package main

import (
	"fmt"
	"log"
	"test-report/testreport"
)

func buildEmail(reportLink string) string {
	header := buildHeader()
	table := buildTable(reportLink)
	errorBoard := buildErrorBoard(reportLink)

	return fmt.Sprintf(`
	<mjml>
	  <mj-body>
	    %s
	    %s
	    %s
	  </mj-body>
	</mjml>
	`, header, table, errorBoard)
}

func buildErrorBoard(reportLink string) string {
	return fmt.Sprintf(`
	<mj-wrapper>
	  {{ range .Failures }}
	  <mj-section>
	    <mj-column>
	      <mj-text>
	        <p><strong>Suite:</strong> {{ .Suite }}</p>
	        <p><strong>Test:</strong> {{ .Test }}</p>
	        <p><strong>Error:</strong> {{ .Error }}</p>
	        <p><strong>Screenshot:</strong> 
	          <a href="%s/{{ .ScreenShot }}">
	            View Screenshot
	          </a>
	        </p>
	      </mj-text>
	    </mj-column>
	  </mj-section>
	  {{ end }}
	</mj-wrapper>
	`, reportLink)
}

func buildTable(reportLink string) string {
	return fmt.Sprintf(`
	<mj-section>
	  <mj-column>
	    <mj-table>
	      <tr>
	        <th>Metric</th>
	        <th>Value</th>
	      </tr>
	      <tr>
	        <td>Suites</td>
	        <td>{{ index .Stats "suites" }}</td>
	      </tr>
	      <tr>
	        <td>Tests</td>
	        <td>{{ index .Stats "tests" }}</td>
	      </tr>
	      <tr>
	        <td>Passes</td>
	        <td>{{ index .Stats "passes" }}</td>
	      </tr>
	      <tr>
	        <td>Failures</td>
	        <td>{{ index .Stats "failures" }}</td>
	      </tr>
	      <tr>
	        <td>Pending</td>
	        <td>{{ index .Stats "pending" }}</td>
	      </tr>
	      <tr>
	        <td>Skipped</td>
	        <td>{{ index .Stats "skipped" }}</td>
	      </tr>
	    </mj-table>
	    <mj-button background-color="#3067DB" href="%s">
	      Report Link
	    </mj-button>
	  </mj-column>
	</mj-section>
	`, reportLink)
}

func buildHeader() string {
	return `
	<mj-section>
	  <mj-column>
	    <mj-text font-size="20px" color="#333333" align="center">Cypress Test Report</mj-text>
	    <mj-divider border-color="#F45E43"></mj-divider>
	  </mj-column>
	</mj-section>
	`
}

func main() {
	// Configuration
	buildNum := "98"
	reportHost := "http://nan4dfc1tst15.custadds.com:8080/job/Payroll_Intelligence_UI_Cypress_Test/"
	reportLink := fmt.Sprintf("%s%s/payroll-intelliigence-ui", reportHost, buildNum)
	reportInput := fmt.Sprintf("%s%s/execution/node/3/ws/cypress/reports/index.json", reportHost, buildNum)

	// Generate the email template
	emailTemplate := buildEmail(reportLink)

	// Email configuration
	mailDomain := "corpadds.com"
	sender := "autotest@yourdomain.com"
	recipient := "k.zhang@dayforce.com"
	subject := fmt.Sprintf("Cypress Test Report - Build %s", buildNum)
	attachmentPath := ""

	// Initialize AutoTestReport
	autoTestReport := testreport.NewAutoTestReport(mailDomain, emailTemplate, nil)

	// Generate and send the report
	err := autoTestReport.GenerateAndSendReport(reportInput, sender, recipient, subject, attachmentPath)
	if err != nil {
		log.Fatalf("Failed to generate and send the report: %v", err)
	}

	fmt.Printf("Test report for build %s sent successfully!\n", buildNum)
}
