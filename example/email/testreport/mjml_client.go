package testreport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Boostport/mjml-go"
	"strings"
)

type MJMLClient struct {
	DefaultTemplate string
	DefaultData     map[string]string
}

func (mj *MJMLClient) CreateHTMLContent(templateInput, jsonInput string) (string, error) {
	template := mj.DefaultTemplate
	if templateInput != "" {
		template = templateInput
	}

	var data map[string]string
	if jsonInput != "" {
		if err := json.Unmarshal([]byte(jsonInput), &data); err != nil {
			return "", fmt.Errorf("failed to parse JSON input: %w", err)
		}
	} else {
		data = mj.DefaultData
	}

	for key, value := range data {
		template = strings.ReplaceAll(template, fmt.Sprintf("{{.%s}}", key), value)
	}

	output, err := mjml.ToHTML(context.Background(), template, mjml.WithMinify(true))
	if err != nil {
		var mjmlError mjml.Error
		if errors.As(err, &mjmlError) {
			return "", fmt.Errorf("MJML Conversion Error: %s", mjmlError.Message)
		}
		return "", err
	}

	return output, nil
}

//// RenderFailuresAsHTML converts Failures to an HTML table
//func RenderFailuresAsHTML(failures []map[string]string) string {
//	var failureDetails strings.Builder
//	if len(failures) > 0 {
//		failureDetails.WriteString("<table><tr><th>Suite</th><th>Test</th><th>Error</th></tr>")
//		for _, failure := range failures {
//			failureDetails.WriteString(fmt.Sprintf(
//				"<tr><td>%s</td><td>%s</td><td>%s</td></tr>",
//				failure["Suite"], failure["Test"], failure["Error"],
//			))
//		}
//		failureDetails.WriteString("</table>")
//	} else {
//		failureDetails.WriteString("<p>No failures found.</p>")
//	}
//	return failureDetails.String()
//}

func (mj *MJMLClient) PrepareHTMLContent(templateInput, jsonData string) (string, error) {
	template := mj.DefaultTemplate
	if templateInput != "" {
		template = templateInput
	}
	var parsedData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &parsedData); err != nil {
		return "", fmt.Errorf("failed to parse JSON data: %w", err)
	}

	fmt.Println(parsedData["Stats"])
	// Render Failures as HTML if it's an array
	if failures, ok := parsedData["Failures"].([]map[string]string); ok {
		fmt.Println(failures)
		parsedData["Failures"] = RenderFailuresAsHTML(failures)
	}

	if data, ok := parsedData["Stats"].(map[string]string); ok {

		for key, value := range data {
			fmt.Println(value)
			template = strings.ReplaceAll(template, fmt.Sprintf("{{.%s}}", key), value)
		}
	}

	// Use the parsedData for template rendering
	// Here, integrate with your MJML rendering logic
	//htmlContent, err := mjml.ToHTML(context.Background(), mjmlTemplate, mjml.WithMinify(parsedData))
	htmlContent, err := mjml.ToHTML(context.Background(), template, mjml.WithMinify(true))
	if err != nil {
		var mjmlError mjml.Error
		if errors.As(err, &mjmlError) {
			return "", fmt.Errorf("MJML Conversion Error: %s", mjmlError.Message)
		}
		return "", err
	}

	return htmlContent, nil
}

func RenderFailuresAsHTML(failures []map[string]string) string {
	var htmlBuilder strings.Builder
	htmlBuilder.WriteString("<table><tr><th>Suite</th><th>Test</th><th>Error</th></tr>")
	for _, failure := range failures {
		htmlBuilder.WriteString(fmt.Sprintf(
			"<tr><td>%s</td><td>%s</td><td>%s</td></tr>",
			failure["Suite"], failure["Test"], failure["Error"],
		))
	}
	htmlBuilder.WriteString("</table>")
	return htmlBuilder.String()
}

func (mj *MJMLClient) PrepareHTMLContent2(templateInput, jsonData string) (string, error) {
	template := mj.DefaultTemplate
	if templateInput != "" {
		template = templateInput
	}

	// Parse JSON data
	var parsedData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &parsedData); err != nil {
		return "", fmt.Errorf("failed to parse JSON data: %w", err)
	}

	// Convert Failures to an HTML table
	if failures, ok := parsedData["Failures"].([]interface{}); ok {
		var failureData []map[string]string
		for _, item := range failures {
			if failure, ok := item.(map[string]interface{}); ok {
				stringMap := make(map[string]string)
				for key, value := range failure {
					stringMap[key] = fmt.Sprintf("%v", value)
				}
				failureData = append(failureData, stringMap)
			}
		}
		parsedData["Failures"] = RenderFailuresAsHTML(failureData)
	}
	//if data, ok := parsedData["Stats"].(map[string]string); ok {
	//	for key, value := range data {
	//		template = strings.ReplaceAll(template, fmt.Sprintf("{{.%s}}", key), value)
	//	}
	//}
	// Manually replace placeholders in the template
	//replacedTemplate := templateInput
	for key, value := range parsedData {
		switch v := value.(type) {
		case string:
			template = strings.ReplaceAll(template, fmt.Sprintf("{{.%s}}", key), v)
		case map[string]interface{}:
			for subKey, subValue := range v {
				template = strings.ReplaceAll(
					template,
					fmt.Sprintf("{{.%s.%s}}", key, subKey),
					fmt.Sprintf("%v", subValue),
				)
			}
		}
	}
	output, err := mjml.ToHTML(context.Background(), template, mjml.WithMinify(true))
	if err != nil {
		var mjmlError mjml.Error
		if errors.As(err, &mjmlError) {
			return "", fmt.Errorf("MJML Conversion Error: %s", mjmlError.Message)
		}
		return "", err
	}

	return output, nil
	// Render MJML to HTML
}
