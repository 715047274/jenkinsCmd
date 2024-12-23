package testreport

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/Boostport/mjml-go"
)

type MjmlHandler struct {
	DefaultTemplate string
	DefaultData     map[string]string
}

// CreateHTMLContent generates HTML content from the provided MJML template and JSON data.
func (mj *MjmlHandler) CreateHTMLContent(templateInput string, jsonData string) (string, error) {
	var tmpl string
	// Determine if templateInput is a filepath or a direct string
	if strings.HasSuffix(templateInput, ".tmpl") {
		// Read the template file
		content, err := os.ReadFile(templateInput)
		if err != nil {
			return "", fmt.Errorf("failed to read template file: %w", err)
		}
		tmpl = string(content)
	} else {
		// Use the input string as the template
		tmpl = templateInput
	}
	// Use the provided template or fall back to the default
	// If no input is provided, fall back to the default template
	if tmpl == "" {
		tmpl = mj.DefaultTemplate
	}

	// Parse the JSON data into a map
	var rawData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &rawData); err != nil {
		return "", fmt.Errorf("failed to parse JSON data: %w", err)
	}

	// Convert float64 to int for numeric values
	parsedData := convertFloatsToInts(rawData)

	// Define custom functions
	funcMap := template.FuncMap{
		"mod": func(a, b int) int {
			return a % b
		},
		"add": func(a, b int) int {
			return a + b
		},
	}

	// Parse the MJML template using html/template with custom functions
	t, err := template.New("mjml").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute the template with the parsed data
	var processedTemplate bytes.Buffer
	if err := t.Execute(&processedTemplate, parsedData); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}
	// Log the generated MJML template
	//fmt.Println("Generated MJML Template:")
	//fmt.Println(processedTemplate.String())

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

func convertFloatsToInts(data map[string]interface{}) map[string]interface{} {
	for key, value := range data {
		switch v := value.(type) {
		case float64:
			data[key] = int(v)
		case []interface{}:
			for i, item := range v {
				if num, ok := item.(float64); ok {
					v[i] = int(num)
				} else if subMap, ok := item.(map[string]interface{}); ok {
					v[i] = convertFloatsToInts(subMap)
				}
			}
		case map[string]interface{}:
			data[key] = convertFloatsToInts(v)
		}
	}
	return data
}
