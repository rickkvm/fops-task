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
	"bytes"
	"fmt"
	"github.com/rickkvm/fops-task/checksum"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var md5 bool
var sha1 bool
var sha256 bool
var algorithm checksum.HashAlgorithm

// checksumCmd represents the checksum command
var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Print checksum of file",
	Long:  "Print checksum of file.",
	PreRun: func(cmd *cobra.Command, args []string) {
		if md5 {
			algorithm = checksum.MD5
			return
		}
		if sha1 {
			algorithm = checksum.SHA1
			return
		}
		if sha256 {
			algorithm = checksum.SHA256
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var e error
		var file io.Reader
		if targetFilename == "" {
			file = bytes.NewReader([]byte(""))
		} else {
			file, e = os.Open(targetFilename)
		}
		if e != nil {
			fmt.Printf("Open file Error [%s]\n", e.Error())
			os.Exit(1)
		}
		result := checksum.ChecksumHash(file, checksum.HashAlgorithm(algorithm))
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(checksumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checksumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	checksumCmd.Flags().BoolVar(&md5, "md5", false, "check sum with md5 algorithm")
	checksumCmd.Flags().BoolVar(&sha1, "sha1", false, "check sum with sha1 algorithm")
	checksumCmd.Flags().BoolVar(&sha256, "sha256", false, "check sum with sha256 algorithm")
}
