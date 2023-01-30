package main

import (
	"bufio"
	"go/scanner"
	"log"
	"os"
)

func main() {
	file, err := os.Open("read_file.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanners := bufio.NewScanner(file)

	success := scanners.Scan()
	if success == false {
		err = scanner.Error{}
		if err == nil {
			log.Fatal("scan was completed and it reached EOF")
		} else {
			log.Fatal(err)
		}
	}

}
