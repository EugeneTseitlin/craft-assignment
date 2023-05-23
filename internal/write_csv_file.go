package internal

import (
	"encoding/csv"
	"os"
)


func WriteCsvFile(path string, data [][]string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.WriteAll(data)
	writer.Flush()
	return nil
}

