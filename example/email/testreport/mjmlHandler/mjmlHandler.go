package mjmlHandler

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
	CustomFuncMap   template.FuncMap
}

// RegisterFunc allows registering custom functions to the handler's function map.
// If the key already exists, the new function will override the existing one.
func (mj *MjmlHandler) RegisterFunc(name string, function interface{}) {
	if mj.CustomFuncMap == nil {
		mj.CustomFuncMap = make(template.FuncMap)
	}

	if _, exists := mj.CustomFuncMap[name]; exists {
		fmt.Printf("Function with name '%s' already exists. Overriding with the new function.\n", name)
	}

	mj.CustomFuncMap[name] = function
}

// CreateHTMLContent generates HTML content from the provided MJML template and JSON data.
func (mj *MjmlHandler) CreateHTMLContent(templateInput string, jsonData string) (string, error) {
	var tmpl string
	// Determine if templateInput is a filepath or a direct string
	if strings.HasSuffix(templateInput, ".tmpl") {
		content, err := os.ReadFile(templateInput)
		if err != nil {
			return "", fmt.Errorf("failed to read template file: %w", err)
		}
		tmpl = string(content)
	} else {
		tmpl = templateInput
	}

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

	// Default template functions
	defaultFuncMap := template.FuncMap{
		"mod": func(a, b int) int {
			return a % b
		},
		"add": func(a, b int) int {
			return a + b
		},
	}

	// Merge default functions with custom functions
	mergedFuncMap := make(template.FuncMap)
	for k, v := range defaultFuncMap {
		mergedFuncMap[k] = v
	}
	for k, v := range mj.CustomFuncMap {
		mergedFuncMap[k] = v
	}

	// Parse the MJML template using html/template with merged functions
	t, err := template.New("mjml").Funcs(mergedFuncMap).Parse(tmpl)
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
