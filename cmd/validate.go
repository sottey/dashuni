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

	"github.com/sottey/dashuni/internal/parser"
)

var validateInput string

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate your universal site JSON schema",
	Long:  `Validate will check that your universal homelab schema JSON is well-formed.`,
	Run: func(cmd *cobra.Command, args []string) {
		if validateInput == "" {
			fmt.Println("❌ Error: --input is required")
			cmd.Usage()
			os.Exit(1)
		}

		root, err := parser.LoadRootFromFile(validateInput)
		if err != nil {
			fmt.Printf("❌ Failed to load or parse JSON: %v\n", err)
			os.Exit(1)
		}

		// Basic built-in validation (could be expanded)
		if root.Site.Name == "" {
			fmt.Println("❌ Validation error: site.name is required")
			os.Exit(1)
		}
		if len(root.Site.Pages) == 0 {
			fmt.Println("❌ Validation error: site.pages must not be empty")
			os.Exit(1)
		}

		fmt.Println("✅ Validation successful!")
	},
}

func init() {
	validateCmd.Flags().StringVarP(&validateInput, "input", "i", "", "Path to universal site JSON input")
	rootCmd.AddCommand(validateCmd)
}
