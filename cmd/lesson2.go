/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// lesson2Cmd represents the lesson2 command
var lesson2Cmd = &cobra.Command{
	Use:   "lesson2",
	Short: "A brief description of your command",
	Long: `Write a program which can compute the factorial of a given numbers. The results should be printed in a 
comma-separated sequence on a single line.
Suppose the following input is supplied to the program: 8
Then, the output should be: 40320

Hints: In case of input data being supplied to the question, it should be assumed to be a console input.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Invalid number of arguments, please enter a valid integer.")
			return
		}
		var input = args[0]
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid integer.")
			return
		}

		var factorial int = 1
		for i := 1; i <= number; i++ {
			factorial *= i
		}

		fmt.Println("lesson2 solution: the factorial of", input, "is: ", factorial)
	},
}

func init() {
	rootCmd.AddCommand(lesson2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lesson2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lesson2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
