package expense

import (
	"encoding/csv"
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

func writeHeaderToCSVFile(file *os.File) {

	writeFile := csv.NewWriter(file)
	defer writeFile.Flush()

	row := []string{"Title", "Amount"}
	err := writeFile.Write(row)
	if err != nil {
		log.Fatal(err)
	}
}

func writeToCSVFile(file *os.File, employee Employee) {

	writeFile := csv.NewWriter(file)
	defer writeFile.Flush()

	row := []string{employee.id, strconv.Itoa(employee.age)}
	err := writeFile.Write(row)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateOrAddRecord(data []string) {
	baseDir := "/home/ubuntu/Documents/my_projects/go-small-tools/expense-tracker"
	fileName := "expense.csv"
	filePath := baseDir + "/" + fileName
	age, _ := strconv.Atoi(data[1])
	employee := Employee{data[0], age}

	if isFileExist(filePath) {
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		writeToCSVFile(file, employee)
	} else {
		csvFile, err := os.Create(filePath)
		defer csvFile.Close()
		if err != nil {
			log.Fatal(err)
		}

		writeHeaderToCSVFile(csvFile)
		writeToCSVFile(csvFile, employee)
	}
}
