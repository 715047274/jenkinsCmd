package testreport

import (
	"fmt"
	"strings"
	"testing"
)

func TestCreateHTMLContent(t *testing.T) {
	MjmlHandler := MjmlHandler{
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

	htmlContent, err := MjmlHandler.CreateHTMLContent("", jsonData)
	if err != nil {
		t.Fatalf("CreateHTMLContent failed: %v", err)
	}

	expectedSubstring := "<p>Hello, Test User!</p>"
	if !strings.Contains(htmlContent, expectedSubstring) {
		t.Errorf("Expected HTML content to contain %q, but it did not. Got: %s", expectedSubstring, htmlContent)
	}
}

func TestCreateHTMLContentWithInvalidJSON(t *testing.T) {
	MjmlHandler := MjmlHandler{
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

	_, err := MjmlHandler.CreateHTMLContent("", invalidJsonData)
	if err == nil {
		t.Fatal("Expected CreateHTMLContent to fail due to invalid JSON, but it succeeded")
	}
}

func TestCreateHTMLContentWithTemplateError(t *testing.T) {
	MjmlHandler := MjmlHandler{
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

	_, err := MjmlHandler.CreateHTMLContent("", jsonData)
	if err == nil {
		t.Fatal("Expected CreateHTMLContent to fail due to template parsing error, but it succeeded")
	}
}

func TestCreateHTMLContentWithComplex(t *testing.T) {
	MjmlHandler := MjmlHandler{
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

	htmlContent, err := MjmlHandler.CreateHTMLContent("", jsonData)
	fmt.Println(htmlContent)
	if err != nil {
		t.Fatalf("CreateHTMLContent failed: %v", err)
	}

	if !strings.Contains(htmlContent, expectedSubstring) {
		t.Errorf("Expected HTML content to::contentReference[oaicite:0]{index=0}")

	}
}

func TestMjmlHandlerWithModAndAdd(t *testing.T) {
	// Test template with mod and add functions
	templateInput := `
	<mjml>
	  <mj-body>
		<mj-section>
		  <mj-column>
			<mj-table>
			  <tr style="background-color:#f0f0f0;text-align:left;">
				<th>Index</th>
				<th>Value</th>
				<th>Mod 2</th>
				<th>Add 10</th>
			  </tr>
			  {{ range $index, $value := .Values }}
			  <tr>
				<td>{{ $index }}</td>
				<td>{{ $value }}</td>
				<td>{{ mod $index 2 }}</td>
				<td>{{ add $value 10 }}</td>
			  </tr>
			  {{ end }}
			</mj-table>
		  </mj-column>
		</mj-section>
	  </mj-body>
	</mjml>
	`

	// Sample JSON data for the test
	jsonData := `{
		"Values": [1, 2, 3, 4, 5]
	}`

	// Create an MjmlHandler instance
	MjmlHandler := MjmlHandler{}

	// Generate the HTML content
	output, err := MjmlHandler.CreateHTMLContent(templateInput, jsonData)
	if err != nil {
		t.Fatalf("Failed to generate HTML content: %v", err)
	}

	// Verify that the output contains expected values
	expectedStrings := []string{
		"<td>0</td><td>1</td><td>0</td><td>11</td>", // First row
		"<td>1</td><td>2</td><td>1</td><td>12</td>", // Second row
		"<td>2</td><td>3</td><td>0</td><td>13</td>", // Third row
		"<td>3</td><td>4</td><td>1</td><td>14</td>", // Fourth row
		"<td>4</td><td>5</td><td>0</td><td>15</td>", // Fifth row
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(output, expected) {
			t.Errorf("Output missing expected value: %s\nOutput:\n%s", expected, output)
		}
	}
}
func TestMjmlHandlerWithGridView(t *testing.T) {
	// Test template to create a grid view with rows of 4 columns
	templateInput := `
	<mjml>
	  <mj-body>
		<mj-section>
		  <mj-group>
			{{ range $i, $num := .Numbers }}
			  {{ if eq (mod $i 4) 0 }}
			  {{ if ne $i 0 }}
			  </mj-group>
			  <mj-group>
			  {{ end }}
			  {{ end }}
			  <mj-column>
				<mj-text font-size="20px" line-height="24px" color="#333333" align="center">
				  {{ $num }}
				</mj-text>
			  </mj-column>
			{{ end }}
		  </mj-group>
		</mj-section>
	  </mj-body>
	</mjml>
	`

	// Sample JSON data for the test
	jsonData := `{
		"Numbers": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
	}`

	// Create an MjmlHandler instance
	MjmlHandler := MjmlHandler{}

	// Generate the HTML content
	output, err := MjmlHandler.CreateHTMLContent(templateInput, jsonData)
	if err != nil {
		t.Fatalf("Failed to generate HTML content: %v", err)
	}

	// Verify that the output contains expected grid structure
	expectedStrings := []string{
		"<mj-column>",          // Column start
		"<mj-text>1</mj-text>", // First number
		"<mj-text>5</mj-text>", // Fifth number
		"</mj-column>",         // Column end
		"</mj-group>",          // Group end
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(output, expected) {
			t.Errorf("Output missing expected value: %s\nOutput:\n%s", expected, output)
		}
	}
}
