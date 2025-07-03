package importer

import (
	"fmt"
	"strings"
)

// ImportWithMapping is the main entry point that converts input data
// into canonical dashuni JSON using a field mapping.
func ImportWithMapping(inputData map[string]interface{}, mapping map[string]string) (map[string]interface{}, error) {
	canonical := make(map[string]interface{})

	for canonicalPath, sourcePath := range mapping {
		value, ok := getValueFromPath(inputData, sourcePath)
		if !ok {
			fmt.Printf("⚠️  Warning: Source path %q not found.\n", sourcePath)
			continue
		}
		setValueAtPath(canonical, canonicalPath, value)
	}

	return canonical, nil
}

// getValueFromPath navigates through the input data using a dot-separated path.
// It handles both map fields and arrays (with [] suffix in the path).
func getValueFromPath(data interface{}, path string) (interface{}, bool) {
	parts := strings.Split(path, ".")
	current := data

	for _, part := range parts {
		if strings.HasSuffix(part, "[]") {
			part = strings.TrimSuffix(part, "[]")
			m, ok := current.(map[string]interface{})
			if !ok {
				return nil, false
			}
			next, ok := m[part]
			if !ok {
				return nil, false
			}
			arr, ok := next.([]interface{})
			if !ok {
				return nil, false
			}

			// Process each item recursively if further path remains
			remaining := strings.Join(parts[1:], ".")
			if remaining == "" {
				return arr, true
			}

			var results []interface{}
			for _, item := range arr {
				v, ok := getValueFromPath(item, strings.Join(parts[1:], "."))
				if ok {
					results = append(results, v)
				}
			}
			return results, true
		}

		m, ok := current.(map[string]interface{})
		if !ok {
			return nil, false
		}
		next, ok := m[part]
		if !ok {
			return nil, false
		}
		current = next
	}
	return current, true
}

// setValueAtPath sets a value in the canonical map at the given dot-separated path.
func setValueAtPath(data map[string]interface{}, path string, value interface{}) {
	parts := strings.Split(path, ".")
	last := len(parts) - 1
	current := data

	for i, part := range parts {
		if i == last {
			current[part] = value
			return
		}

		if next, ok := current[part]; ok {
			if m, ok := next.(map[string]interface{}); ok {
				current = m
			} else {
				newMap := make(map[string]interface{})
				current[part] = newMap
				current = newMap
			}
		} else {
			newMap := make(map[string]interface{})
			current[part] = newMap
			current = newMap
		}
	}
}
