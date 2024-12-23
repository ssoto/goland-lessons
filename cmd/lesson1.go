/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lesson1Cmd represents the lesson1 command
var lesson1Cmd = &cobra.Command{
	Use:   "lesson1",
	Short: "Lesson 1 command",
	Long: `Write a program which will find all such numbers which are divisible by 7 but are not a multiple of 5, 
between 2000 and 3200 (both included). The numbers obtained should be printed in a comma-separated sequence on a 
single line.

Hint: Consider using strconv and strings.Join
`,
	Run: func(cmd *cobra.Command, args []string) {
		var solution string = ""
		var total int = 0
		for i := 2000; i <= 3200; i++ {
			if i%7 == 0 && i%5 != 0 {
				if i < 3200 {
					solution += fmt.Sprintf("%d,", i)
					total += 1
				}
			}
		}
		// remove last comma
		solution = solution[:len(solution)-1]

		fmt.Println("lesson1 solution has", total, "elements: ", solution)

	},
}

func init() {
	rootCmd.AddCommand(lesson1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lesson1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lesson1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
