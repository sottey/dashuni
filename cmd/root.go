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
)

var rootCmd = &cobra.Command{
    Use:   "dashuni",
    Short: "dashuni - Universal homelab dashboard config converter",
    Long: `dashuni is a CLI tool for managing a universal homelab config schema
and converting it to the format of many popular self-hosted dashboards.

Examples:
  dashuni validate --input site.json
  dashuni convert --input site.json --mapping dashy.map.json --output dashy-config.yaml
  dashuni list
`,
    Version: "0.1.0",
    Run: func(cmd *cobra.Command, args []string) {
        // Default behavior if no subcommand
        fmt.Println("dashuni - A CLI for universal homelab dashboard config conversion.")
        fmt.Println("Use --help to see available commands.")
    },
}

// Execute runs the root command.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize()
    // Add your subcommands here
    rootCmd.AddCommand(validateCmd)
    rootCmd.AddCommand(convertCmd)
    rootCmd.AddCommand(listCmd)
}