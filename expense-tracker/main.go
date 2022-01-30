package main

import (
	"fmt"
	"os"

	"./expense"
)

func main() {
	fmt.Println("Starting application")

	title := os.Args[1]
	amount := os.Args[2]

	expense.CreateOrAddRecord([]string{title, amount})

}
