package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
)

func readCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func csvToJSON(records [][]string) ([]byte, error) {

	var jsonData []map[string]string
	headers := records[0]

	for _, record := range records[1:] {
		rowData := make(map[string]string)

		for i, header := range headers {
			rowData[header] = record[i]
		}

		jsonData = append(jsonData, rowData)
	}

	return json.MarshalIndent(jsonData, "", " ")

}
