package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	// Put the file in FileNameChanger folder and run bellow command
	// go run main.go oldFileName newFileName

	baseDir := "/home/ubuntu/Documents/my_projects/go-small-tools/fileNameChanger"
	fmt.Println("Main function starts")
	oldFileName := os.Args[1]
	newFileName := os.Args[2]

	oldFileFullPath := path.Join(baseDir, "/", oldFileName)
	newFileFullPath := path.Join(baseDir, "/", newFileName)

	err := os.Rename(oldFileFullPath, newFileFullPath)

	if err != nil {
		fmt.Println("Fatal error")
		log.Fatal(err)
	} else {
		fmt.Println("File name changed")
		// log.Default()
	}

}
