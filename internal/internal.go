// internal/internal.go
package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/715047274/jenkinsCmd/model"
	"github.com/Boostport/mjml-go"
	"io/ioutil"
	"os"
)

func readJSON(filePath string) (model.TestReport, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return model.TestReport{}, err
	}

	var report model.TestReport
	err = json.Unmarshal(data, &report)
	if err != nil {
		return model.TestReport{}, err
	}

	return report, nil
}

func GenerateHTML(inputFile, outputDir string) {
	// Read Cypress test results from JSON file
	//report, err := readJSON(inputFile)
	// Implement JSON parsing logic (using gojsonschema) to extract relevant information

	// Implement MJML template generation logic
	mjmlTemplate := `<mjml><mj-body><mj-section><mj-column><mj-divider border-color="#F45E43"></mj-divider><mj-text font-size="20px" color="#F45E43" font-family="helvetica">Hello World</mj-text></mj-column></mj-section></mj-body></mjml>`

	// Convert MJML to HTML
	htmlOutput, err := mjml.ToHTML(context.Background(), mjmlTemplate, mjml.WithMinify(false))
	if err != nil {
		fmt.Println("Error converting MJML to HTML:", err)
		return
	}

	// Save HTML output to a file in the specified output directory
	outputFile := outputDir + "/output.html"
	err = os.WriteFile(outputFile, []byte(htmlOutput), 0644)
	if err != nil {
		fmt.Println("Error writing HTML output:", err)
		return
	}

	fmt.Println("HTML email template generated successfully at:", outputFile)
}
