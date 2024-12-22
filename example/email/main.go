package main

import (
	"fmt"
	"log"
	"test-report/testreport"
)

func main() {
	// Define the input for the Cypress report
	// This can be a URL, a local file path, or a JSON string
	reportInput := "http://nan4dfc1tst15.custadds.com:8080/job/Payroll_Intelligence_UI_Cypress_Test/95/execution/node/3/ws/cypress/reports/index.json"
	// Initialize the CypressReport struct
	report := testreport.CypressReport{}

	// Load and process the Cypress report
	if err := report.LoadData(reportInput); err != nil {
		log.Fatalf("Failed to load and process Cypress report: %v", err)
	}

	// Generate JSON data summarizing the report
	jsonData, err := report.GenerateJSONData()
	if err != nil {
		log.Fatalf("Failed to generate JSON data: %v", err)
	}

	// Print the JSON data to verify the output
	fmt.Println("Generated JSON Data:")
	fmt.Println(jsonData)

	// Initialize the MJML template and email details
	mjmlClient := testreport.MJMLClient{
		DefaultTemplate: `
<mjml>
  <mj-body>
    <mj-section>
      <mj-column>
		<mj-image width="600px" src="https://storage-thumbnails.bananatag.com/images/zsfKYb/1b35331be2d0a05a7d6ce2531ebc2ab4.png" />
        <mj-divider border-color="#F45E43"></mj-divider>
        <mj-text font-size="16px" font-family="helvetica" color="#555555">Cypress Test Report:</mj-text>
        <mj-table>
          <tr style="background-color:#f0f0f0;text-align:left;">
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
            <td style="color:green">{{ index .Stats "passes" }}</td>
          </tr>
          <tr>
            <td>Failures</td>
            <td style="color:red">{{ index .Stats "failures" }}</td>
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
        <mj-button background-color="#3067DB" href="http://nan4dfc1tst15.custadds.com:8080/job/Payroll_Intelligence_UI_Cypress_Test/97/payroll-intelliigence-ui">Report Link</mj-button>

        <mj-text font-size="16px" font-family="helvetica" color="#555555">Failure Details:</mj-text>
        <mj-accordion>
         {{ range .Failures }}
 	     <mj-accordion-element>
          <mj-accordion-title>Suite:</strong> {{ .Suite }}</mj-accordion-title>
		  <mj-accordion-text>
          <span style="line-height:20px; font-size:14px;">
          <strong>Test:</strong> {{ .Test }}<br/>
          <strong>Error Stack:</strong> {{ .Error }}<br/>
		  </span>	
          </mj-accordion-text>
          </mj-accordion-element>
          {{ end }}
       </mj-accordion>
      </mj-column>
    </mj-section>
	<mj-section>
      <mj-column>
        <mj-image width="300px" src="https://storage-thumbnails.bananatag.com/images/zd8JyS/831be28551fbbe786a569f3d1b7ee525.png" />
      </mj-column>
    </mj-section>
  </mj-body>
</mjml>
`,
		DefaultData: map[string]string{},
	}

	// Create the HTML content from the JSON data
	htmlContent, err := mjmlClient.CreateHTMLContent("", jsonData)
	if err != nil {
		log.Fatalf("Failed to generate HTML content: %v", err)
	}

	fmt.Println(htmlContent)

	// Initialize the MailClient
	mailClient := testreport.MailClient{Domain: "corpadds.com"}

	// Send the email
	sender := "autotest@yourdomain.com"
	recipient := "k.zhang@ceridian.com"
	subject := "Cypress Test Report"
	attachmentPath := "" // Add an attachment if needed

	err = mailClient.SendHTMLEmailWithAttachment(sender, recipient, subject, htmlContent, attachmentPath)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	fmt.Println("Email sent successfully!")
}
