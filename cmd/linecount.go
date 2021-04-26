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
	"github.com/rickkvm/fops-task/linecount"
	"github.com/spf13/cobra"
	"os"
)

// linecountCmd represents the linecount command
var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Print line count of file",
	RunE: func(cmd *cobra.Command, args []string) error {
		e := checkFileExist(targetFilename)
		if e != nil {
			return e
		}
		file, e := os.Open(targetFilename)
		if e != nil {
			return e
		}
		result, e := linecount.Count(file)
		if e == linecount.ErrCountingBinary {
			return fmt.Errorf("error: Cannot do linecount for binary file '%s'", targetFilename)
		}
		if e != nil {
			return e
		}
		fmt.Println(result)
		return nil
	},
}

func init() {
	//rootCmd.AddCommand(linecountCmd)
	linecountCmd.Flags().StringVarP(&targetFilename, "file", "f", "", "the input file")
	linecountCmd.Flags().BoolP("help", "h", false, "Print usage")
	linecountCmd.Flags().MarkHidden("help")
	linecountCmd.MarkFlagRequired("file")
}
