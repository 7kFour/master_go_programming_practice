package main

import (
	"log"
	"os"
)

func main () {
	// ex 1
	// creating a new file
	theFile, err := os.Create("the_file_name.txt")

	// checking for errors
	if err != nil {
		//log the error and exit the program
		log.Fatal(err)
	}

	// idiomatic way to handle file closure
	//defer theFile.Close()
	// closing immediately so that we can rename the file later and not wait to exit main func
	theFile.Close()

	// ex 2
	// check if file exists
	fileInfo, err := os.Stat("the_file_name.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist")
		}
	}
	_ = fileInfo

	// rename file
	oldName := "the_file_name.txt"
	newName := "data.txt"
	err = os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}

	// ex 3
	// remove the file
	err = os.Remove(newName)
	if err != nil{
		log.Fatal(err)
	}
}
