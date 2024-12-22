package testreport

import (
	"fmt"
	"strings"
	"testing"
)

func TestCreateHTMLContent(t *testing.T) {
	mjmlClient := MJMLClient{
		DefaultTemplate: `
<mjml>
  <mj-body>
    <mj-section>
      <mj-column>
        <mj-text font-size="20px">Hello, {{.Name}}!</mj-text>
      </mj-column>
    </mj-section>
  </mj-body>
</mjml>`,
	}

	jsonData := `{"Name": "Test User"}`

	htmlContent, err := mjmlClient.CreateHTMLContent("", jsonData)
	if err != nil {
		t.Fatalf("CreateHTMLContent failed: %v", err)
	}

	expectedSubstring := "<p>Hello, Test User!</p>"
	if !strings.Contains(htmlContent, expectedSubstring) {
		t.Errorf("Expected HTML content to contain %q, but it did not. Got: %s", expectedSubstring, htmlContent)
	}
}

func TestCreateHTMLContentWithInvalidJSON(t *testing.T) {
	mjmlClient := MJMLClient{
		DefaultTemplate: `
<mjml>
  <mj-body>
    <mj-section>
      <mj-column>
        <mj-text font-size="20px">Hello, {{.Name}}!</mj-text>
      </mj-column>
    </mj-section>
  </mj-body>
</mjml>`,
	}

	invalidJsonData := `{"Name": "Test User"`

	_, err := mjmlClient.CreateHTMLContent("", invalidJsonData)
	if err == nil {
		t.Fatal("Expected CreateHTMLContent to fail due to invalid JSON, but it succeeded")
	}
}

func TestCreateHTMLContentWithTemplateError(t *testing.T) {
	mjmlClient := MJMLClient{
		DefaultTemplate: `
<mjml>
  <mj-body>
    <mj-section>
      <mj-column>
        <mj-text font-size="20px">Hello, {{.Name</mj-text>
      </mj-column>
    </mj-section>
  </mj-body>
</mjml>`,
	}

	jsonData := `{"Name": "Test User"}`

	_, err := mjmlClient.CreateHTMLContent("", jsonData)
	if err == nil {
		t.Fatal("Expected CreateHTMLContent to fail due to template parsing error, but it succeeded")
	}
}

func TestCreateHTMLContentWithComplex(t *testing.T) {
	mjmlClient := MJMLClient{
		DefaultTemplate: `
<mjml>
  <mj-body>
    <mj-section>
      <mj-column>
        <mj-text font-size="24px" font-family="helvetica" color="#333333" align="center">Cypress Test Report</mj-text>
        <mj-divider border-color="#F45E43"></mj-divider>
        <mj-text font-size="16px" font-family="helvetica" color="#555555">Test Statistics:</mj-text>
        <mj-table>
          <tr style="background-color:#f0f0f0;text-align:left;">
            <th>Metric</th>
            <th>Value</th>
          </tr>
          <tr>
            <td>Suites</td>
            <td>{{ .Stats.suites }}</td>
          </tr>
          <tr>
            <td>Tests</td>
            <td>{{ .Stats.tests }}</td>
          </tr>
          <tr>
            <td>Passes</td>
            <td>{{ .Stats.passes }}</td>
          </tr>
          <tr>
            <td>Failures</td>
            <td>{{ .Stats.failures }}</td>
          </tr>
          <tr>
            <td>Pending</td>
            <td>{{ .Stats.pending }}</td>
          </tr>
          <tr>
            <td>Skipped</td>
            <td>{{ .Stats.skipped }}</td>
          </tr>
        </mj-table>
        <mj-text font-size="16px" font-family="helvetica" color="#555555">Failure Details:</mj-text>
        {{ range .Failures }}
        <mj-text font-size="14px" font-family="helvetica" color="#555555">
          <strong>Suite:</strong> {{ .Suite }}<br/>
          <strong>Test:</strong> {{ .Test }}<br/>
          <strong>Error:</strong> {{ .Error }}<br/>
        </mj-text>
        {{ end }}
      </mj-column>
    </mj-section>
  </mj-body>
</mjml>`,
	}

	jsonData := `{
		"Stats": {
			"suites": 23,
			"tests": 231,
			"passes": 211,
			"failures": 13,
			"pending": 3,
			"skipped": 4
		},
		"Failures": [
			{
				"Suite": "Dashboard Responsiveness Wide to Narrow",
				"Test": "Dashboard Responsiveness Wide to Narrow GPTCI-16459 - Dashboard_Responsiveness_Breakpoint1200_WideToNarrow",
				"Error": "title: Failed screenshot\nvalue: screenshots\\Dashboard\\Responsiveness.spec.js/Dashboard%20Responsiveness%20Wide%20to%20Narrow%20--%20GPTCI-16459%20-%20Dashboard_Responsiveness_Breakpoint1200_WideToNarrow%20(failed)%20(attempt%202).png\n"
			}
		]
	}`

	expectedSubstring := `<td>Pending</td><td>3</td>`

	htmlContent, err := mjmlClient.CreateHTMLContent("", jsonData)
	fmt.Println(htmlContent)
	if err != nil {
		t.Fatalf("CreateHTMLContent failed: %v", err)
	}

	if !strings.Contains(htmlContent, expectedSubstring) {
		t.Errorf("Expected HTML content to::contentReference[oaicite:0]{index=0}")

	}
}
