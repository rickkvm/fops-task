/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"io/ioutil"
	"os"

	"github.com/rickkvm/fops-task/linecount"
	"github.com/spf13/cobra"
)

// linecountCmd represents the linecount command
var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Print line count of file",
	Long:  `Print line count of file`,
	Run: func(cmd *cobra.Command, args []string) {

		data, e := ioutil.ReadFile(targetFilename)
		if e != nil {
			fmt.Println("error: Cannot open file '%s'\n", targetFilename)
			os.Exit(1)
		}
		result := linecount.Count(data)
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(linecountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// linecountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// linecountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
