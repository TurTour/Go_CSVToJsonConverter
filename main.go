package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	fmt.Println("Info: Source file needs to be in current exe running directory")
	
	//get source file Name
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file name: ")
	fileName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fileName = strings.TrimSuffix(fileName, "\n")
	fileName = strings.TrimSuffix(fileName, "\r")
	//*****

	//set destination file Name
	fmt.Print("Enter resulting file name: ")
	resultingfileName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	resultingfileName = strings.TrimSuffix(resultingfileName, "\n")
	resultingfileName = strings.TrimSuffix(resultingfileName, "\r")
	//*****

	//get current working directory
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	//*****

	records, err := readCSV(exPath + "\\" + fileName)
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	jsonOutput, err := csvToJSON(records)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	err = os.WriteFile(exPath+"\\"+resultingfileName, jsonOutput, os.ModePerm)
	if err != nil {
		fmt.Println("Error writing to JSON file:", err)
		return
	}

	fmt.Println("JSON data writen to file")
}
