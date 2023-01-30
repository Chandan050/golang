package main

import (
	"bufio"
	"fmt"
	"go/scanner"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// file, err := os.Create("text.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// byteSlice := make([]byte, 2)

	// numberBYtesRead, err := io.ReadFull(file, byteSlice)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("number of bytes %d \n", numberBYtesRead)
	// log.Printf("the data in it %s\n", byteSlice)

	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("data in that pkg %#v %d\n", string(data), len(data))

	data1, err1 := ioutil.ReadFile("write_file.go")
	if err1 != nil {
		log.Fatal(err)
	}
	fmt.Printf("data in that pkg %s %d\n", data1, len(data1))

	fmt.Println("bufio scanning ---------------------------------------------------")

	file2, err2 := os.Open("text.txt")
	if err1 != nil {
		log.Fatal(err2)
	}
	defer file2.Close()

	scanners := bufio.NewScanner(file2)

	success := scanners.Scan()
	if success == false {
		err2 = scanner.Error{}
		if err2 == nil {
			log.Fatal("scan was completed and it reached EOF")
		} else {
			log.Fatal(err)
		}
	}

}
