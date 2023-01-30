package main

import (
	f "fmt"
	"log"
	"os"
)

func main() {

	//OS package

	//new pfile
	var newFile *os.File
	f.Printf("%T\n", newFile)

	fileNew, err := os.Create("a.txt")

	if err != nil {
		log.Fatal(err)
	}
	fileNew.Close()

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")

	p := f.Println
	f.Println("fileInfo", fileInfo.Name())
	f.Println("size in bytes", fileInfo.Size())
	p("last modified", fileInfo.ModTime())
	p("its directory", fileInfo.IsDir())
	p("permissions", fileInfo.Mode())

	// renaming the file
	os.Rename("a.txt", "chandan.txt")

	os.Remove("chandan.txt")

}
