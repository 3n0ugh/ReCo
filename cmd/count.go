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

	"github.com/3n0ugh/ReCo/counters"
	"github.com/spf13/cobra"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// checking args...
		if len(args) == 1 {

			// calling the read and calculate function
			outs, err := counters.CountOthers(args[0])
			if err != nil {
				fmt.Println(err)
			} else {
				// checking flags here
				arr := []string{"vowel", "consonant", "letter", "space", "digit", "pmark"}
				for order, f := range arr {
					fstatus, _ := cmd.Flags().GetBool(f)
					// if the flag called we print the count
					if fstatus {
						fmt.Printf("%s count: %d\n", f, outs[order])
					}
				}
			}

			// calling the wordCount function
			results, err := counters.CountWord(args[0])
			if err != nil {
				fmt.Println(err)
			} else {
				// checking word flag calling or not
				fstatus, _ := cmd.Flags().GetBool("word")
				if fstatus {
					// if word flag called we print the count
					fmt.Printf("Word Count: %d\n", results)
				}
			}
		} else {
			fmt.Println("just entering one filename")
		}

	},
}

func init() {
	rootCmd.AddCommand(countCmd)
	countCmd.Flags().BoolP("word", "w", false, "Show Word Count")
	countCmd.Flags().BoolP("vowel", "v", false, "Show Vowel Count")
	countCmd.Flags().BoolP("consonant", "c", false, "Show Consonant Count")
	countCmd.Flags().BoolP("letter", "l", false, "Show Letter Count")
	countCmd.Flags().BoolP("space", "s", false, "Show Space Count")
	countCmd.Flags().BoolP("digit", "d", false, "Show Digit Count")
	countCmd.Flags().BoolP("pmark", "p", false, "Show Punctuation Mark Count")
	countCmd.Flags().BoolP("time", "t", false, "Show Elapsed Time")
}
