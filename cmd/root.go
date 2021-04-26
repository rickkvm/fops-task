/*
Copyright © 2021 Rick Chen <rick.kvm@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// files for subcommands, will be initialized in linecount, checksum
var targetFilename string

// update during compile time
var VERSION string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fops",
	Short: "File Ops",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	SilenceErrors: true,
	SilenceUsage:  true,
	//Long:  `File Operaters with linecount and checksum`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	VERSION = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		fmt.Println()
		os.Exit(1)
	}
}

func init() {
	cobra.EnableCommandSorting = false
	rootCmd.SetHelpTemplate(helpTemplate)
	rootCmd.SetUsageTemplate(usageTemplate)
	rootCmd.AddCommand(linecountCmd)
	rootCmd.AddCommand(checksumCmd)
	rootCmd.SetHelpCommand(helpCmd)
}

func checkFileExist(path string) error {
	stat, e := os.Lstat(targetFilename)
	if e != nil {
		if os.IsNotExist(e) {
			return fmt.Errorf("No such file '%s'", targetFilename)
		}
		return e
	}

	if stat.IsDir() {
		return fmt.Errorf("Expected file got directory '%s'", targetFilename)
	}
	return nil
}
