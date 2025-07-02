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
	"path/filepath"
	"strings"

	"github.com/sottey/dashuni/internal"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available mapping templates",
	Run: func(cmd *cobra.Command, args []string) {
		mappingDir := "mappings"

		entries, err := os.ReadDir(mappingDir)
		if err != nil {
			fmt.Printf("Error reading mappings directory '%s': %v\n", mappingDir, err)
			return
		}

		fmt.Println("Available templates:")
		count := 0
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			if filepath.Ext(entry.Name()) == ".tmpl" {
				dashName := strings.Replace(entry.Name(), ".tmpl", "", 1)
				dashName = internal.ToProperCase(dashName)
				fmt.Printf("- %s (%s)\n", dashName, "./mappings/"+entry.Name())
				count++
			}
		}

		if count == 0 {
			fmt.Println("No .tmpl files found in mappings directory.")
		}
	},
}
