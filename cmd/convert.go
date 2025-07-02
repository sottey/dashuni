/*
Copyright © 2025 sottey

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/sottey/dashuni/internal/converter"
	"github.com/sottey/dashuni/internal/parser"
)

var (
	inputFile   string
	mappingFile string
	outputFile  string
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert your universal site schema using a mapping template",
	Long: `Convert takes your universal homelab site JSON 
and renders it to another dashboard's config format using a Go text/template mapping.`,
	Run: func(cmd *cobra.Command, args []string) {
		if inputFile == "" || mappingFile == "" || outputFile == "" {
			fmt.Println("❌ Error: --input, --mapping, and --output are required")
			cmd.Usage()
			os.Exit(1)
		}

		// Load the universal site schema
		root, err := parser.LoadRootFromFile(inputFile)
		if err != nil {
			fmt.Printf("❌ Failed to load input: %v\n", err)
			os.Exit(1)
		}

		// Apply the mapping template
		result, err := converter.ApplyTemplate(&root.Site, mappingFile)
		if err != nil {
			fmt.Printf("❌ Conversion failed: %v\n", err)
			os.Exit(1)
		}

		// Write the output
		err = os.WriteFile(outputFile, []byte(result), 0644)
		if err != nil {
			fmt.Printf("❌ Failed to write output: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ Conversion successful! Output written to %s\n", outputFile)
	},
}

func init() {
	convertCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Path to universal site JSON input")
	convertCmd.Flags().StringVarP(&mappingFile, "mapping", "m", "", "Path to mapping template file")
	convertCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Path to output file")
	rootCmd.AddCommand(convertCmd)
}
