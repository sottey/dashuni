package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sottey/dashuni/internal/model"
)

// LoadRootFromFile loads the universal site config from a JSON file.
func LoadRootFromFile(path string) (*model.Root, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", path, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var root model.Root
	if err := decoder.Decode(&root); err != nil {
		return nil, fmt.Errorf("failed to parse JSON in %s: %w", path, err)
	}

	return &root, nil
}

// SaveRootToFile writes the site config to a JSON file.
// If pretty is true, it will indent the output.
func SaveRootToFile(root *model.Root, path string, pretty bool) error {
	var data []byte
	var err error

	if pretty {
		data, err = json.MarshalIndent(root, "", "  ")
	} else {
		data, err = json.Marshal(root)
	}

	if err != nil {
		return fmt.Errorf("failed to serialize JSON: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", path, err)
	}

	return nil
}
