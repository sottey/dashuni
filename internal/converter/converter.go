package converter

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/google/uuid"
	"github.com/sottey/dashuni/internal/model"
)

// ApplyTemplate takes a Site object and a template file path,
// renders the template with the Site data (and optional FA map), and returns the resulting string.
func ApplyTemplate(site *model.Site, mappingPath string) (string, error) {
	// Check if template declares FA requirement
	needsFA, err := TemplateRequiresFA(mappingPath)
	if err != nil {
		log.Fatalf("Error checking template header: %v", err)
	}

	var faMap map[string]string
	if needsFA {
		fmt.Println("🔎 Template declares it requires FA mapping. Loading fa-mapping.json...")
		faMap, err = LoadFaMap("./mappings/fa-mapping.json")
		if err != nil {
			log.Fatalf("❌ Failed to load fa-mapping.json: %v", err)
		}
	}

	// Load template file
	tmplBytes, err := os.ReadFile(mappingPath)
	if err != nil {
		return "", fmt.Errorf("failed to read mapping template: %w", err)
	}

	// Parse template with helpers
	tmpl, err := template.New("mapping").Funcs(templateFuncs()).Parse(string(tmplBytes))
	if err != nil {
		return "", fmt.Errorf("failed to parse mapping template: %w", err)
	}

	// Compose data with Site and FAMap
	data := struct {
		Site  *model.Site
		FAMap map[string]string
	}{
		Site:  site,
		FAMap: faMap,
	}

	// Execute template
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"upper": strings.ToUpper,
		"lower": strings.ToLower,
		"trim":  strings.TrimSpace,
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, fmt.Errorf("invalid dict call: odd number of args")
			}
			d := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, fmt.Errorf("dict keys must be strings")
				}
				d[key] = values[i+1]
			}
			return d, nil
		},
		"add": func(a, b int) int {
			return a + b
		},
		"uuid": func() string {
			return uuid.NewString()
		},
	}
}

func TemplateRequiresFA(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "requiresFA: true") {
			return true, nil
		}
	}
	return false, scanner.Err()
}

func LoadFaMap(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result map[string]string
	err = json.Unmarshal(data, &result)
	return result, err
}
