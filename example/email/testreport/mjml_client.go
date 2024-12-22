package testreport

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"strings"

	"github.com/Boostport/mjml-go"
)

type MJMLClient struct {
	DefaultTemplate string
	DefaultData     map[string]string
}

// CreateHTMLContent generates HTML content from the provided MJML template and JSON data.
func (mj *MJMLClient) CreateHTMLContent(templateInput string, jsonData string) (string, error) {
	//// Use the provided template or fall back to the default
	//template := mj.DefaultTemplate
	//if templateInput != "" {
	//	template = templateInput
	//}
	//
	//// Parse the JSON data into a map
	//var parsedData map[string]interface{}
	//if err := json.Unmarshal([]byte(jsonData), &parsedData); err != nil {
	//	return "", fmt.Errorf("failed to parse JSON data: %w", err)
	//}
	//// Replace placeholders in the template with actual data
	//processedTemplate := replaceTemplatePlaceholders(template, parsedData)
	////fmt.Println("-----------------------------------------------")
	////
	////fmt.Println(processedTemplate)
	//// Convert the processed MJML template to HTML
	//output, err := mjml.ToHTML(context.Background(), processedTemplate, mjml.WithMinify(true))
	//
	//if err != nil {
	//	var mjmlError mjml.Error
	//	if errors.As(err, &mjmlError) {
	//		return "", fmt.Errorf("MJML Conversion Error: %s", mjmlError.Message)
	//	}
	//	return "", err
	//}
	//return output, nil

	// Use the provided template or fall back to the default
	tmpl := mj.DefaultTemplate
	if templateInput != "" {
		tmpl = templateInput
	}

	// Parse the JSON data into a map
	var parsedData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &parsedData); err != nil {
		return "", fmt.Errorf("failed to parse JSON data: %w", err)
	}

	// Parse the MJML template using html/template
	t, err := template.New("mjml").Parse(tmpl)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute the template with the parsed data
	var processedTemplate bytes.Buffer
	if err := t.Execute(&processedTemplate, parsedData); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	// Convert the processed MJML template to HTML
	output, err := mjml.ToHTML(context.Background(), processedTemplate.String(), mjml.WithMinify(true))
	if err != nil {
		var mjmlError mjml.Error
		if errors.As(err, &mjmlError) {
			return "", fmt.Errorf("MJML Conversion Error: %s", mjmlError.Message)
		}
		return "", err
	}

	return output, nil

}

// replaceTemplatePlaceholders replaces placeholders in the MJML template with actual data.
func replaceTemplatePlaceholders(template string, data map[string]interface{}) string {
	for key, value := range data {
		switch v := value.(type) {
		case string:
			template = strings.ReplaceAll(template, fmt.Sprintf("{{.%s}}", key), v)
		case float64, int:
			template = strings.ReplaceAll(template, fmt.Sprintf("{{.%s}}", key), fmt.Sprintf("%v", v))
		case map[string]interface{}:
			for subKey, subValue := range v {
				template = strings.ReplaceAll(template, fmt.Sprintf("{{.%s.%s}}", key, subKey), fmt.Sprintf("%v", subValue))
			}
		case []interface{}:
			if key == "Failures" {
				template = strings.ReplaceAll(template, fmt.Sprintf("{{.%s}}", key), renderFailuresAsHTML(v))
			}
		}
	}
	return template
}

// renderFailuresAsHTML converts the Failures data into an HTML table.
func renderFailuresAsHTML(failures []interface{}) string {
	var htmlBuilder strings.Builder
	htmlBuilder.WriteString("<table><tr><th>Suite</th><th>Test</th><th>Error</th></tr>")
	for _, failure := range failures {
		if failureMap, ok := failure.(map[string]interface{}); ok {
			htmlBuilder.WriteString(fmt.Sprintf(
				"<tr><td>%s</td><td>%s</td><td>%s</td></tr>",
				failureMap["Suite"], failureMap["Test"], failureMap["Error"],
			))
		}
	}
	htmlBuilder.WriteString("</table>")
	return htmlBuilder.String()
}
