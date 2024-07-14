package cmd

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/nathan-gomez/qti-parser/models"
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
	var items models.QtiItem

	err := parseXml(filePath, &items)
	if err != nil {
		panic(err)
	}

	err = writeCsv(&items)
	if err != nil {
		fmt.Printf("an error has ocurred: %s", err)
		panic(err)
	}
}

func parseXml(filePath string, qtiItems *models.QtiItem) error {
	var err error

	xmlFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(xmlFile, qtiItems)
	if err != nil {
		return err
	}

	return nil
}

func writeCsv(qtiItem *models.QtiItem) error {
	csvFile, err := os.Create("test.csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	headers, err := csvHeaders()
	if err != nil {
		return err
	}

	err = csvWriter.Write(headers)
	if err != nil {
		return err
	}

	// for _, item := range qtiItems {
	csvString, err := toStringArr(qtiItem)
	if err != nil {
		return err
	}

	err = csvWriter.Write(csvString)
	if err != nil {
		return err
	}
	// }
	return nil
}

func csvHeaders() ([]string, error) {
	headers := []string{"Identifier", "Cardinality", "BaseType", "CorrectResponse"}
	return headers, nil
}

func toStringArr(qtiItem *models.QtiItem) ([]string, error) {
	// for _, item := range qtiItem.ResponseDeclaration {
	// }
	item := qtiItem.ResponseDeclaration[0]
	row := []string{item.Identifier, item.Cardinality, item.BaseType, ""}

	return row, nil
}
