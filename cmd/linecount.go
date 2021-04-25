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
	"io/ioutil"
	"os"

	"github.com/rickkvm/fops-task/linecount"
	"github.com/spf13/cobra"
)

// linecountCmd represents the linecount command
var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Print line count of file",
	Run: func(cmd *cobra.Command, args []string) {
		if targetFilename == "" {
			cmd.Help()
			os.Exit(1)
		}
		stat, e := os.Lstat(targetFilename)
		if e != nil {
			if os.IsNotExist(e) {
				fmt.Printf("error: No such file '%s'\n", targetFilename)
				os.Exit(1)
			}
			fmt.Println("unknwon error", e.Error())
			os.Exit(1)
		}

		if stat.IsDir() {
			fmt.Printf("error: Expected file got directory '%s'\n", targetFilename)
			os.Exit(1)
		}
		data, e := ioutil.ReadFile(targetFilename)
		if e != nil {
			fmt.Printf("error: Cannot open file '%s'\n", targetFilename)
			os.Exit(1)
		}
		result := linecount.Count(data)
		fmt.Println(result)
	},
}

func init() {
	//rootCmd.AddCommand(linecountCmd)
	linecountCmd.Flags().StringVarP(&targetFilename, "file", "f", "", "the input file")
	linecountCmd.Flags().BoolP("help", "h", false, "Print usage")
	linecountCmd.Flags().MarkHidden("help")
}
