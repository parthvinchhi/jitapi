package jitapi

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func CreateDir(dirName string) error {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		// Create the new directory
		err = os.Mkdir(dirName, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func CreateFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func ReadCsvFile(inputCsv string) [][]string {
	file, err := os.Open(inputCsv)
	if err != nil {
		log.Fatalf("error opening file %s: %v", inputCsv, err)
	}
	defer file.Close()

	read := csv.NewReader(file)
	records, err := read.ReadAll()
	if err != nil {
		log.Fatalf("error reading file %s: %v", inputCsv, err)
	}

	return records
}

func ReadJsonFile(jsonFile string) ([]Data, error) {
	bytes, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	var data []Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func WriteDataToJson(filename string, data []Data) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func WriteDataToCsv(filename string, data [][]string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range data {
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("could not write row to CSV: %v", err)
		}
	}

	return nil
}
