package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	// by using os madule

	file, err := os.OpenFile("c.txt", os.O_CREATE|os.O_APPEND|os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	bs := []byte("i love golang!")
	bWritten, err := file.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(bWritten)
	file.Close()

	// by using "ioutil" package

	// bufio package

	fl, err := os.Create("file_text.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer fl.Close()
	bufferedWriter := bufio.NewWriter(fl)
	ba := []byte{97, 98, 99}
	bytesWritten, err := bufferedWriter.Write(ba)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("bytes written to buffer (not file) %d \n", bytesWritten)

	bytesAvailable := bufferedWriter.Available()
	log.Println(bytesAvailable)

}
