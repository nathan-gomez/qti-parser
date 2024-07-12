package cmd

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/nathan-gomez/qti-parser/models"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "A brief description of your command",
	Run:   cmdFunc,
	Args:  cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(parseCmd)
}

func cmdFunc(cmd *cobra.Command, args []string) {
	filePath := args[0]
	fmt.Printf("%v", filePath)

	xmlFile, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var items models.Item

	err = xml.Unmarshal(xmlFile, &items)
	if err != nil {
		fmt.Printf("error: %v", err)
		panic(err)
	}

	for _, v := range items.ResponseDeclaration.CorrectResponse {
		fmt.Printf("\n MapKey: %v", v)
	}
}
