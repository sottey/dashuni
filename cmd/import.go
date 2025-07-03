package cmd

import (
	"encoding/json"
	"fmt"

	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/sottey/dashuni/internal/importer"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a dashboard config into canonical dashuni.json",
	Run: func(cmd *cobra.Command, args []string) {
		inputPath, _ := cmd.Flags().GetString("input")
		mappingPath, _ := cmd.Flags().GetString("mapping")
		outputPath, _ := cmd.Flags().GetString("output")

		inputBytes, err := os.ReadFile(inputPath)
		if err != nil {
			fmt.Printf("❌ Failed to read input: %v\n", err)
			os.Exit(1)
		}

		mappingBytes, err := os.ReadFile(mappingPath)
		if err != nil {
			fmt.Printf("❌ Failed to read mapping: %v\n", err)
			os.Exit(1)
		}

		// Load input (YAML)
		var inputData map[string]interface{}
		if err := yaml.Unmarshal(inputBytes, &inputData); err != nil {
			fmt.Printf("❌ Failed to parse input YAML: %v\n", err)
			os.Exit(1)
		}

		// Load mapping (JSON)
		var mapping map[string]string
		if err := json.Unmarshal(mappingBytes, &mapping); err != nil {
			fmt.Printf("❌ Failed to parse mapping JSON: %v\n", err)
			os.Exit(1)
		}

		// Import
		canonical, err := importer.ImportWithMapping(inputData, mapping)
		if err != nil {
			fmt.Printf("❌ Import failed: %v\n", err)
			os.Exit(1)
		}

		// Write output
		outBytes, _ := json.MarshalIndent(canonical, "", "  ")
		if err := os.WriteFile(outputPath, outBytes, 0644); err != nil {
			fmt.Printf("❌ Failed to write output: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ Imported canonical dashuni.json written to %s\n", outputPath)
	},
}

func init() {
	importCmd.Flags().String("input", "", "Path to the dashboard config file (YAML/JSON)")
	importCmd.Flags().String("mapping", "", "Path to the mapping file (JSON)")
	importCmd.Flags().String("output", "", "Path to write the canonical dashuni.json")
	rootCmd.AddCommand(importCmd)
}
