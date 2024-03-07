package csv

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Csv []struct{}

func CsvToJson(csvFilePath string, jsonFilePath string) error {
	// Open the CSV file
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	// Create a CSV reader
	reader := csv.NewReader(csvFile)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	WriteJsonToFile(jsonFilePath, records)
	return nil
}

func Transform(jsonFilePath string, dataFilePath string) error {
	// open the json file
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// reader the json data
	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	var jsonData [][]string
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}
	objArr := []map[string]any{}
	obj := make(map[string]any)
	headers := jsonData[0]
	objects := jsonData[1:]

	for i := 0; i < len(objects); i++ {
		row := objects[i]
		for j := 0; j < len(row); j++ {
			obj[headers[j]] = row[j]
		}
		objArr = append(objArr, obj)
	}

	WriteJsonToFile(dataFilePath, objArr)
	return nil
}

func Extract(jsonFilePath string, keys []string) ([]map[string]interface{}, error) {
	// open the json file
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	// reader the json data
	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var jsonData []map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}
	objectArray := []map[string]interface{}{}
	for i := 0; i < len(jsonData); i++ {
		// Create a new map for each iteration
		object := make(map[string]interface{})

		// Extract specified keys
		for _, key := range keys {
			if value, ok := jsonData[i][key]; ok {
				object[key] = value
			}
		}

		objectArray = append(objectArray, object)
	}

	return objectArray, nil
}

func WriteJsonToFile(path string, data any) error {
	// Create a new JSON file
	dataFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dataFile.Close()

	// Encode CSV data to JSON and write to the file
	jsonEncoder := json.NewEncoder(dataFile)
	err = jsonEncoder.Encode(data)
	if err != nil {
		return err
	}

	fmt.Printf("Conversion complete. JSON data written to %s\n", path)
	return nil
}
