package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	filePath := flag.String("p", "none", "pathname of the csv file")
	flag.Parse()

	ext := filepath.Ext(*filePath)
	dir, fileName := filepath.Split(*filePath)
	fileNameWithOutExt := strings.Split(fileName, ".")

	if ext != ".csv" {
		log.Fatal("File must of the .csv type!")
	}

	fmt.Printf("filePath: %v\n", *filePath)

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Printf("Fatal Error: %v\n", err)
	}

	defer file.Close()

	record := csv.NewReader(file)
	records, err := record.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	var header []string
	var body []any

	for idx, text := range records {
		if idx == 0 {
			header = text
		} else {

			_map := make(map[string]any)
			for index, val := range text {
				_map[header[index]] = val

			}
			body = append(body, _map)

		}
	}

	data, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(dir+fileNameWithOutExt[0]+".json", data, 0666)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %v\n", string(data))

}
