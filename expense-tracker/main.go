package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Employee struct {
	id  string
	age int
}

func isFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	return false
}

func writeToCSVFile(file *os.File) {
	records := []Employee{
		Employee{"E01", 45},
		Employee{"E02", 50},
	}

	writeFile := csv.NewWriter(file)
	defer writeFile.Flush()

	for _, record := range records {
		row := []string{record.id, strconv.Itoa(record.age)}
		err := writeFile.Write(row)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	fmt.Println("Starting application")
	baseDir := "/home/ubuntu/Documents/my_projects/go-small-tools/expense-tracker"
	fileName := "test.csv"
	filePath := baseDir + "/" + fileName

	records := []Employee{
		Employee{"E01", 45},
		Employee{"E02", 50},
	}
	if isFileExist(filePath) {
		fmt.Println("File exists")
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		writeToCSVFile(file)
	} else {
		fmt.Println("File not found")
		csvFile, err := os.Create(filePath)
		defer csvFile.Close()
		if err != nil {
			log.Fatal(err)
		}

		writeFile := csv.NewWriter(csvFile)
		defer writeFile.Flush()

		for _, record := range records {
			row := []string{record.id, strconv.Itoa(record.age)}
			fmt.Println(row)
			err := writeFile.Write(row)
			if err != nil {
				log.Fatal(err)
			}
		}

		// var data [][]string

		// for _, record := range records {
		// 	row := []string{record.id, strconv.Itoa(record.age)}
		// 	data = append(data, row)
		// }

		// writeFile.WriteAll(data)

	}
}
