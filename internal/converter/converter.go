package converter

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/sottey/dashuni/internal/model"
)

// ApplyTemplate takes a Site object and a template file path,
// renders the template with the Site data, and returns the resulting string.
func ApplyTemplate(site *model.Site, mappingPath string) (string, error) {
	// Load the mapping template from file
	tmplBytes, err := os.ReadFile(mappingPath)
	if err != nil {
		return "", fmt.Errorf("failed to read mapping template: %w", err)
	}

	// Parse the template
	tmpl, err := template.New("mapping").Funcs(templateFuncs()).Parse(string(tmplBytes))
	if err != nil {
		return "", fmt.Errorf("failed to parse mapping template: %w", err)
	}

	// Execute the template with Site data
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, site)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// templateFuncs returns custom template functions if needed.
func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"upper": strings.ToUpper,
		"lower": strings.ToLower,
		"trim":  strings.TrimSpace,
		// Add more helpers here as needed
	}
}
