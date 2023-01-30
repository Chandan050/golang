package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%T\n", scanner)

	scanner.Scan()
	text := scanner.Text()
	bytes := scanner.Bytes()

	fmt.Println("input text", text)
	fmt.Println("input bytes", bytes)

	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("you entered:", text)
		fmt.Println(scanner.Bytes())
		if text == "exit" {
			fmt.Println("Exiting the Scannig..........")
			break
		}
	}
	fmt.Println("just after the for loop")

}
