/*
Copyright © 2021 NAME HERE <rick.kvm@gmail.com>

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

var cfgFile string
var targetFilename string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fops",
	Short: "File Ops",
	Long:  `File Operaters with linecount and checksum`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
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
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	//cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fops.yaml)")

	rootCmd.PersistentFlags().StringVarP(&targetFilename, "file", "f", "", "")
	//rootCmd.MarkPersistentFlagRequired("file")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
