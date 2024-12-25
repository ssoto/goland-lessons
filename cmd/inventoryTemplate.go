/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"text/template"

	"github.com/spf13/cobra"
)

// inventoryTemplate represents the lesson3 command
var inventoryTemplate = &cobra.Command{
	Use:   "inventoryTemplate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			fmt.Println("Invalid call, please enter a valid string and integer: ", args)
			return
		}

		templatePath := args[0]
		// check if the file path exists
		if _, err := os.Stat(templatePath); os.IsNotExist(err) {
			fmt.Println("Invalid file path, please enter a valid file path.")
			return
		}

		material := args[1]
		// extract second argument and convert to uint
		count, err := strconv.ParseUint(args[2], 10, 64)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid integer.")
			return
		}

		type Inventory struct {
			Material string
			Count    uint
		}

		sweaters := Inventory{material, uint(count)}
		tmpl, err := template.ParseFiles(templatePath)
		if err != nil {
			fmt.Println("Error parsing template:", err)
			return
		} else {
			err = tmpl.Execute(os.Stdout, sweaters)
			if err != nil {
				fmt.Println("Error executing template:", err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(inventoryTemplate)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inventoryTemplate.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inventoryTemplate.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
