package main

import (
	"csv_svc/csv"
	"fmt"
)

func main() {
	csvPath := "./data/sa_data.csv"
	jsonPath := "./data/sa_data.json"
	destinationFile := "./data/data.json"
	extractedData := "./data/products.json"
	csv.CsvToJson(csvPath, jsonPath)
	csv.Transform(jsonPath, destinationFile)
	keysToExtract := []string{"id", "sku", "supplier_id"}
	data, err := csv.Extract(destinationFile, keysToExtract)

	if err != nil {
		fmt.Println(err)
	}

	csv.WriteJsonToFile(extractedData, data)
}
