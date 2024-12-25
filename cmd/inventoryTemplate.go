/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

var templatePath string
var material string
var count uint

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

		// check if the path is valid
		if _, err := os.Stat(templatePath); os.IsNotExist(err) {
			fmt.Println("Template file does not exist")
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
	inventoryTemplate.Flags().StringVarP(&templatePath, "template", "t", "", "Path to the template file")
	inventoryTemplate.MarkFlagRequired("template")
	inventoryTemplate.Flags().StringVarP(&material, "material", "m", "", "Material of the inventory")
	inventoryTemplate.MarkFlagRequired("material")
	inventoryTemplate.Flags().UintVarP(&count, "count", "c", 0, "Count of the inventory")
	inventoryTemplate.MarkFlagRequired("count")
}
