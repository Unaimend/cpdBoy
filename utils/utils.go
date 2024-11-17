// utils/utils.go

package utils

import (
	"bufio"
	"os"
	//"fmt"
	"strings"
)

type Row = map[string]string
type DataBase = []map[string]string

func ReadTSV(filename string) (DataBase, error) {
	// Open the TSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []map[string]string
	var headers []string

	// Create a scanner to read through the file
	scanner := bufio.NewScanner(file)
	lineCount := 0

	// Read each line
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")

		if lineCount == 0 {
			// First line contains headers
			headers = fields
		} else {
			// Create a map for each row, using headers as keys
			row := make(map[string]string)
			for i, value := range fields {
				if i < len(headers) {
					row[headers[i]] = value
				}
			}
			result = append(result, row)
		}
		lineCount++
	}

	// Check for errors from scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}


// filterBy filters the data by a specific column and value.
func FilterBy(data []map[string]string, column string, value string) DataBase {
	var filtered []map[string]string
	for _, row := range data {
		if row[column] == value {
			filtered = append(filtered, row)
		}
	}
	return filtered
}
