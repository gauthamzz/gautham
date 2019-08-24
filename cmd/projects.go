/*
Copyright © 2019 GAUTHAM SANTHOSH thabeatsz@gmail.com

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
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "A brief description of my projects",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
		👔  Machine Learning Jobs List Machine Learning Jobs Portal.

		🚀  ayyo - a paywall with one line of code. Ayyo won 1k euro bounty at ETHParis.
		
		🐘  Anna Assistant Your personal assistant for Chrome. #3 Product Of the Day
		
		🎯  Karma - The Smart Habit Tracker. #4 Product of the Day.
		
		✨  Hazel - Build, train, and ship custom deep learning models using a simple visual interface. Won Prototype.
		
		🙏  Pelican Facebook Make your Facebook experince minimal and focused. Featured Product.
		
		🚀  Passerine Product Hunt Desktop Client`)
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
